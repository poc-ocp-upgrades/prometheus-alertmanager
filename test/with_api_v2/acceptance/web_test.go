package test

import (
	"testing"
	a "github.com/prometheus/alertmanager/test/with_api_v2"
)

func TestWebWithPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: []
  group_wait:      1s
  group_interval:  1s
  repeat_interval: 1h

receivers:
- name: "default"
`
	at := a.NewAcceptanceTest(t, &a.AcceptanceOpts{RoutePrefix: "/foo"})
	at.AlertmanagerCluster(conf, 1)
	at.Run()
}
