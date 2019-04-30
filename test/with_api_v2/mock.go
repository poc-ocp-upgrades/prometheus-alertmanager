package test

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"sync"
	"time"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
	"github.com/go-openapi/strfmt"
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
func (s *TestSilence) nativeSilence(opts *AcceptanceOpts) *models.Silence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nsil := &models.Silence{}
	for i := 0; i < len(s.match); i += 2 {
		nsil.Matchers = append(nsil.Matchers, &models.Matcher{Name: &s.match[i], Value: &s.match[i+1]})
	}
	t := true
	for i := 0; i < len(s.matchRE); i += 2 {
		nsil.Matchers = append(nsil.Matchers, &models.Matcher{Name: &s.matchRE[i], Value: &s.matchRE[i+1], IsRegex: &t})
	}
	if s.startsAt > 0 {
		start := strfmt.DateTime(opts.expandTime(s.startsAt))
		nsil.StartsAt = &start
	}
	if s.endsAt > 0 {
		end := strfmt.DateTime(opts.expandTime(s.endsAt))
		nsil.EndsAt = &end
	}
	comment := "some comment"
	createdBy := "admin@example.com"
	nsil.Comment = &comment
	nsil.CreatedBy = &createdBy
	return nsil
}

type TestAlert struct {
	labels			models.LabelSet
	annotations		models.LabelSet
	startsAt, endsAt	float64
}

func Alert(keyval ...interface{}) *TestAlert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keyval)%2 == 1 {
		panic("bad key/values")
	}
	a := &TestAlert{labels: models.LabelSet{}, annotations: models.LabelSet{}}
	for i := 0; i < len(keyval); i += 2 {
		ln := keyval[i].(string)
		lv := keyval[i+1].(string)
		a.labels[ln] = lv
	}
	return a
}
func (a *TestAlert) nativeAlert(opts *AcceptanceOpts) *models.GettableAlert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	na := &models.GettableAlert{Alert: models.Alert{Labels: a.labels}, Annotations: a.annotations, StartsAt: &strfmt.DateTime{}, EndsAt: &strfmt.DateTime{}}
	if a.startsAt > 0 {
		start := strfmt.DateTime(opts.expandTime(a.startsAt))
		na.StartsAt = &start
	}
	if a.endsAt > 0 {
		end := strfmt.DateTime(opts.expandTime(a.endsAt))
		na.EndsAt = &end
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
		ln := keyval[i].(string)
		lv := keyval[i+1].(string)
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
func equalAlerts(a, b *models.GettableAlert, opts *AcceptanceOpts) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !reflect.DeepEqual(a.Labels, b.Labels) {
		return false
	}
	if !reflect.DeepEqual(a.Annotations, b.Annotations) {
		return false
	}
	if !equalTime(time.Time(*a.StartsAt), time.Time(*b.StartsAt), opts) {
		return false
	}
	if (a.EndsAt == nil) != (b.EndsAt == nil) {
		return false
	}
	if !(a.EndsAt == nil) && !(b.EndsAt == nil) && !equalTime(time.Time(*a.EndsAt), time.Time(*b.EndsAt), opts) {
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
	var alerts models.GettableAlerts
	for _, a := range v.Alerts {
		var (
			labels		= models.LabelSet{}
			annotations	= models.LabelSet{}
		)
		for k, v := range a.Labels {
			labels[k] = v
		}
		for k, v := range a.Annotations {
			annotations[k] = v
		}
		start := strfmt.DateTime(a.StartsAt)
		end := strfmt.DateTime(a.EndsAt)
		alerts = append(alerts, &models.GettableAlert{Alert: models.Alert{Labels: labels, GeneratorURL: strfmt.URI(a.GeneratorURL)}, Annotations: annotations, StartsAt: &start, EndsAt: &end})
	}
	ws.collector.add(alerts...)
}
func (ws *MockWebhook) Address() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ws.listener.Addr().String()
}
