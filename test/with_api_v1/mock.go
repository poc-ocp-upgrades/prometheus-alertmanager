package test

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"sync"
	"time"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/types"
)

func At(ts float64) float64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ts
}

type Interval struct{ start, end float64 }

func (iv Interval) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[%v,%v]", iv.start, iv.end)
}
func (iv Interval) contains(f float64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f >= iv.start && f <= iv.end
}
func Between(start, end float64) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{start: start, end: end}
}

type TestSilence struct {
	id			string
	match			[]string
	matchRE			[]string
	startsAt, endsAt	float64
	mtx			sync.RWMutex
}

func Silence(start, end float64) *TestSilence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &TestSilence{startsAt: start, endsAt: end}
}
func (s *TestSilence) Match(v ...string) *TestSilence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.match = append(s.match, v...)
	return s
}
func (s *TestSilence) MatchRE(v ...string) *TestSilence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(v)%2 == 1 {
		panic("bad key/values")
	}
	s.matchRE = append(s.matchRE, v...)
	return s
}
func (s *TestSilence) SetID(ID string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.id = ID
}
func (s *TestSilence) ID() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.id
}
func (s *TestSilence) nativeSilence(opts *AcceptanceOpts) *types.Silence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nsil := &types.Silence{}
	for i := 0; i < len(s.match); i += 2 {
		nsil.Matchers = append(nsil.Matchers, &types.Matcher{Name: s.match[i], Value: s.match[i+1]})
	}
	for i := 0; i < len(s.matchRE); i += 2 {
		nsil.Matchers = append(nsil.Matchers, &types.Matcher{Name: s.matchRE[i], Value: s.matchRE[i+1], IsRegex: true})
	}
	if s.startsAt > 0 {
		nsil.StartsAt = opts.expandTime(s.startsAt)
	}
	if s.endsAt > 0 {
		nsil.EndsAt = opts.expandTime(s.endsAt)
	}
	nsil.Comment = "some comment"
	nsil.CreatedBy = "admin@example.com"
	return nsil
}

type TestAlert struct {
	labels			model.LabelSet
	annotations		model.LabelSet
	startsAt, endsAt	float64
}

func Alert(keyval ...interface{}) *TestAlert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keyval)%2 == 1 {
		panic("bad key/values")
	}
	a := &TestAlert{labels: model.LabelSet{}, annotations: model.LabelSet{}}
	for i := 0; i < len(keyval); i += 2 {
		ln := model.LabelName(keyval[i].(string))
		lv := model.LabelValue(keyval[i+1].(string))
		a.labels[ln] = lv
	}
	return a
}
func (a *TestAlert) nativeAlert(opts *AcceptanceOpts) *model.Alert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	na := &model.Alert{Labels: a.labels, Annotations: a.annotations}
	if a.startsAt > 0 {
		na.StartsAt = opts.expandTime(a.startsAt)
	}
	if a.endsAt > 0 {
		na.EndsAt = opts.expandTime(a.endsAt)
	}
	return na
}
func (a *TestAlert) Annotate(keyval ...interface{}) *TestAlert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keyval)%2 == 1 {
		panic("bad key/values")
	}
	for i := 0; i < len(keyval); i += 2 {
		ln := model.LabelName(keyval[i].(string))
		lv := model.LabelValue(keyval[i+1].(string))
		a.annotations[ln] = lv
	}
	return a
}
func (a *TestAlert) Active(tss ...float64) *TestAlert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(tss) > 2 || len(tss) == 0 {
		panic("only one or two timestamps allowed")
	}
	if len(tss) == 2 {
		a.endsAt = tss[1]
	}
	a.startsAt = tss[0]
	return a
}
func equalAlerts(a, b *model.Alert, opts *AcceptanceOpts) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !reflect.DeepEqual(a.Labels, b.Labels) {
		return false
	}
	if !reflect.DeepEqual(a.Annotations, b.Annotations) {
		return false
	}
	if !equalTime(a.StartsAt, b.StartsAt, opts) {
		return false
	}
	if !equalTime(a.EndsAt, b.EndsAt, opts) {
		return false
	}
	return true
}
func equalTime(a, b time.Time, opts *AcceptanceOpts) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.IsZero() != b.IsZero() {
		return false
	}
	diff := a.Sub(b)
	if diff < 0 {
		diff = -diff
	}
	return diff <= opts.Tolerance
}

type MockWebhook struct {
	opts		*AcceptanceOpts
	collector	*Collector
	listener	net.Listener
	Func		func(timestamp float64) bool
}

func NewWebhook(c *Collector) *MockWebhook {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l, err := net.Listen("tcp4", "localhost:0")
	if err != nil {
		panic(err)
	}
	wh := &MockWebhook{listener: l, collector: c, opts: c.opts}
	go func() {
		if err := http.Serve(l, wh); err != nil {
			panic(err)
		}
	}()
	return wh
}
func (ws *MockWebhook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ws.Func != nil {
		if ws.Func(ws.opts.relativeTime(time.Now())) {
			return
		}
	}
	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var v notify.WebhookMessage
	if err := dec.Decode(&v); err != nil {
		panic(err)
	}
	var alerts model.Alerts
	for _, a := range v.Alerts {
		var (
			labels		= model.LabelSet{}
			annotations	= model.LabelSet{}
		)
		for k, v := range a.Labels {
			labels[model.LabelName(k)] = model.LabelValue(v)
		}
		for k, v := range a.Annotations {
			annotations[model.LabelName(k)] = model.LabelValue(v)
		}
		alerts = append(alerts, &model.Alert{Labels: labels, Annotations: annotations, StartsAt: a.StartsAt, EndsAt: a.EndsAt, GeneratorURL: a.GeneratorURL})
	}
	ws.collector.add(alerts...)
}
func (ws *MockWebhook) Address() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ws.listener.Addr().String()
}
