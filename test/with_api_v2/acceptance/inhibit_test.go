package test

import (
	"fmt"
	"testing"
	"time"
	. "github.com/prometheus/alertmanager/test/with_api_v2"
)

func TestInhibiting(t *testing.T) {
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
  repeat_interval: 1s

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'

inhibit_rules:
- source_match:
    alertname: JobDown
  target_match:
    alertname: InstanceDown
  equal:
    - job
    - zone
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	amc := at.AlertmanagerCluster(fmt.Sprintf(conf, wh.Address()), 1)
	amc.Push(At(1), Alert("alertname", "test1", "job", "testjob", "zone", "aa"))
	amc.Push(At(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa"))
	amc.Push(At(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "ab"))
	amc.Push(At(2.2), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa"))
	amc.Push(At(3.6), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(2.2, 3.6))
	co.Want(Between(2, 2.5), Alert("alertname", "test1", "job", "testjob", "zone", "aa").Active(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa").Active(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "ab").Active(1))
	co.Want(Between(3, 3.5), Alert("alertname", "test1", "job", "testjob", "zone", "aa").Active(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "ab").Active(1), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(2.2))
	co.Want(Between(4, 4.5), Alert("alertname", "test1", "job", "testjob", "zone", "aa").Active(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa").Active(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "ab").Active(1), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(2.2, 3.6))
	at.Run()
	t.Log(co.Check())
}
func TestAlwaysInhibiting(t *testing.T) {
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
  repeat_interval: 1s

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'

inhibit_rules:
- source_match:
    alertname: JobDown
  target_match:
    alertname: InstanceDown
  equal:
    - job
    - zone
`
	at := NewAcceptanceTest(t, &AcceptanceOpts{Tolerance: 150 * time.Millisecond})
	co := at.Collector("webhook")
	wh := NewWebhook(co)
	amc := at.AlertmanagerCluster(fmt.Sprintf(conf, wh.Address()), 1)
	amc.Push(At(1), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa"))
	amc.Push(At(1), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa"))
	amc.Push(At(2.6), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(1, 2.6))
	amc.Push(At(2.6), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa").Active(1, 2.6))
	co.Want(Between(2, 2.5), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(1))
	co.Want(Between(3, 3.5), Alert("alertname", "InstanceDown", "job", "testjob", "zone", "aa").Active(1, 2.6), Alert("alertname", "JobDown", "job", "testjob", "zone", "aa").Active(1, 2.6))
	at.Run()
	t.Log(co.Check())
}
