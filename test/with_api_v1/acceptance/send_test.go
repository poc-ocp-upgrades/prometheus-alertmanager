package test

import (
	"fmt"
	"testing"
	"time"
	. "github.com/prometheus/alertmanager/test/with_api_v1"
)

func TestMergeAlerts(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: [alertname]
  group_wait:      1s
  group_interval:  1s
  repeat_interval: 1ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
    send_resolved: true
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
	am.Push(At(1), Alert("alertname", "test").Active(1.1))
	am.Push(At(1.2), Alert("alertname", "test").Active(1))
	co.Want(Between(2, 2.5), Alert("alertname", "test").Active(1))
	am.Push(At(2.1), Alert("alertname", "test").Annotate("ann", "v1").Active(2))
	co.Want(Between(3, 3.5), Alert("alertname", "test").Annotate("ann", "v1").Active(1))
	am.Push(At(3.6), Alert("alertname", "test").Annotate("ann", "v2").Active(1.5))
	co.Want(Between(4, 4.5), Alert("alertname", "test").Annotate("ann", "v2").Active(1))
	am.Push(At(4.6), Alert("alertname", "test").Annotate("ann", "v2").Active(3, 4.5))
	am.Push(At(4.8), Alert("alertname", "test").Annotate("ann", "v3").Active(2.9, 4.8))
	am.Push(At(4.8), Alert("alertname", "test").Annotate("ann", "v3").Active(2.9, 4.1))
	co.Want(Between(5, 5.5), Alert("alertname", "test").Annotate("ann", "v3").Active(1, 4.8))
	am.Push(At(5.3), Alert("alertname", "test"))
	co.Want(Between(6, 6.5), Alert("alertname", "test").Active(5.3))
	at.Run()
}
func TestRepeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: [alertname]
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
	am.Push(At(1), Alert("alertname", "test").Active(1))
	am.Push(At(3.5), Alert("alertname", "test").Active(1, 3))
	co.Want(Between(2, 2.5), Alert("alertname", "test").Active(1))
	co.Want(Between(3, 3.5), Alert("alertname", "test").Active(1))
	co.Want(Between(4, 4.5), Alert("alertname", "test").Active(1, 3))
	at.Run()
}
func TestRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: [alertname]
  group_wait:      1s
  group_interval:  1s
  repeat_interval: 3s

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
  - url: 'http://%s'
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co1 := at.Collector("webhook")
	wh1 := NewWebhook(co1)
	co2 := at.Collector("webhook_failing")
	wh2 := NewWebhook(co2)
	wh2.Func = func(ts float64) bool {
		return ts < 4.5
	}
	am := at.Alertmanager(fmt.Sprintf(conf, wh1.Address(), wh2.Address()))
	am.Push(At(1), Alert("alertname", "test1"))
	co1.Want(Between(2, 2.5), Alert("alertname", "test1").Active(1))
	co1.Want(Between(5, 5.5), Alert("alertname", "test1").Active(1))
	co2.Want(Between(4.5, 5), Alert("alertname", "test1").Active(1))
}
func TestBatching(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: []
  group_wait:      1s
  group_interval:  1s
  # use a value slightly below the 5s interval to avoid timing issues
  repeat_interval: 4900ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
	am.Push(At(1.1), Alert("alertname", "test1").Active(1))
	am.Push(At(1.7), Alert("alertname", "test5").Active(1))
	co.Want(Between(2.0, 2.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test5").Active(1))
	am.Push(At(3.3), Alert("alertname", "test2").Active(1.5), Alert("alertname", "test3").Active(1.5), Alert("alertname", "test4").Active(1.6))
	co.Want(Between(4.1, 4.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test5").Active(1), Alert("alertname", "test2").Active(1.5), Alert("alertname", "test3").Active(1.5), Alert("alertname", "test4").Active(1.6))
	co.Want(Between(9.1, 9.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test5").Active(1), Alert("alertname", "test2").Active(1.5), Alert("alertname", "test3").Active(1.5), Alert("alertname", "test4").Active(1.6))
	at.Run()
}
func TestResolved(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	for i := 0; i < 2; i++ {
		conf := `
global:
  resolve_timeout: 10s

route:
  receiver: "default"
  group_by: [alertname]
  group_wait: 1s
  group_interval: 5s

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`
		at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
		co := at.Collector("webhook")
		wh := NewWebhook(co)
		am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
		am.Push(At(1), Alert("alertname", "test", "lbl", "v1"), Alert("alertname", "test", "lbl", "v2"), Alert("alertname", "test", "lbl", "v3"))
		co.Want(Between(2, 2.5), Alert("alertname", "test", "lbl", "v1").Active(1), Alert("alertname", "test", "lbl", "v2").Active(1), Alert("alertname", "test", "lbl", "v3").Active(1))
		co.Want(Between(12, 13), Alert("alertname", "test", "lbl", "v1").Active(1, 11), Alert("alertname", "test", "lbl", "v2").Active(1, 11), Alert("alertname", "test", "lbl", "v3").Active(1, 11))
		at.Run()
	}
}
func TestResolvedFilter(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
global:
  resolve_timeout: 10s

route:
  receiver: "default"
  group_by: [alertname]
  group_wait: 1s
  group_interval: 5s

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
    send_resolved: true
  - url: 'http://%s'
    send_resolved: false
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co1 := at.Collector("webhook1")
	wh1 := NewWebhook(co1)
	co2 := at.Collector("webhook2")
	wh2 := NewWebhook(co2)
	am := at.Alertmanager(fmt.Sprintf(conf, wh1.Address(), wh2.Address()))
	am.Push(At(1), Alert("alertname", "test", "lbl", "v1"), Alert("alertname", "test", "lbl", "v2"))
	am.Push(At(3), Alert("alertname", "test", "lbl", "v1").Active(1, 4), Alert("alertname", "test", "lbl", "v3"))
	am.Push(At(8), Alert("alertname", "test", "lbl", "v3").Active(3))
	co1.Want(Between(2, 2.5), Alert("alertname", "test", "lbl", "v1").Active(1), Alert("alertname", "test", "lbl", "v2").Active(1))
	co1.Want(Between(7, 7.5), Alert("alertname", "test", "lbl", "v1").Active(1, 4), Alert("alertname", "test", "lbl", "v2").Active(1), Alert("alertname", "test", "lbl", "v3").Active(3))
	co1.Want(Between(12, 12.5), Alert("alertname", "test", "lbl", "v2").Active(1, 11), Alert("alertname", "test", "lbl", "v3").Active(3))
	co2.Want(Between(2, 2.5), Alert("alertname", "test", "lbl", "v1").Active(1), Alert("alertname", "test", "lbl", "v2").Active(1))
	co2.Want(Between(7, 7.5), Alert("alertname", "test", "lbl", "v2").Active(1), Alert("alertname", "test", "lbl", "v3").Active(3))
	co2.Want(Between(12, 12.5))
	at.Run()
}
func TestReload(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	conf := `
route:
  receiver: "default"
  group_by: []
  group_wait:      1s
  group_interval:  6s
  repeat_interval: 10m

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	am := at.Alertmanager(fmt.Sprintf(conf, wh.Address()))
	am.Push(At(1), Alert("alertname", "test1"))
	at.Do(At(3), am.Reload)
	am.Push(At(4), Alert("alertname", "test2"))
	co.Want(Between(2, 2.5), Alert("alertname", "test1").Active(1))
	co.Want(Between(9, 9.5), Alert("alertname", "test1").Active(1), Alert("alertname", "test2").Active(4))
	at.Run()
}
