package notify

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
	"github.com/cenkalti/backoff"
	"github.com/cespare/xxhash"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/cluster"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/nflog"
	"github.com/prometheus/alertmanager/nflog/nflogpb"
	"github.com/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/alertmanager/types"
)

var (
	numNotifications		= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "alertmanager", Name: "notifications_total", Help: "The total number of attempted notifications."}, []string{"integration"})
	numFailedNotifications		= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "alertmanager", Name: "notifications_failed_total", Help: "The total number of failed notifications."}, []string{"integration"})
	notificationLatencySeconds	= prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "alertmanager", Name: "notification_latency_seconds", Help: "The latency of notifications in seconds.", Buckets: []float64{1, 5, 10, 15, 20}}, []string{"integration"})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	numNotifications.WithLabelValues("email")
	numNotifications.WithLabelValues("hipchat")
	numNotifications.WithLabelValues("pagerduty")
	numNotifications.WithLabelValues("wechat")
	numNotifications.WithLabelValues("pushover")
	numNotifications.WithLabelValues("slack")
	numNotifications.WithLabelValues("opsgenie")
	numNotifications.WithLabelValues("webhook")
	numNotifications.WithLabelValues("victorops")
	numFailedNotifications.WithLabelValues("email")
	numFailedNotifications.WithLabelValues("hipchat")
	numFailedNotifications.WithLabelValues("pagerduty")
	numFailedNotifications.WithLabelValues("wechat")
	numFailedNotifications.WithLabelValues("pushover")
	numFailedNotifications.WithLabelValues("slack")
	numFailedNotifications.WithLabelValues("opsgenie")
	numFailedNotifications.WithLabelValues("webhook")
	numFailedNotifications.WithLabelValues("victorops")
	notificationLatencySeconds.WithLabelValues("email")
	notificationLatencySeconds.WithLabelValues("hipchat")
	notificationLatencySeconds.WithLabelValues("pagerduty")
	notificationLatencySeconds.WithLabelValues("wechat")
	notificationLatencySeconds.WithLabelValues("pushover")
	notificationLatencySeconds.WithLabelValues("slack")
	notificationLatencySeconds.WithLabelValues("opsgenie")
	notificationLatencySeconds.WithLabelValues("webhook")
	notificationLatencySeconds.WithLabelValues("victorops")
	prometheus.MustRegister(numNotifications)
	prometheus.MustRegister(numFailedNotifications)
	prometheus.MustRegister(notificationLatencySeconds)
}

type notifierConfig interface{ SendResolved() bool }

const MinTimeout = 10 * time.Second

type notifyKey int

const (
	keyReceiverName	notifyKey	= iota
	keyRepeatInterval
	keyGroupLabels
	keyGroupKey
	keyFiringAlerts
	keyResolvedAlerts
	keyNow
)

func WithReceiverName(ctx context.Context, rcv string) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyReceiverName, rcv)
}
func WithGroupKey(ctx context.Context, s string) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyGroupKey, s)
}
func WithFiringAlerts(ctx context.Context, alerts []uint64) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyFiringAlerts, alerts)
}
func WithResolvedAlerts(ctx context.Context, alerts []uint64) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyResolvedAlerts, alerts)
}
func WithGroupLabels(ctx context.Context, lset model.LabelSet) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyGroupLabels, lset)
}
func WithNow(ctx context.Context, t time.Time) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyNow, t)
}
func WithRepeatInterval(ctx context.Context, t time.Duration) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return context.WithValue(ctx, keyRepeatInterval, t)
}
func RepeatInterval(ctx context.Context) (time.Duration, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyRepeatInterval).(time.Duration)
	return v, ok
}
func ReceiverName(ctx context.Context) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyReceiverName).(string)
	return v, ok
}
func receiverName(ctx context.Context, l log.Logger) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	recv, ok := ReceiverName(ctx)
	if !ok {
		level.Error(l).Log("msg", "Missing receiver")
	}
	return recv
}
func GroupKey(ctx context.Context) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyGroupKey).(string)
	return v, ok
}
func groupLabels(ctx context.Context, l log.Logger) model.LabelSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupLabels, ok := GroupLabels(ctx)
	if !ok {
		level.Error(l).Log("msg", "Missing group labels")
	}
	return groupLabels
}
func GroupLabels(ctx context.Context) (model.LabelSet, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyGroupLabels).(model.LabelSet)
	return v, ok
}
func Now(ctx context.Context) (time.Time, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyNow).(time.Time)
	return v, ok
}
func FiringAlerts(ctx context.Context) ([]uint64, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyFiringAlerts).([]uint64)
	return v, ok
}
func ResolvedAlerts(ctx context.Context) ([]uint64, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v, ok := ctx.Value(keyResolvedAlerts).([]uint64)
	return v, ok
}

type Stage interface {
	Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error)
}
type StageFunc func(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error)

func (f StageFunc) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(ctx, l, alerts...)
}

type NotificationLog interface {
	Log(r *nflogpb.Receiver, gkey string, firingAlerts, resolvedAlerts []uint64) error
	Query(params ...nflog.QueryParam) ([]*nflogpb.Entry, error)
}

func BuildPipeline(confs []*config.Receiver, tmpl *template.Template, wait func() time.Duration, muter types.Muter, silences *silence.Silences, notificationLog NotificationLog, marker types.Marker, peer *cluster.Peer, logger log.Logger) RoutingStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs := RoutingStage{}
	ms := NewGossipSettleStage(peer)
	is := NewInhibitStage(muter)
	ss := NewSilenceStage(silences, marker)
	for _, rc := range confs {
		rs[rc.Name] = MultiStage{ms, is, ss, createStage(rc, tmpl, wait, notificationLog, logger)}
	}
	return rs
}
func createStage(rc *config.Receiver, tmpl *template.Template, wait func() time.Duration, notificationLog NotificationLog, logger log.Logger) Stage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var fs FanoutStage
	for _, i := range BuildReceiverIntegrations(rc, tmpl, logger) {
		recv := &nflogpb.Receiver{GroupName: rc.Name, Integration: i.name, Idx: uint32(i.idx)}
		var s MultiStage
		s = append(s, NewWaitStage(wait))
		s = append(s, NewDedupStage(i, notificationLog, recv))
		s = append(s, NewRetryStage(i, rc.Name))
		s = append(s, NewSetNotifiesStage(notificationLog, recv))
		fs = append(fs, s)
	}
	return fs
}

type RoutingStage map[string]Stage

func (rs RoutingStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	receiver, ok := ReceiverName(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("receiver missing")
	}
	s, ok := rs[receiver]
	if !ok {
		return ctx, nil, fmt.Errorf("stage for receiver missing")
	}
	return s.Exec(ctx, l, alerts...)
}

type MultiStage []Stage

func (ms MultiStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	for _, s := range ms {
		if len(alerts) == 0 {
			return ctx, nil, nil
		}
		ctx, alerts, err = s.Exec(ctx, l, alerts...)
		if err != nil {
			return ctx, nil, err
		}
	}
	return ctx, alerts, nil
}

type FanoutStage []Stage

func (fs FanoutStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		wg	sync.WaitGroup
		me	types.MultiError
	)
	wg.Add(len(fs))
	for _, s := range fs {
		go func(s Stage) {
			if _, _, err := s.Exec(ctx, l, alerts...); err != nil {
				me.Add(err)
				level.Error(l).Log("msg", "Error on notify", "err", err)
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
	if me.Len() > 0 {
		return ctx, alerts, &me
	}
	return ctx, alerts, nil
}

type GossipSettleStage struct{ peer *cluster.Peer }

func NewGossipSettleStage(p *cluster.Peer) *GossipSettleStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GossipSettleStage{peer: p}
}
func (n *GossipSettleStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.peer != nil {
		n.peer.WaitReady()
	}
	return ctx, alerts, nil
}

type InhibitStage struct{ muter types.Muter }

func NewInhibitStage(m types.Muter) *InhibitStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &InhibitStage{muter: m}
}
func (n *InhibitStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var filtered []*types.Alert
	for _, a := range alerts {
		if !n.muter.Mutes(a.Labels) {
			filtered = append(filtered, a)
		}
	}
	return ctx, filtered, nil
}

type SilenceStage struct {
	silences	*silence.Silences
	marker		types.Marker
}

func NewSilenceStage(s *silence.Silences, mk types.Marker) *SilenceStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SilenceStage{silences: s, marker: mk}
}
func (n *SilenceStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var filtered []*types.Alert
	for _, a := range alerts {
		sils, err := n.silences.Query(silence.QState(types.SilenceStateActive), silence.QMatches(a.Labels))
		if err != nil {
			level.Error(l).Log("msg", "Querying silences failed", "err", err)
		}
		if len(sils) == 0 {
			filtered = append(filtered, a)
			n.marker.SetSilenced(a.Labels.Fingerprint())
		} else {
			ids := make([]string, len(sils))
			for i, s := range sils {
				ids[i] = s.Id
			}
			n.marker.SetSilenced(a.Labels.Fingerprint(), ids...)
		}
	}
	return ctx, filtered, nil
}

type WaitStage struct{ wait func() time.Duration }

func NewWaitStage(wait func() time.Duration) *WaitStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &WaitStage{wait: wait}
}
func (ws *WaitStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-time.After(ws.wait()):
	case <-ctx.Done():
		return ctx, nil, ctx.Err()
	}
	return ctx, alerts, nil
}

type DedupStage struct {
	nflog	NotificationLog
	recv	*nflogpb.Receiver
	conf	notifierConfig
	now	func() time.Time
	hash	func(*types.Alert) uint64
}

func NewDedupStage(i Integration, l NotificationLog, recv *nflogpb.Receiver) *DedupStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DedupStage{nflog: l, recv: recv, conf: i.conf, now: utcNow, hash: hashAlert}
}
func utcNow() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Now().UTC()
}

var hashBuffers = sync.Pool{}

func getHashBuffer() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := hashBuffers.Get()
	if b == nil {
		return make([]byte, 0, 1024)
	}
	return b.([]byte)
}
func putHashBuffer(b []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b = b[:0]
	hashBuffers.Put(b)
}
func hashAlert(a *types.Alert) uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const sep = '\xff'
	b := getHashBuffer()
	defer putHashBuffer(b)
	names := make(model.LabelNames, 0, len(a.Labels))
	for ln := range a.Labels {
		names = append(names, ln)
	}
	sort.Sort(names)
	for _, ln := range names {
		b = append(b, string(ln)...)
		b = append(b, sep)
		b = append(b, string(a.Labels[ln])...)
		b = append(b, sep)
	}
	hash := xxhash.Sum64(b)
	return hash
}
func (n *DedupStage) needsUpdate(entry *nflogpb.Entry, firing, resolved map[uint64]struct{}, repeat time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if entry == nil {
		return len(firing) > 0
	}
	if !entry.IsFiringSubset(firing) {
		return true
	}
	if len(firing) == 0 {
		return len(entry.FiringAlerts) > 0
	}
	if n.conf.SendResolved() && !entry.IsResolvedSubset(resolved) {
		return true
	}
	return entry.Timestamp.Before(n.now().Add(-repeat))
}
func (n *DedupStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gkey, ok := GroupKey(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("group key missing")
	}
	repeatInterval, ok := RepeatInterval(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("repeat interval missing")
	}
	firingSet := map[uint64]struct{}{}
	resolvedSet := map[uint64]struct{}{}
	firing := []uint64{}
	resolved := []uint64{}
	var hash uint64
	for _, a := range alerts {
		hash = n.hash(a)
		if a.Resolved() {
			resolved = append(resolved, hash)
			resolvedSet[hash] = struct{}{}
		} else {
			firing = append(firing, hash)
			firingSet[hash] = struct{}{}
		}
	}
	ctx = WithFiringAlerts(ctx, firing)
	ctx = WithResolvedAlerts(ctx, resolved)
	entries, err := n.nflog.Query(nflog.QGroupKey(gkey), nflog.QReceiver(n.recv))
	if err != nil && err != nflog.ErrNotFound {
		return ctx, nil, err
	}
	var entry *nflogpb.Entry
	switch len(entries) {
	case 0:
	case 1:
		entry = entries[0]
	case 2:
		return ctx, nil, fmt.Errorf("unexpected entry result size %d", len(entries))
	}
	if n.needsUpdate(entry, firingSet, resolvedSet, repeatInterval) {
		return ctx, alerts, nil
	}
	return ctx, nil, nil
}

type RetryStage struct {
	integration	Integration
	groupName	string
}

func NewRetryStage(i Integration, groupName string) *RetryStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RetryStage{integration: i, groupName: groupName}
}
func (r RetryStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sent []*types.Alert
	if !r.integration.conf.SendResolved() {
		firing, ok := FiringAlerts(ctx)
		if !ok {
			return ctx, nil, fmt.Errorf("firing alerts missing")
		}
		if len(firing) == 0 {
			return ctx, alerts, nil
		}
		for _, a := range alerts {
			if a.Status() != model.AlertResolved {
				sent = append(sent, a)
			}
		}
	} else {
		sent = alerts
	}
	var (
		i	= 0
		b	= backoff.NewExponentialBackOff()
		tick	= backoff.NewTicker(b)
		iErr	error
	)
	defer tick.Stop()
	for {
		i++
		select {
		case <-ctx.Done():
			if iErr != nil {
				return ctx, nil, iErr
			}
			return ctx, nil, ctx.Err()
		default:
		}
		select {
		case <-tick.C:
			now := time.Now()
			retry, err := r.integration.Notify(ctx, sent...)
			notificationLatencySeconds.WithLabelValues(r.integration.name).Observe(time.Since(now).Seconds())
			numNotifications.WithLabelValues(r.integration.name).Inc()
			if err != nil {
				numFailedNotifications.WithLabelValues(r.integration.name).Inc()
				level.Debug(l).Log("msg", "Notify attempt failed", "attempt", i, "integration", r.integration.name, "receiver", r.groupName, "err", err)
				if !retry {
					return ctx, alerts, fmt.Errorf("cancelling notify retry for %q due to unrecoverable error: %s", r.integration.name, err)
				}
				iErr = err
			} else {
				return ctx, alerts, nil
			}
		case <-ctx.Done():
			if iErr != nil {
				return ctx, nil, iErr
			}
			return ctx, nil, ctx.Err()
		}
	}
}

type SetNotifiesStage struct {
	nflog	NotificationLog
	recv	*nflogpb.Receiver
}

func NewSetNotifiesStage(l NotificationLog, recv *nflogpb.Receiver) *SetNotifiesStage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SetNotifiesStage{nflog: l, recv: recv}
}
func (n SetNotifiesStage) Exec(ctx context.Context, l log.Logger, alerts ...*types.Alert) (context.Context, []*types.Alert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gkey, ok := GroupKey(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("group key missing")
	}
	firing, ok := FiringAlerts(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("firing alerts missing")
	}
	resolved, ok := ResolvedAlerts(ctx)
	if !ok {
		return ctx, nil, fmt.Errorf("resolved alerts missing")
	}
	return ctx, alerts, n.nflog.Log(n.recv, gkey, firing, resolved)
}
