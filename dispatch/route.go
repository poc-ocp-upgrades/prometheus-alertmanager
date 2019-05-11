package dispatch

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/types"
)

var DefaultRouteOpts = RouteOpts{GroupWait: 30 * time.Second, GroupInterval: 5 * time.Minute, RepeatInterval: 4 * time.Hour, GroupBy: map[model.LabelName]struct{}{}, GroupByAll: false}

type Route struct {
	parent		*Route
	RouteOpts	RouteOpts
	Matchers	types.Matchers
	Continue	bool
	Routes		[]*Route
}

func NewRoute(cr *config.Route, parent *Route) *Route {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := DefaultRouteOpts
	if parent != nil {
		opts = parent.RouteOpts
	}
	if cr.Receiver != "" {
		opts.Receiver = cr.Receiver
	}
	if cr.GroupBy != nil {
		opts.GroupBy = map[model.LabelName]struct{}{}
		for _, ln := range cr.GroupBy {
			opts.GroupBy[ln] = struct{}{}
		}
	}
	opts.GroupByAll = cr.GroupByAll
	if cr.GroupWait != nil {
		opts.GroupWait = time.Duration(*cr.GroupWait)
	}
	if cr.GroupInterval != nil {
		opts.GroupInterval = time.Duration(*cr.GroupInterval)
	}
	if cr.RepeatInterval != nil {
		opts.RepeatInterval = time.Duration(*cr.RepeatInterval)
	}
	var matchers types.Matchers
	for ln, lv := range cr.Match {
		matchers = append(matchers, types.NewMatcher(model.LabelName(ln), lv))
	}
	for ln, lv := range cr.MatchRE {
		matchers = append(matchers, types.NewRegexMatcher(model.LabelName(ln), lv.Regexp))
	}
	sort.Sort(matchers)
	route := &Route{parent: parent, RouteOpts: opts, Matchers: matchers, Continue: cr.Continue}
	route.Routes = NewRoutes(cr.Routes, route)
	return route
}
func NewRoutes(croutes []*config.Route, parent *Route) []*Route {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := []*Route{}
	for _, cr := range croutes {
		res = append(res, NewRoute(cr, parent))
	}
	return res
}
func (r *Route) Match(lset model.LabelSet) []*Route {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !r.Matchers.Match(lset) {
		return nil
	}
	var all []*Route
	for _, cr := range r.Routes {
		matches := cr.Match(lset)
		all = append(all, matches...)
		if matches != nil && !cr.Continue {
			break
		}
	}
	if len(all) == 0 {
		all = append(all, r)
	}
	return all
}
func (r *Route) Key() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := make([]byte, 0, 1024)
	if r.parent != nil {
		b = append(b, r.parent.Key()...)
		b = append(b, '/')
	}
	return string(append(b, r.Matchers.String()...))
}

type RouteOpts struct {
	Receiver		string
	GroupBy			map[model.LabelName]struct{}
	GroupByAll		bool
	GroupWait		time.Duration
	GroupInterval	time.Duration
	RepeatInterval	time.Duration
}

func (ro *RouteOpts) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var labels []model.LabelName
	for ln := range ro.GroupBy {
		labels = append(labels, ln)
	}
	return fmt.Sprintf("<RouteOpts send_to:%q group_by:%q group_by_all:%t timers:%q|%q>", ro.Receiver, labels, ro.GroupByAll, ro.GroupWait, ro.GroupInterval)
}
func (ro *RouteOpts) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := struct {
		Receiver		string				`json:"receiver"`
		GroupBy			model.LabelNames	`json:"groupBy"`
		GroupByAll		bool				`json:"groupByAll"`
		GroupWait		time.Duration		`json:"groupWait"`
		GroupInterval	time.Duration		`json:"groupInterval"`
		RepeatInterval	time.Duration		`json:"repeatInterval"`
	}{Receiver: ro.Receiver, GroupByAll: ro.GroupByAll, GroupWait: ro.GroupWait, GroupInterval: ro.GroupInterval, RepeatInterval: ro.RepeatInterval}
	for ln := range ro.GroupBy {
		v.GroupBy = append(v.GroupBy, ln)
	}
	return json.Marshal(&v)
}
