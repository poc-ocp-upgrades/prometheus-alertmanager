package inhibit

import (
	"testing"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/provider"
	"github.com/prometheus/alertmanager/store"
	"github.com/prometheus/alertmanager/types"
)

var nopLogger = log.NewNopLogger()

func TestInhibitRuleHasEqual(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	now := time.Now()
	cases := []struct {
		initial	map[model.Fingerprint]*types.Alert
		equal	model.LabelNames
		input	model.LabelSet
		result	bool
	}{{initial: map[model.Fingerprint]*types.Alert{}, input: model.LabelSet{"a": "b"}, result: false}, {initial: map[model.Fingerprint]*types.Alert{1: &types.Alert{}}, input: model.LabelSet{"a": "b"}, result: true}, {initial: map[model.Fingerprint]*types.Alert{1: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "b", "b": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second)}}, 2: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "b", "b": "c"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second)}}}, equal: model.LabelNames{"a", "b"}, input: model.LabelSet{"a": "b", "b": "c"}, result: false}, {initial: map[model.Fingerprint]*types.Alert{1: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "b", "c": "d"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second)}}, 2: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "b", "c": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(time.Hour)}}}, equal: model.LabelNames{"a"}, input: model.LabelSet{"a": "b"}, result: true}, {initial: map[model.Fingerprint]*types.Alert{1: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "c", "c": "d"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second)}}, 2: &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"a": "c", "c": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second)}}}, equal: model.LabelNames{"a"}, input: model.LabelSet{"a": "b"}, result: false}}
	for _, c := range cases {
		r := &InhibitRule{Equal: map[model.LabelName]struct{}{}, scache: store.NewAlerts(5 * time.Minute)}
		for _, ln := range c.equal {
			r.Equal[ln] = struct{}{}
		}
		for _, v := range c.initial {
			r.scache.Set(v)
		}
		if _, have := r.hasEqual(c.input); have != c.result {
			t.Errorf("Unexpected result %t, expected %t", have, c.result)
		}
	}
}
func TestInhibitRuleMatches(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	cr := config.InhibitRule{SourceMatch: map[string]string{"s": "1"}, TargetMatch: map[string]string{"t": "1"}, Equal: model.LabelNames{"e"}}
	m := types.NewMarker()
	ih := NewInhibitor(nil, []*config.InhibitRule{&cr}, m, nopLogger)
	ir := ih.rules[0]
	now := time.Now()
	sourceAlert := &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"s": "1", "e": "1"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(time.Hour)}}
	ir.scache = store.NewAlerts(5 * time.Minute)
	ir.scache.Set(sourceAlert)
	cases := []struct {
		target		model.LabelSet
		expected	bool
	}{{target: model.LabelSet{"t": "1", "e": "1"}, expected: true}, {target: model.LabelSet{"t": "1", "t2": "1", "e": "1"}, expected: true}, {target: model.LabelSet{"t": "0", "e": "1"}, expected: false}, {target: model.LabelSet{"s": "1", "t": "1", "e": "1"}, expected: false}, {target: model.LabelSet{"t": "1", "e": "0"}, expected: false}}
	for _, c := range cases {
		if actual := ih.Mutes(c.target); actual != c.expected {
			t.Errorf("Expected (*Inhibitor).Mutes(%v) to return %t but got %t", c.target, c.expected, actual)
		}
	}
}

type fakeAlerts struct {
	alerts		[]*types.Alert
	finished	chan struct{}
}

func newFakeAlerts(alerts []*types.Alert) *fakeAlerts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &fakeAlerts{alerts: alerts, finished: make(chan struct{})}
}
func (f *fakeAlerts) GetPending() provider.AlertIterator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (f *fakeAlerts) Get(model.Fingerprint) (*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (f *fakeAlerts) Put(...*types.Alert) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (f *fakeAlerts) Subscribe() provider.AlertIterator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan *types.Alert)
	done := make(chan struct{})
	go func() {
		for _, a := range f.alerts {
			ch <- a
		}
		ch <- &types.Alert{Alert: model.Alert{Labels: model.LabelSet{}, StartsAt: time.Now()}}
		close(f.finished)
		<-done
	}()
	return provider.NewAlertIterator(ch, done, nil)
}
func TestInhibit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	now := time.Now()
	inhibitRule := func() *config.InhibitRule {
		return &config.InhibitRule{SourceMatch: map[string]string{"s": "1"}, TargetMatch: map[string]string{"t": "1"}, Equal: model.LabelNames{"e"}}
	}
	alertOne := func() *types.Alert {
		return &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"t": "1", "e": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(time.Hour)}}
	}
	alertTwo := func(resolved bool) *types.Alert {
		var end time.Time
		if resolved {
			end = now.Add(-time.Second)
		} else {
			end = now.Add(time.Hour)
		}
		return &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"s": "1", "e": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: end}}
	}
	type exp struct {
		lbls	model.LabelSet
		muted	bool
	}
	for i, tc := range []struct {
		alerts		[]*types.Alert
		expected	[]exp
	}{{alerts: []*types.Alert{alertOne()}, expected: []exp{{lbls: model.LabelSet{"t": "1", "e": "f"}, muted: false}}}, {alerts: []*types.Alert{alertOne(), alertTwo(false)}, expected: []exp{{lbls: model.LabelSet{"t": "1", "e": "f"}, muted: true}, {lbls: model.LabelSet{"s": "1", "e": "f"}, muted: false}}}, {alerts: []*types.Alert{alertOne(), alertTwo(false), alertTwo(true)}, expected: []exp{{lbls: model.LabelSet{"t": "1", "e": "f"}, muted: false}, {lbls: model.LabelSet{"s": "1", "e": "f"}, muted: false}}}} {
		ap := newFakeAlerts(tc.alerts)
		mk := types.NewMarker()
		inhibitor := NewInhibitor(ap, []*config.InhibitRule{inhibitRule()}, mk, nopLogger)
		go func() {
			for ap.finished != nil {
				select {
				case <-ap.finished:
					ap.finished = nil
				default:
				}
			}
			inhibitor.Stop()
		}()
		inhibitor.Run()
		for _, expected := range tc.expected {
			if inhibitor.Mutes(expected.lbls) != expected.muted {
				mute := "unmuted"
				if expected.muted {
					mute = "muted"
				}
				t.Errorf("tc: %d, expected alert with labels %q to be %s", i, expected.lbls, mute)
			}
		}
	}
}
