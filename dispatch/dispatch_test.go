package dispatch

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/types"
)

func TestAggrGroup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lset := model.LabelSet{"a": "v1", "b": "v2"}
	opts := &RouteOpts{Receiver: "n1", GroupBy: map[model.LabelName]struct{}{}, GroupWait: 1 * time.Second, GroupInterval: 300 * time.Millisecond, RepeatInterval: 1 * time.Hour}
	route := &Route{RouteOpts: *opts}
	var (
		a1	= &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "v1", "b": "v2", "c": "v3"}, StartsAt: time.Now().Add(time.Minute), EndsAt: time.Now().Add(time.Hour)}, UpdatedAt: time.Now()}
		a2	= &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "v1", "b": "v2", "c": "v4"}, StartsAt: time.Now().Add(-time.Hour), EndsAt: time.Now().Add(2 * time.Hour)}, UpdatedAt: time.Now()}
		a3	= &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "v1", "b": "v2", "c": "v5"}, StartsAt: time.Now().Add(time.Minute), EndsAt: time.Now().Add(5 * time.Minute)}, UpdatedAt: time.Now()}
	)
	var (
		last		= time.Now()
		current		= time.Now()
		lastCurMtx	= &sync.Mutex{}
		alertsCh	= make(chan types.AlertSlice)
	)
	ntfy := func(ctx context.Context, alerts ...*types.Alert) bool {
		if _, ok := notify.Now(ctx); !ok {
			t.Errorf("now missing")
		}
		if _, ok := notify.GroupKey(ctx); !ok {
			t.Errorf("group key missing")
		}
		if lbls, ok := notify.GroupLabels(ctx); !ok || !reflect.DeepEqual(lbls, lset) {
			t.Errorf("wrong group labels: %q", lbls)
		}
		if rcv, ok := notify.ReceiverName(ctx); !ok || rcv != opts.Receiver {
			t.Errorf("wrong receiver: %q", rcv)
		}
		if ri, ok := notify.RepeatInterval(ctx); !ok || ri != opts.RepeatInterval {
			t.Errorf("wrong repeat interval: %q", ri)
		}
		lastCurMtx.Lock()
		last = current
		current = time.Now().Add(-time.Millisecond)
		lastCurMtx.Unlock()
		alertsCh <- types.AlertSlice(alerts)
		return true
	}
	removeEndsAt := func(as types.AlertSlice) types.AlertSlice {
		for i, a := range as {
			ac := *a
			ac.EndsAt = time.Time{}
			as[i] = &ac
		}
		return as
	}
	ag := newAggrGroup(context.Background(), lset, route, nil, log.NewNopLogger())
	go ag.run(ntfy)
	ag.insert(a1)
	select {
	case <-time.After(2 * opts.GroupWait):
		t.Fatalf("expected initial batch after group_wait")
	case batch := <-alertsCh:
		lastCurMtx.Lock()
		s := time.Since(last)
		lastCurMtx.Unlock()
		if s < opts.GroupWait {
			t.Fatalf("received batch too early after %v", s)
		}
		exp := removeEndsAt(types.AlertSlice{a1})
		sort.Sort(batch)
		if !reflect.DeepEqual(batch, exp) {
			t.Fatalf("expected alerts %v but got %v", exp, batch)
		}
	}
	for i := 0; i < 3; i++ {
		ag.insert(a3)
		select {
		case <-time.After(2 * opts.GroupInterval):
			t.Fatalf("expected new batch after group interval but received none")
		case batch := <-alertsCh:
			lastCurMtx.Lock()
			s := time.Since(last)
			lastCurMtx.Unlock()
			if s < opts.GroupInterval {
				t.Fatalf("received batch too early after %v", s)
			}
			exp := removeEndsAt(types.AlertSlice{a1, a3})
			sort.Sort(batch)
			if !reflect.DeepEqual(batch, exp) {
				t.Fatalf("expected alerts %v but got %v", exp, batch)
			}
		}
	}
	ag.stop()
	ag = newAggrGroup(context.Background(), lset, route, nil, log.NewNopLogger())
	go ag.run(ntfy)
	ag.insert(a1)
	ag.insert(a2)
	select {
	case <-time.After(opts.GroupWait / 2):
		t.Fatalf("expected immediate alert but received none")
	case batch := <-alertsCh:
		exp := removeEndsAt(types.AlertSlice{a1, a2})
		sort.Sort(batch)
		if !reflect.DeepEqual(batch, exp) {
			t.Fatalf("expected alerts %v but got %v", exp, batch)
		}
	}
	for i := 0; i < 3; i++ {
		ag.insert(a3)
		select {
		case <-time.After(2 * opts.GroupInterval):
			t.Fatalf("expected new batch after group interval but received none")
		case batch := <-alertsCh:
			lastCurMtx.Lock()
			s := time.Since(last)
			lastCurMtx.Unlock()
			if s < opts.GroupInterval {
				t.Fatalf("received batch too early after %v", s)
			}
			exp := removeEndsAt(types.AlertSlice{a1, a2, a3})
			sort.Sort(batch)
			if !reflect.DeepEqual(batch, exp) {
				t.Fatalf("expected alerts %v but got %v", exp, batch)
			}
		}
	}
	a1r, a2r, a3r := *a1, *a2, *a3
	resolved := types.AlertSlice{&a1r, &a2r, &a3r}
	for _, a := range resolved {
		a.EndsAt = time.Now()
		ag.insert(a)
	}
	select {
	case <-time.After(2 * opts.GroupInterval):
		t.Fatalf("expected new batch after group interval but received none")
	case batch := <-alertsCh:
		lastCurMtx.Lock()
		s := time.Since(last)
		lastCurMtx.Unlock()
		if s < opts.GroupInterval {
			t.Fatalf("received batch too early after %v", s)
		}
		sort.Sort(batch)
		if !reflect.DeepEqual(batch, resolved) {
			t.Fatalf("expected alerts %v but got %v", resolved, batch)
		}
		if !ag.empty() {
			t.Fatalf("Expected aggregation group to be empty after resolving alerts: %v", ag)
		}
	}
	ag.stop()
}
func TestGroupLabels(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var a = &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "v1", "b": "v2", "c": "v3"}}}
	route := &Route{RouteOpts: RouteOpts{GroupBy: map[model.LabelName]struct{}{"a": struct{}{}, "b": struct{}{}}, GroupByAll: false}}
	expLs := model.LabelSet{"a": "v1", "b": "v2"}
	ls := getGroupLabels(a, route)
	if !reflect.DeepEqual(ls, expLs) {
		t.Fatalf("expected labels are %v, but got %v", expLs, ls)
	}
}
func TestGroupByAllLabels(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var a = &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "v1", "b": "v2", "c": "v3"}}}
	route := &Route{RouteOpts: RouteOpts{GroupBy: map[model.LabelName]struct{}{}, GroupByAll: true}}
	expLs := model.LabelSet{"a": "v1", "b": "v2", "c": "v3"}
	ls := getGroupLabels(a, route)
	if !reflect.DeepEqual(ls, expLs) {
		t.Fatalf("expected labels are %v, but got %v", expLs, ls)
	}
}
