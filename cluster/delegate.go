package cluster

import (
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/memberlist"
	"github.com/prometheus/alertmanager/cluster/clusterpb"
	"github.com/prometheus/client_golang/prometheus"
)

const maxQueueSize = 4096

type delegate struct {
	*Peer
	logger			log.Logger
	bcast			*memberlist.TransmitLimitedQueue
	messagesReceived	*prometheus.CounterVec
	messagesReceivedSize	*prometheus.CounterVec
	messagesSent		*prometheus.CounterVec
	messagesSentSize	*prometheus.CounterVec
	messagesPruned		prometheus.Counter
}

func newDelegate(l log.Logger, reg prometheus.Registerer, p *Peer, retransmit int) *delegate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcast := &memberlist.TransmitLimitedQueue{NumNodes: p.ClusterSize, RetransmitMult: retransmit}
	messagesReceived := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "alertmanager_cluster_messages_received_total", Help: "Total number of cluster messsages received."}, []string{"msg_type"})
	messagesReceivedSize := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "alertmanager_cluster_messages_received_size_total", Help: "Total size of cluster messages received."}, []string{"msg_type"})
	messagesSent := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "alertmanager_cluster_messages_sent_total", Help: "Total number of cluster messsages sent."}, []string{"msg_type"})
	messagesSentSize := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "alertmanager_cluster_messages_sent_size_total", Help: "Total size of cluster messages sent."}, []string{"msg_type"})
	messagesPruned := prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_messages_pruned_total", Help: "Total number of cluster messsages pruned."})
	gossipClusterMembers := prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "alertmanager_cluster_members", Help: "Number indicating current number of members in cluster."}, func() float64 {
		return float64(p.ClusterSize())
	})
	peerPosition := prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "alertmanager_peer_position", Help: "Position the Alertmanager instance believes it's in. The position determines a peer's behavior in the cluster."}, func() float64 {
		return float64(p.Position())
	})
	healthScore := prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "alertmanager_cluster_health_score", Help: "Health score of the cluster. Lower values are better and zero means 'totally healthy'."}, func() float64 {
		return float64(p.mlist.GetHealthScore())
	})
	messagesQueued := prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "alertmanager_cluster_messages_queued", Help: "Number of cluster messsages which are queued."}, func() float64 {
		return float64(bcast.NumQueued())
	})
	messagesReceived.WithLabelValues("full_state")
	messagesReceivedSize.WithLabelValues("full_state")
	messagesReceived.WithLabelValues("update")
	messagesReceivedSize.WithLabelValues("update")
	messagesSent.WithLabelValues("full_state")
	messagesSentSize.WithLabelValues("full_state")
	messagesSent.WithLabelValues("update")
	messagesSentSize.WithLabelValues("update")
	reg.MustRegister(messagesReceived, messagesReceivedSize, messagesSent, messagesSentSize, gossipClusterMembers, peerPosition, healthScore, messagesQueued, messagesPruned)
	d := &delegate{logger: l, Peer: p, bcast: bcast, messagesReceived: messagesReceived, messagesReceivedSize: messagesReceivedSize, messagesSent: messagesSent, messagesSentSize: messagesSentSize, messagesPruned: messagesPruned}
	go d.handleQueueDepth()
	return d
}
func (d *delegate) NodeMeta(limit int) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte{}
}
func (d *delegate) NotifyMsg(b []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.messagesReceived.WithLabelValues("update").Inc()
	d.messagesReceivedSize.WithLabelValues("update").Add(float64(len(b)))
	var p clusterpb.Part
	if err := proto.Unmarshal(b, &p); err != nil {
		level.Warn(d.logger).Log("msg", "decode broadcast", "err", err)
		return
	}
	s, ok := d.states[p.Key]
	if !ok {
		return
	}
	if err := s.Merge(p.Data); err != nil {
		level.Warn(d.logger).Log("msg", "merge broadcast", "err", err, "key", p.Key)
		return
	}
}
func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	msgs := d.bcast.GetBroadcasts(overhead, limit)
	d.messagesSent.WithLabelValues("update").Add(float64(len(msgs)))
	for _, m := range msgs {
		d.messagesSentSize.WithLabelValues("update").Add(float64(len(m)))
	}
	return msgs
}
func (d *delegate) LocalState(_ bool) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	all := &clusterpb.FullState{Parts: make([]clusterpb.Part, 0, len(d.states))}
	for key, s := range d.states {
		b, err := s.MarshalBinary()
		if err != nil {
			level.Warn(d.logger).Log("msg", "encode local state", "err", err, "key", key)
			return nil
		}
		all.Parts = append(all.Parts, clusterpb.Part{Key: key, Data: b})
	}
	b, err := proto.Marshal(all)
	if err != nil {
		level.Warn(d.logger).Log("msg", "encode local state", "err", err)
		return nil
	}
	d.messagesSent.WithLabelValues("full_state").Inc()
	d.messagesSentSize.WithLabelValues("full_state").Add(float64(len(b)))
	return b
}
func (d *delegate) MergeRemoteState(buf []byte, _ bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.messagesReceived.WithLabelValues("full_state").Inc()
	d.messagesReceivedSize.WithLabelValues("full_state").Add(float64(len(buf)))
	var fs clusterpb.FullState
	if err := proto.Unmarshal(buf, &fs); err != nil {
		level.Warn(d.logger).Log("msg", "merge remote state", "err", err)
		return
	}
	d.mtx.RLock()
	defer d.mtx.RUnlock()
	for _, p := range fs.Parts {
		s, ok := d.states[p.Key]
		if !ok {
			level.Warn(d.logger).Log("received", "unknown state key", "len", len(buf), "key", p.Key)
			continue
		}
		if err := s.Merge(p.Data); err != nil {
			level.Warn(d.logger).Log("msg", "merge remote state", "err", err, "key", p.Key)
			return
		}
	}
}
func (d *delegate) NotifyJoin(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	level.Debug(d.logger).Log("received", "NotifyJoin", "node", n.Name, "addr", n.Address())
	d.Peer.peerJoin(n)
}
func (d *delegate) NotifyLeave(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	level.Debug(d.logger).Log("received", "NotifyLeave", "node", n.Name, "addr", n.Address())
	d.Peer.peerLeave(n)
}
func (d *delegate) NotifyUpdate(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	level.Debug(d.logger).Log("received", "NotifyUpdate", "node", n.Name, "addr", n.Address())
	d.Peer.peerUpdate(n)
}
func (d *delegate) handleQueueDepth() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case <-d.stopc:
			return
		case <-time.After(15 * time.Minute):
			n := d.bcast.NumQueued()
			if n > maxQueueSize {
				level.Warn(d.logger).Log("msg", "dropping messages because too many are queued", "current", n, "limit", maxQueueSize)
				d.bcast.Prune(maxQueueSize)
				d.messagesPruned.Add(float64(n - maxQueueSize))
			}
		}
	}
}
