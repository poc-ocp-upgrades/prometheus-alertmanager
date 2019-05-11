package types

import (
	"strings"
	"sync"
	"time"
	"github.com/prometheus/common/model"
)

type AlertState string

const (
	AlertStateUnprocessed	AlertState	= "unprocessed"
	AlertStateActive		AlertState	= "active"
	AlertStateSuppressed	AlertState	= "suppressed"
)

type AlertStatus struct {
	State		AlertState	`json:"state"`
	SilencedBy	[]string	`json:"silencedBy"`
	InhibitedBy	[]string	`json:"inhibitedBy"`
}
type Marker interface {
	SetActive(alert model.Fingerprint)
	SetInhibited(alert model.Fingerprint, ids ...string)
	SetSilenced(alert model.Fingerprint, ids ...string)
	Count(...AlertState) int
	Status(model.Fingerprint) AlertStatus
	Delete(model.Fingerprint)
	Unprocessed(model.Fingerprint) bool
	Active(model.Fingerprint) bool
	Silenced(model.Fingerprint) ([]string, bool)
	Inhibited(model.Fingerprint) ([]string, bool)
}

func NewMarker() Marker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &memMarker{m: map[model.Fingerprint]*AlertStatus{}}
}

type memMarker struct {
	m	map[model.Fingerprint]*AlertStatus
	mtx	sync.RWMutex
}

func (m *memMarker) Count(states ...AlertState) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	count := 0
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if len(states) == 0 {
		count = len(m.m)
	} else {
		for _, status := range m.m {
			for _, state := range states {
				if status.State == state {
					count++
				}
			}
		}
	}
	return count
}
func (m *memMarker) SetSilenced(alert model.Fingerprint, ids ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.mtx.Lock()
	s, found := m.m[alert]
	if !found {
		s = &AlertStatus{}
		m.m[alert] = s
	}
	if len(ids) == 0 && len(s.InhibitedBy) == 0 {
		m.mtx.Unlock()
		m.SetActive(alert)
		return
	}
	s.State = AlertStateSuppressed
	s.SilencedBy = ids
	m.mtx.Unlock()
}
func (m *memMarker) SetInhibited(alert model.Fingerprint, ids ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.mtx.Lock()
	s, found := m.m[alert]
	if !found {
		s = &AlertStatus{}
		m.m[alert] = s
	}
	if len(ids) == 0 && len(s.SilencedBy) == 0 {
		m.mtx.Unlock()
		m.SetActive(alert)
		return
	}
	s.State = AlertStateSuppressed
	s.InhibitedBy = ids
	m.mtx.Unlock()
}
func (m *memMarker) SetActive(alert model.Fingerprint) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.mtx.Lock()
	defer m.mtx.Unlock()
	s, found := m.m[alert]
	if !found {
		s = &AlertStatus{SilencedBy: []string{}, InhibitedBy: []string{}}
		m.m[alert] = s
	}
	s.State = AlertStateActive
	s.SilencedBy = []string{}
	s.InhibitedBy = []string{}
}
func (m *memMarker) Status(alert model.Fingerprint) AlertStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	s, found := m.m[alert]
	if !found {
		s = &AlertStatus{State: AlertStateUnprocessed, SilencedBy: []string{}, InhibitedBy: []string{}}
	}
	return *s
}
func (m *memMarker) Delete(alert model.Fingerprint) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.m, alert)
}
func (m *memMarker) Unprocessed(alert model.Fingerprint) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Status(alert).State == AlertStateUnprocessed
}
func (m *memMarker) Active(alert model.Fingerprint) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Status(alert).State == AlertStateActive
}
func (m *memMarker) Inhibited(alert model.Fingerprint) ([]string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := m.Status(alert)
	return s.InhibitedBy, s.State == AlertStateSuppressed && len(s.InhibitedBy) > 0
}
func (m *memMarker) Silenced(alert model.Fingerprint) ([]string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := m.Status(alert)
	return s.SilencedBy, s.State == AlertStateSuppressed && len(s.SilencedBy) > 0
}

type MultiError struct {
	mtx		sync.Mutex
	errors	[]error
}

func (e *MultiError) Add(err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mtx.Lock()
	defer e.mtx.Unlock()
	e.errors = append(e.errors, err)
}
func (e *MultiError) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mtx.Lock()
	defer e.mtx.Unlock()
	return len(e.errors)
}
func (e *MultiError) Errors() []error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mtx.Lock()
	defer e.mtx.Unlock()
	return append(make([]error, 0, len(e.errors)), e.errors...)
}
func (e *MultiError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mtx.Lock()
	defer e.mtx.Unlock()
	es := make([]string, 0, len(e.errors))
	for _, err := range e.errors {
		es = append(es, err.Error())
	}
	return strings.Join(es, "; ")
}

type Alert struct {
	model.Alert
	UpdatedAt	time.Time
	Timeout		bool
}
type AlertSlice []*Alert

func (as AlertSlice) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, overrideKey := range [...]model.LabelName{"job", "instance"} {
		iVal, iOk := as[i].Labels[overrideKey]
		jVal, jOk := as[j].Labels[overrideKey]
		if !iOk && !jOk {
			continue
		}
		if !iOk {
			return false
		}
		if !jOk {
			return true
		}
		if iVal != jVal {
			return iVal < jVal
		}
	}
	return as[i].Labels.Before(as[j].Labels)
}
func (as AlertSlice) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	as[i], as[j] = as[j], as[i]
}
func (as AlertSlice) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(as)
}
func Alerts(alerts ...*Alert) model.Alerts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := make(model.Alerts, 0, len(alerts))
	for _, a := range alerts {
		v := a.Alert
		if !a.Resolved() {
			v.EndsAt = time.Time{}
		}
		res = append(res, &v)
	}
	return res
}
func (a *Alert) Merge(o *Alert) *Alert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.UpdatedAt.Before(a.UpdatedAt) {
		return o.Merge(a)
	}
	res := *o
	if a.StartsAt.Before(o.StartsAt) {
		res.StartsAt = a.StartsAt
	}
	if o.Resolved() {
		if a.Resolved() && a.EndsAt.After(o.EndsAt) {
			res.EndsAt = a.EndsAt
		}
	} else {
		if a.EndsAt.After(o.EndsAt) && !a.Timeout {
			res.EndsAt = a.EndsAt
		}
	}
	return &res
}

type Muter interface{ Mutes(model.LabelSet) bool }
type MuteFunc func(model.LabelSet) bool

func (f MuteFunc) Mutes(lset model.LabelSet) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(lset)
}

type Silence struct {
	ID			string			`json:"id"`
	Matchers	Matchers		`json:"matchers"`
	StartsAt	time.Time		`json:"startsAt"`
	EndsAt		time.Time		`json:"endsAt"`
	UpdatedAt	time.Time		`json:"updatedAt"`
	CreatedBy	string			`json:"createdBy"`
	Comment		string			`json:"comment,omitempty"`
	Status		SilenceStatus	`json:"status"`
}

func (s *Silence) Expired() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.StartsAt.Equal(s.EndsAt)
}

type SilenceStatus struct {
	State SilenceState `json:"state"`
}
type SilenceState string

const (
	SilenceStateExpired	SilenceState	= "expired"
	SilenceStateActive	SilenceState	= "active"
	SilenceStatePending	SilenceState	= "pending"
)

func CalcSilenceState(start, end time.Time) SilenceState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	current := time.Now()
	if current.Before(start) {
		return SilenceStatePending
	}
	if current.Before(end) {
		return SilenceStateActive
	}
	return SilenceStateExpired
}
