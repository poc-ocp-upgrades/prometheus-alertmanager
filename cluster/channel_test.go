package cluster

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"github.com/go-kit/kit/log"
	"github.com/hashicorp/memberlist"
	"github.com/prometheus/client_golang/prometheus"
)

func TestNormalMessagesGossiped(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sent bool
	c := newChannel(func(_ []byte) {
		sent = true
	}, func() []*memberlist.Node {
		return nil
	}, func(_ *memberlist.Node, _ []byte) error {
		return nil
	})
	c.Broadcast([]byte{})
	if sent != true {
		t.Fatalf("small message not sent")
	}
}
func TestOversizedMessagesGossiped(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sent bool
	ctx, cancel := context.WithCancel(context.Background())
	c := newChannel(func(_ []byte) {
	}, func() []*memberlist.Node {
		return []*memberlist.Node{&memberlist.Node{}}
	}, func(_ *memberlist.Node, _ []byte) error {
		sent = true
		cancel()
		return nil
	})
	f, err := os.Open("/dev/zero")
	if err != nil {
		t.Fatalf("failed to open /dev/zero: %v", err)
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	toCopy := int64(800)
	if n, err := io.CopyN(buf, f, toCopy); err != nil {
		t.Fatalf("failed to copy bytes: %v", err)
	} else if n != toCopy {
		t.Fatalf("wanted to copy %d bytes, only copied %d", toCopy, n)
	}
	c.Broadcast(buf.Bytes())
	<-ctx.Done()
	if sent != true {
		t.Fatalf("oversized message not sent")
	}
}
func newChannel(send func([]byte), peers func() []*memberlist.Node, sendOversize func(*memberlist.Node, []byte) error) *Channel {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewChannel("test", send, peers, sendOversize, log.NewNopLogger(), make(chan struct{}), prometheus.NewRegistry())
}
