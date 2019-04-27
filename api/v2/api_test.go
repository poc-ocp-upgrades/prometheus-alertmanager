package v2

import (
	"testing"
	"time"
	general_ops "github.com/prometheus/alertmanager/api/v2/restapi/operations/general"
	"github.com/prometheus/alertmanager/config"
)

func TestGetStatusHandlerWithNilPeer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := API{uptime: time.Now(), peer: nil, alertmanagerConfig: &config.Config{}}
	status := api.getStatusHandler(general_ops.GetStatusParams{}).(*general_ops.GetStatusOK)
	c := status.Payload.Cluster
	if c == nil || c.Status == nil || c.Name == nil || c.Peers == nil {
		t.Fatal("expected cluster {status,name,peers} not to be nil, violating the openapi specification")
	}
}
