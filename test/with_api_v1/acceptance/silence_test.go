package test

import (
	"fmt"
	"testing"
	"time"
	. "github.com/prometheus/alertmanager/test/with_api_v1"
)

func TestSilencing(t *testing.T) {
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
  repeat_interval: 1ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
	am.Push(At(1), Alert("alertname", "test1").Active(1))
	am.Push(At(1), Alert("alertname", "test2").Active(1))
	co.Want(Between(2, 2.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test2").Active(1))
	am.SetSilence(At(2.3), Silence(2.5, 4.5).Match("alertname", "test1"))
	co.Want(Between(3, 3.5), Alert("alertname", "test2").Active(1))
	co.Want(Between(4, 4.5), Alert("alertname", "test2").Active(1))
	co.Want(Between(5, 5.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test2").Active(1))
	at.Run()
}
func TestSilenceDelete(t *testing.T) {
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
  repeat_interval: 1ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
	am.Push(At(1), Alert("alertname", "test1").Active(1))
	am.Push(At(1), Alert("alertname", "test2").Active(1))
	sil := Silence(1.5, 100).MatchRE("alertname", ".*")
	am.SetSilence(At(1.3), sil)
	am.DelSilence(At(3.5), sil)
	co.Want(Between(3.5, 4.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test2").Active(1))
	at.Run()
}
