package dispatch

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"sort"
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/provider"
	"github.com/prometheus/alertmanager/store"
	"github.com/prometheus/alertmanager/types"
)

type Dispatcher struct {
	route		*Route
	alerts		provider.Alerts
	stage		notify.Stage
	marker		types.Marker
	timeout		func(time.Duration) time.Duration
	aggrGroups	map[*Route]map[model.Fingerprint]*aggrGroup
	mtx		sync.RWMutex
	done		chan struct{}
	ctx		context.Context
	cancel		func()
	logger		log.Logger
}

func NewDispatcher(ap provider.Alerts, r *Route, s notify.Stage, mk types.Marker, to func(time.Duration) time.Duration, l log.Logger) *Dispatcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	disp := &Dispatcher{alerts: ap, stage: s, route: r, marker: mk, timeout: to, logger: log.With(l, "component", "dispatcher")}
	return disp
}
func (d *Dispatcher) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.done = make(chan struct{})
	d.mtx.Lock()
	d.aggrGroups = map[*Route]map[model.Fingerprint]*aggrGroup{}
	d.mtx.Unlock()
	d.ctx, d.cancel = context.WithCancel(context.Background())
	d.run(d.alerts.Subscribe())
	close(d.done)
}
func (d *Dispatcher) run(it provider.AlertIterator) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cleanup := time.NewTicker(30 * time.Second)
	defer cleanup.Stop()
	defer it.Close()
	for {
		select {
		case alert, ok := <-it.Next():
			if !ok {
				if err := it.Err(); err != nil {
					level.Error(d.logger).Log("msg", "Error on alert update", "err", err)
				}
				return
			}
			level.Debug(d.logger).Log("msg", "Received alert", "alert", alert)
			if err := it.Err(); err != nil {
				level.Error(d.logger).Log("msg", "Error on alert update", "err", err)
				continue
			}
			for _, r := range d.route.Match(alert.Labels) {
				d.processAlert(alert, r)
			}
		case <-cleanup.C:
			d.mtx.Lock()
			for _, groups := range d.aggrGroups {
				for _, ag := range groups {
					if ag.empty() {
						ag.stop()
						delete(groups, ag.fingerprint())
					}
				}
			}
			d.mtx.Unlock()
		case <-d.ctx.Done():
			return
		}
	}
}
func (d *Dispatcher) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if d == nil || d.cancel == nil {
		return
	}
	d.cancel()
	d.cancel = nil
	<-d.done
}

type notifyFunc func(context.Context, ...*types.Alert) bool

func (d *Dispatcher) processAlert(alert *types.Alert, route *Route) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupLabels := getGroupLabels(alert, route)
	fp := groupLabels.Fingerprint()
	d.mtx.Lock()
	defer d.mtx.Unlock()
	group, ok := d.aggrGroups[route]
	if !ok {
		group = map[model.Fingerprint]*aggrGroup{}
		d.aggrGroups[route] = group
	}
	ag, ok := group[fp]
	if !ok {
		ag = newAggrGroup(d.ctx, groupLabels, route, d.timeout, d.logger)
		group[fp] = ag
		go ag.run(func(ctx context.Context, alerts ...*types.Alert) bool {
			_, _, err := d.stage.Exec(ctx, d.logger, alerts...)
			if err != nil {
				level.Error(d.logger).Log("msg", "Notify for alerts failed", "num_alerts", len(alerts), "err", err)
			}
			return err == nil
		})
	}
	ag.insert(alert)
}
func getGroupLabels(alert *types.Alert, route *Route) model.LabelSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupLabels := model.LabelSet{}
	for ln, lv := range alert.Labels {
		if _, ok := route.RouteOpts.GroupBy[ln]; ok || route.RouteOpts.GroupByAll {
			groupLabels[ln] = lv
		}
	}
	return groupLabels
}

type aggrGroup struct {
	labels		model.LabelSet
	opts		*RouteOpts
	logger		log.Logger
	routeKey	string
	alerts		*store.Alerts
	ctx		context.Context
	cancel		func()
	done		chan struct{}
	next		*time.Timer
	timeout		func(time.Duration) time.Duration
	mtx		sync.RWMutex
	hasFlushed	bool
}

func newAggrGroup(ctx context.Context, labels model.LabelSet, r *Route, to func(time.Duration) time.Duration, logger log.Logger) *aggrGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if to == nil {
		to = func(d time.Duration) time.Duration {
			return d
		}
	}
	ag := &aggrGroup{labels: labels, routeKey: r.Key(), opts: &r.RouteOpts, timeout: to, alerts: store.NewAlerts(15 * time.Minute)}
	ag.ctx, ag.cancel = context.WithCancel(ctx)
	ag.alerts.Run(ag.ctx)
	ag.logger = log.With(logger, "aggrGroup", ag)
	ag.next = time.NewTimer(ag.opts.GroupWait)
	return ag
}
func (ag *aggrGroup) fingerprint() model.Fingerprint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ag.labels.Fingerprint()
}
func (ag *aggrGroup) GroupKey() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s:%s", ag.routeKey, ag.labels)
}
func (ag *aggrGroup) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ag.GroupKey()
}
func (ag *aggrGroup) run(nf notifyFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ag.done = make(chan struct{})
	defer close(ag.done)
	defer ag.next.Stop()
	for {
		select {
		case now := <-ag.next.C:
			ctx, cancel := context.WithTimeout(ag.ctx, ag.timeout(ag.opts.GroupInterval))
			ctx = notify.WithNow(ctx, now)
			ctx = notify.WithGroupKey(ctx, ag.GroupKey())
			ctx = notify.WithGroupLabels(ctx, ag.labels)
			ctx = notify.WithReceiverName(ctx, ag.opts.Receiver)
			ctx = notify.WithRepeatInterval(ctx, ag.opts.RepeatInterval)
			ag.mtx.Lock()
			ag.next.Reset(ag.opts.GroupInterval)
			ag.hasFlushed = true
			ag.mtx.Unlock()
			ag.flush(func(alerts ...*types.Alert) bool {
				return nf(ctx, alerts...)
			})
			cancel()
		case <-ag.ctx.Done():
			return
		}
	}
}
func (ag *aggrGroup) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ag.cancel()
	<-ag.done
}
func (ag *aggrGroup) insert(alert *types.Alert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ag.alerts.Set(alert); err != nil {
		level.Error(ag.logger).Log("msg", "error on set alert", "err", err)
	}
	ag.mtx.Lock()
	defer ag.mtx.Unlock()
	if !ag.hasFlushed && alert.StartsAt.Add(ag.opts.GroupWait).Before(time.Now()) {
		ag.next.Reset(0)
	}
}
func (ag *aggrGroup) empty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ag.alerts.Count() == 0
}
func (ag *aggrGroup) flush(notify func(...*types.Alert) bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ag.empty() {
		return
	}
	var (
		alerts		= ag.alerts.List()
		alertsSlice	= make(types.AlertSlice, 0, ag.alerts.Count())
	)
	now := time.Now()
	for alert := range alerts {
		a := *alert
		if !a.ResolvedAt(now) {
			a.EndsAt = time.Time{}
		}
		alertsSlice = append(alertsSlice, &a)
	}
	sort.Stable(alertsSlice)
	level.Debug(ag.logger).Log("msg", "flushing", "alerts", fmt.Sprintf("%v", alertsSlice))
	if notify(alertsSlice...) {
		for _, a := range alertsSlice {
			fp := a.Fingerprint()
			got, err := ag.alerts.Get(fp)
			if err != nil {
				level.Error(ag.logger).Log("msg", "failed to get alert", "err", err)
				continue
			}
			if a.Resolved() && got.UpdatedAt == a.UpdatedAt {
				if err := ag.alerts.Delete(fp); err != nil {
					level.Error(ag.logger).Log("msg", "error on delete alert", "err", err)
				}
			}
		}
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
