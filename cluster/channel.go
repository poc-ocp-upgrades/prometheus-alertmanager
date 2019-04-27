package cluster

import (
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/memberlist"
	"github.com/prometheus/alertmanager/cluster/clusterpb"
	"github.com/prometheus/client_golang/prometheus"
)

type Channel struct {
	key					string
	send					func([]byte)
	peers					func() []*memberlist.Node
	sendOversize				func(*memberlist.Node, []byte) error
	msgc					chan []byte
	logger					log.Logger
	oversizeGossipMessageFailureTotal	prometheus.Counter
	oversizeGossipMessageDroppedTotal	prometheus.Counter
	oversizeGossipMessageSentTotal		prometheus.Counter
	oversizeGossipDuration			prometheus.Histogram
}

func NewChannel(key string, send func([]byte), peers func() []*memberlist.Node, sendOversize func(*memberlist.Node, []byte) error, logger log.Logger, stopc chan struct{}, reg prometheus.Registerer) *Channel {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oversizeGossipMessageFailureTotal := prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_oversized_gossip_message_failure_total", Help: "Number of oversized gossip message sends that failed.", ConstLabels: prometheus.Labels{"key": key}})
	oversizeGossipMessageSentTotal := prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_oversized_gossip_message_sent_total", Help: "Number of oversized gossip message sent.", ConstLabels: prometheus.Labels{"key": key}})
	oversizeGossipMessageDroppedTotal := prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_oversized_gossip_message_dropped_total", Help: "Number of oversized gossip messages that were dropped due to a full message queue.", ConstLabels: prometheus.Labels{"key": key}})
	oversizeGossipDuration := prometheus.NewHistogram(prometheus.HistogramOpts{Name: "alertmanager_oversize_gossip_message_duration_seconds", Help: "Duration of oversized gossip message requests.", ConstLabels: prometheus.Labels{"key": key}})
	reg.MustRegister(oversizeGossipDuration, oversizeGossipMessageFailureTotal, oversizeGossipMessageDroppedTotal, oversizeGossipMessageSentTotal)
	c := &Channel{key: key, send: send, peers: peers, logger: logger, msgc: make(chan []byte, 200), sendOversize: sendOversize, oversizeGossipMessageFailureTotal: oversizeGossipMessageFailureTotal, oversizeGossipMessageDroppedTotal: oversizeGossipMessageDroppedTotal, oversizeGossipMessageSentTotal: oversizeGossipMessageSentTotal, oversizeGossipDuration: oversizeGossipDuration}
	go c.handleOverSizedMessages(stopc)
	return c
}
func (c *Channel) handleOverSizedMessages(stopc chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	for {
		select {
		case b := <-c.msgc:
			for _, n := range c.peers() {
				wg.Add(1)
				go func(n *memberlist.Node) {
					defer wg.Done()
					c.oversizeGossipMessageSentTotal.Inc()
					start := time.Now()
					if err := c.sendOversize(n, b); err != nil {
						level.Debug(c.logger).Log("msg", "failed to send reliable", "key", c.key, "node", n, "err", err)
						c.oversizeGossipMessageFailureTotal.Inc()
						return
					}
					c.oversizeGossipDuration.Observe(time.Since(start).Seconds())
				}(n)
			}
			wg.Wait()
		case <-stopc:
			return
		}
	}
}
func (c *Channel) Broadcast(b []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := proto.Marshal(&clusterpb.Part{Key: c.key, Data: b})
	if err != nil {
		return
	}
	if OversizedMessage(b) {
		select {
		case c.msgc <- b:
		default:
			level.Debug(c.logger).Log("msg", "oversized gossip channel full")
			c.oversizeGossipMessageDroppedTotal.Inc()
		}
	} else {
		c.send(b)
	}
}
func OversizedMessage(b []byte) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b) > maxGossipPacketSize/2
}
