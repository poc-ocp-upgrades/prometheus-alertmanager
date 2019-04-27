package cluster

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/hashicorp/memberlist"
	"github.com/oklog/ulid"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

type Peer struct {
	mlist				*memberlist.Memberlist
	delegate			*delegate
	resolvedPeers			[]string
	mtx				sync.RWMutex
	states				map[string]State
	stopc				chan struct{}
	readyc				chan struct{}
	peerLock			sync.RWMutex
	peers				map[string]peer
	failedPeers			[]peer
	knownPeers			[]string
	advertiseAddr			string
	failedReconnectionsCounter	prometheus.Counter
	reconnectionsCounter		prometheus.Counter
	failedRefreshCounter		prometheus.Counter
	refreshCounter			prometheus.Counter
	peerLeaveCounter		prometheus.Counter
	peerUpdateCounter		prometheus.Counter
	peerJoinCounter			prometheus.Counter
	logger				log.Logger
}
type peer struct {
	status		PeerStatus
	leaveTime	time.Time
	*memberlist.Node
}
type PeerStatus int

const (
	StatusNone	PeerStatus	= iota
	StatusAlive
	StatusFailed
)

func (s PeerStatus) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch s {
	case StatusNone:
		return "none"
	case StatusAlive:
		return "alive"
	case StatusFailed:
		return "failed"
	default:
		panic(fmt.Sprintf("unknown PeerStatus: %d", s))
	}
}

const (
	DefaultPushPullInterval		= 60 * time.Second
	DefaultGossipInterval		= 200 * time.Millisecond
	DefaultTcpTimeout		= 10 * time.Second
	DefaultProbeTimeout		= 500 * time.Millisecond
	DefaultProbeInterval		= 1 * time.Second
	DefaultReconnectInterval	= 10 * time.Second
	DefaultReconnectTimeout		= 6 * time.Hour
	DefaultRefreshInterval		= 15 * time.Second
	maxGossipPacketSize		= 1400
)

func Create(l log.Logger, reg prometheus.Registerer, bindAddr string, advertiseAddr string, knownPeers []string, waitIfEmpty bool, pushPullInterval time.Duration, gossipInterval time.Duration, tcpTimeout time.Duration, probeTimeout time.Duration, probeInterval time.Duration) (*Peer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bindHost, bindPortStr, err := net.SplitHostPort(bindAddr)
	if err != nil {
		return nil, err
	}
	bindPort, err := strconv.Atoi(bindPortStr)
	if err != nil {
		return nil, errors.Wrap(err, "invalid listen address")
	}
	var advertiseHost string
	var advertisePort int
	if advertiseAddr != "" {
		var advertisePortStr string
		advertiseHost, advertisePortStr, err = net.SplitHostPort(advertiseAddr)
		if err != nil {
			return nil, errors.Wrap(err, "invalid advertise address")
		}
		advertisePort, err = strconv.Atoi(advertisePortStr)
		if err != nil {
			return nil, errors.Wrap(err, "invalid advertise address, wrong port")
		}
	}
	resolvedPeers, err := resolvePeers(context.Background(), knownPeers, advertiseAddr, net.Resolver{}, waitIfEmpty)
	if err != nil {
		return nil, errors.Wrap(err, "resolve peers")
	}
	level.Debug(l).Log("msg", "resolved peers to following addresses", "peers", strings.Join(resolvedPeers, ","))
	addr, err := calculateAdvertiseAddress(bindHost, advertiseHost)
	if err != nil {
		level.Warn(l).Log("err", "couldn't deduce an advertise address: "+err.Error())
	} else if hasNonlocal(resolvedPeers) && isUnroutable(addr.String()) {
		level.Warn(l).Log("err", "this node advertises itself on an unroutable address", "addr", addr.String())
		level.Warn(l).Log("err", "this node will be unreachable in the cluster")
		level.Warn(l).Log("err", "provide --cluster.advertise-address as a routable IP address or hostname")
	} else if isAny(bindAddr) && advertiseHost == "" {
		level.Info(l).Log("msg", "setting advertise address explicitly", "addr", addr.String(), "port", bindPort)
		advertiseHost = addr.String()
		advertisePort = bindPort
	}
	name, err := ulid.New(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil {
		return nil, err
	}
	p := &Peer{states: map[string]State{}, stopc: make(chan struct{}), readyc: make(chan struct{}), logger: l, peers: map[string]peer{}, resolvedPeers: resolvedPeers, knownPeers: knownPeers}
	p.register(reg)
	retransmit := len(knownPeers) / 2
	if retransmit < 3 {
		retransmit = 3
	}
	p.delegate = newDelegate(l, reg, p, retransmit)
	cfg := memberlist.DefaultLANConfig()
	cfg.Name = name.String()
	cfg.BindAddr = bindHost
	cfg.BindPort = bindPort
	cfg.Delegate = p.delegate
	cfg.Events = p.delegate
	cfg.GossipInterval = gossipInterval
	cfg.PushPullInterval = pushPullInterval
	cfg.TCPTimeout = tcpTimeout
	cfg.ProbeTimeout = probeTimeout
	cfg.ProbeInterval = probeInterval
	cfg.LogOutput = &logWriter{l: l}
	cfg.GossipNodes = retransmit
	cfg.UDPBufferSize = maxGossipPacketSize
	if advertiseHost != "" {
		cfg.AdvertiseAddr = advertiseHost
		cfg.AdvertisePort = advertisePort
		p.setInitialFailed(resolvedPeers, fmt.Sprintf("%s:%d", advertiseHost, advertisePort))
	} else {
		p.setInitialFailed(resolvedPeers, bindAddr)
	}
	ml, err := memberlist.Create(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "create memberlist")
	}
	p.mlist = ml
	return p, nil
}
func (p *Peer) Join(reconnectInterval time.Duration, reconnectTimeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, err := p.mlist.Join(p.resolvedPeers)
	if err != nil {
		level.Warn(p.logger).Log("msg", "failed to join cluster", "err", err)
		if reconnectInterval != 0 {
			level.Info(p.logger).Log("msg", fmt.Sprintf("will retry joining cluster every %v", reconnectInterval.String()))
		}
	} else {
		level.Debug(p.logger).Log("msg", "joined cluster", "peers", n)
	}
	if reconnectInterval != 0 {
		go p.handleReconnect(reconnectInterval)
	}
	if reconnectTimeout != 0 {
		go p.handleReconnectTimeout(5*time.Minute, reconnectTimeout)
	}
	go p.handleRefresh(DefaultRefreshInterval)
	return err
}
func (p *Peer) setInitialFailed(peers []string, myAddr string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(peers) == 0 {
		return
	}
	p.peerLock.RLock()
	defer p.peerLock.RUnlock()
	now := time.Now()
	for _, peerAddr := range peers {
		if peerAddr == myAddr {
			continue
		}
		host, port, err := net.SplitHostPort(peerAddr)
		if err != nil {
			continue
		}
		ip := net.ParseIP(host)
		if ip == nil {
			continue
		}
		portUint, err := strconv.ParseUint(port, 10, 16)
		if err != nil {
			continue
		}
		pr := peer{status: StatusFailed, leaveTime: now, Node: &memberlist.Node{Addr: ip, Port: uint16(portUint)}}
		p.failedPeers = append(p.failedPeers, pr)
		p.peers[peerAddr] = pr
	}
}

type logWriter struct{ l log.Logger }

func (l *logWriter) Write(b []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b), level.Debug(l.l).Log("memberlist", string(b))
}
func (p *Peer) register(reg prometheus.Registerer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterFailedPeers := prometheus.NewGaugeFunc(prometheus.GaugeOpts{Name: "alertmanager_cluster_failed_peers", Help: "Number indicating the current number of failed peers in the cluster."}, func() float64 {
		p.peerLock.RLock()
		defer p.peerLock.RUnlock()
		return float64(len(p.failedPeers))
	})
	p.failedReconnectionsCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_reconnections_failed_total", Help: "A counter of the number of failed cluster peer reconnection attempts."})
	p.reconnectionsCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_reconnections_total", Help: "A counter of the number of cluster peer reconnections."})
	p.failedRefreshCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_refresh_join_failed_total", Help: "A counter of the number of failed cluster peer joined attempts via refresh."})
	p.refreshCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_refresh_join_total", Help: "A counter of the number of cluster peer joined via refresh."})
	p.peerLeaveCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_peers_left_total", Help: "A counter of the number of peers that have left."})
	p.peerUpdateCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_peers_update_total", Help: "A counter of the number of peers that have updated metadata."})
	p.peerJoinCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "alertmanager_cluster_peers_joined_total", Help: "A counter of the number of peers that have joined."})
	reg.MustRegister(clusterFailedPeers, p.failedReconnectionsCounter, p.reconnectionsCounter, p.peerLeaveCounter, p.peerUpdateCounter, p.peerJoinCounter, p.refreshCounter, p.failedRefreshCounter)
}
func (p *Peer) handleReconnectTimeout(d time.Duration, timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tick := time.NewTicker(d)
	defer tick.Stop()
	for {
		select {
		case <-p.stopc:
			return
		case <-tick.C:
			p.removeFailedPeers(timeout)
		}
	}
}
func (p *Peer) removeFailedPeers(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.peerLock.Lock()
	defer p.peerLock.Unlock()
	now := time.Now()
	keep := make([]peer, 0, len(p.failedPeers))
	for _, pr := range p.failedPeers {
		if pr.leaveTime.Add(timeout).After(now) {
			keep = append(keep, pr)
		} else {
			level.Debug(p.logger).Log("msg", "failed peer has timed out", "peer", pr.Node, "addr", pr.Address())
			delete(p.peers, pr.Name)
		}
	}
	p.failedPeers = keep
}
func (p *Peer) handleReconnect(d time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tick := time.NewTicker(d)
	defer tick.Stop()
	for {
		select {
		case <-p.stopc:
			return
		case <-tick.C:
			p.reconnect()
		}
	}
}
func (p *Peer) reconnect() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.peerLock.RLock()
	failedPeers := p.failedPeers
	p.peerLock.RUnlock()
	logger := log.With(p.logger, "msg", "reconnect")
	for _, pr := range failedPeers {
		if _, err := p.mlist.Join([]string{pr.Address()}); err != nil {
			p.failedReconnectionsCounter.Inc()
			level.Debug(logger).Log("result", "failure", "peer", pr.Node, "addr", pr.Address())
		} else {
			p.reconnectionsCounter.Inc()
			level.Debug(logger).Log("result", "success", "peer", pr.Node, "addr", pr.Address())
		}
	}
}
func (p *Peer) handleRefresh(d time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tick := time.NewTicker(d)
	defer tick.Stop()
	for {
		select {
		case <-p.stopc:
			return
		case <-tick.C:
			p.refresh()
		}
	}
}
func (p *Peer) refresh() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger := log.With(p.logger, "msg", "refresh")
	resolvedPeers, err := resolvePeers(context.Background(), p.knownPeers, p.advertiseAddr, net.Resolver{}, false)
	if err != nil {
		level.Debug(logger).Log("peers", p.knownPeers, "err", err)
		return
	}
	members := p.mlist.Members()
	for _, peer := range resolvedPeers {
		var isPeerFound bool
		for _, member := range members {
			if member.Address() == peer {
				isPeerFound = true
				break
			}
		}
		if !isPeerFound {
			if _, err := p.mlist.Join([]string{peer}); err != nil {
				p.failedRefreshCounter.Inc()
				level.Warn(logger).Log("result", "failure", "addr", peer)
			} else {
				p.refreshCounter.Inc()
				level.Debug(logger).Log("result", "success", "addr", peer)
			}
		}
	}
}
func (p *Peer) peerJoin(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.peerLock.Lock()
	defer p.peerLock.Unlock()
	var oldStatus PeerStatus
	pr, ok := p.peers[n.Address()]
	if !ok {
		oldStatus = StatusNone
		pr = peer{status: StatusAlive, Node: n}
	} else {
		oldStatus = pr.status
		pr.Node = n
		pr.status = StatusAlive
		pr.leaveTime = time.Time{}
	}
	p.peers[n.Address()] = pr
	p.peerJoinCounter.Inc()
	if oldStatus == StatusFailed {
		level.Debug(p.logger).Log("msg", "peer rejoined", "peer", pr.Node)
		p.failedPeers = removeOldPeer(p.failedPeers, pr.Address())
	}
}
func (p *Peer) peerLeave(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.peerLock.Lock()
	defer p.peerLock.Unlock()
	pr, ok := p.peers[n.Address()]
	if !ok {
		return
	}
	pr.status = StatusFailed
	pr.leaveTime = time.Now()
	p.failedPeers = append(p.failedPeers, pr)
	p.peers[n.Address()] = pr
	p.peerLeaveCounter.Inc()
	level.Debug(p.logger).Log("msg", "peer left", "peer", pr.Node)
}
func (p *Peer) peerUpdate(n *memberlist.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.peerLock.Lock()
	defer p.peerLock.Unlock()
	pr, ok := p.peers[n.Address()]
	if !ok {
		return
	}
	pr.Node = n
	p.peers[n.Address()] = pr
	p.peerUpdateCounter.Inc()
	level.Debug(p.logger).Log("msg", "peer updated", "peer", pr.Node)
}
func (p *Peer) AddState(key string, s State, reg prometheus.Registerer) *Channel {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.states[key] = s
	send := func(b []byte) {
		p.delegate.bcast.QueueBroadcast(simpleBroadcast(b))
	}
	peers := func() []*memberlist.Node {
		nodes := p.Peers()
		for i, n := range nodes {
			if n.Name == p.Self().Name {
				nodes = append(nodes[:i], nodes[i+1:]...)
				break
			}
		}
		return nodes
	}
	sendOversize := func(n *memberlist.Node, b []byte) error {
		return p.mlist.SendReliable(n, b)
	}
	return NewChannel(key, send, peers, sendOversize, p.logger, p.stopc, reg)
}
func (p *Peer) Leave(timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(p.stopc)
	level.Debug(p.logger).Log("msg", "leaving cluster")
	return p.mlist.Leave(timeout)
}
func (p *Peer) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.mlist.LocalNode().Name
}
func (p *Peer) ClusterSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.mlist.NumMembers()
}
func (p *Peer) Ready() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-p.readyc:
		return true
	default:
	}
	return false
}
func (p *Peer) WaitReady() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-p.readyc
}
func (p *Peer) Status() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.Ready() {
		return "ready"
	} else {
		return "settling"
	}
}
func (p *Peer) Info() map[string]interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return map[string]interface{}{"self": p.mlist.LocalNode(), "members": p.mlist.Members()}
}
func (p *Peer) Self() *memberlist.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.mlist.LocalNode()
}
func (p *Peer) Peers() []*memberlist.Node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.mlist.Members()
}
func (p *Peer) Position() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	all := p.Peers()
	sort.Slice(all, func(i, j int) bool {
		return all[i].Name < all[j].Name
	})
	k := 0
	for _, n := range all {
		if n.Name == p.Self().Name {
			break
		}
		k++
	}
	return k
}
func (p *Peer) Settle(ctx context.Context, interval time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	const NumOkayRequired = 3
	level.Info(p.logger).Log("msg", "Waiting for gossip to settle...", "interval", interval)
	start := time.Now()
	nPeers := 0
	nOkay := 0
	totalPolls := 0
	for {
		select {
		case <-ctx.Done():
			elapsed := time.Since(start)
			level.Info(p.logger).Log("msg", "gossip not settled but continuing anyway", "polls", totalPolls, "elapsed", elapsed)
			close(p.readyc)
			return
		case <-time.After(interval):
		}
		elapsed := time.Since(start)
		n := len(p.Peers())
		if nOkay >= NumOkayRequired {
			level.Info(p.logger).Log("msg", "gossip settled; proceeding", "elapsed", elapsed)
			break
		}
		if n == nPeers {
			nOkay++
			level.Debug(p.logger).Log("msg", "gossip looks settled", "elapsed", elapsed)
		} else {
			nOkay = 0
			level.Info(p.logger).Log("msg", "gossip not settled", "polls", totalPolls, "before", nPeers, "now", n, "elapsed", elapsed)
		}
		nPeers = n
		totalPolls++
	}
	close(p.readyc)
}

type State interface {
	MarshalBinary() ([]byte, error)
	Merge(b []byte) error
}
type simpleBroadcast []byte

func (b simpleBroadcast) Message() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte(b)
}
func (b simpleBroadcast) Invalidates(memberlist.Broadcast) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (b simpleBroadcast) Finished() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func resolvePeers(ctx context.Context, peers []string, myAddress string, res net.Resolver, waitIfEmpty bool) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var resolvedPeers []string
	for _, peer := range peers {
		host, port, err := net.SplitHostPort(peer)
		if err != nil {
			return nil, errors.Wrapf(err, "split host/port for peer %s", peer)
		}
		retryCtx, cancel := context.WithCancel(ctx)
		ips, err := res.LookupIPAddr(ctx, host)
		if err != nil {
			resolvedPeers = append(resolvedPeers, peer)
			continue
		}
		if len(ips) == 0 {
			var lookupErrSpotted bool
			err := retry(2*time.Second, retryCtx.Done(), func() error {
				if lookupErrSpotted {
					cancel()
				}
				ips, err = res.LookupIPAddr(retryCtx, host)
				if err != nil {
					lookupErrSpotted = true
					return errors.Wrapf(err, "IP Addr lookup for peer %s", peer)
				}
				ips = removeMyAddr(ips, port, myAddress)
				if len(ips) == 0 {
					if !waitIfEmpty {
						return nil
					}
					return errors.New("empty IPAddr result. Retrying")
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		for _, ip := range ips {
			resolvedPeers = append(resolvedPeers, net.JoinHostPort(ip.String(), port))
		}
	}
	return resolvedPeers, nil
}
func removeMyAddr(ips []net.IPAddr, targetPort string, myAddr string) []net.IPAddr {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result []net.IPAddr
	for _, ip := range ips {
		if net.JoinHostPort(ip.String(), targetPort) == myAddr {
			continue
		}
		result = append(result, ip)
	}
	return result
}
func hasNonlocal(clusterPeers []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, peer := range clusterPeers {
		if host, _, err := net.SplitHostPort(peer); err == nil {
			peer = host
		}
		if ip := net.ParseIP(peer); ip != nil && !ip.IsLoopback() {
			return true
		} else if ip == nil && strings.ToLower(peer) != "localhost" {
			return true
		}
	}
	return false
}
func isUnroutable(addr string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if host, _, err := net.SplitHostPort(addr); err == nil {
		addr = host
	}
	if ip := net.ParseIP(addr); ip != nil && (ip.IsUnspecified() || ip.IsLoopback()) {
		return true
	} else if ip == nil && strings.ToLower(addr) == "localhost" {
		return true
	}
	return false
}
func isAny(addr string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if host, _, err := net.SplitHostPort(addr); err == nil {
		addr = host
	}
	return addr == "" || net.ParseIP(addr).IsUnspecified()
}
func retry(interval time.Duration, stopc <-chan struct{}, f func() error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tick := time.NewTicker(interval)
	defer tick.Stop()
	var err error
	for {
		if err = f(); err == nil {
			return nil
		}
		select {
		case <-stopc:
			return err
		case <-tick.C:
		}
	}
}
func removeOldPeer(old []peer, addr string) []peer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	new := make([]peer, 0, len(old))
	for _, p := range old {
		if p.Address() != addr {
			new = append(new, p)
		}
	}
	return new
}
