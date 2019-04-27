package cluster

import (
	"errors"
	"net"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCalculateAdvertiseAddress(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	old := getPrivateAddress
	defer func() {
		getPrivateAddress = old
	}()
	cases := []struct {
		fn		getPrivateIPFunc
		bind, advertise	string
		expectedIP	net.IP
		err		bool
	}{{bind: "192.0.2.1", advertise: "", expectedIP: net.ParseIP("192.0.2.1"), err: false}, {bind: "192.0.2.1", advertise: "192.0.2.2", expectedIP: net.ParseIP("192.0.2.2"), err: false}, {fn: func() (string, error) {
		return "192.0.2.1", nil
	}, bind: "0.0.0.0", advertise: "", expectedIP: net.ParseIP("192.0.2.1"), err: false}, {fn: func() (string, error) {
		return "", errors.New("some error")
	}, bind: "0.0.0.0", advertise: "", err: true}, {fn: func() (string, error) {
		return "invalid", nil
	}, bind: "0.0.0.0", advertise: "", err: true}, {fn: func() (string, error) {
		return "", nil
	}, bind: "0.0.0.0", advertise: "", err: true}}
	for _, c := range cases {
		getPrivateAddress = c.fn
		got, err := calculateAdvertiseAddress(c.bind, c.advertise)
		if c.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, c.expectedIP.String(), got.String())
		}
	}
}
