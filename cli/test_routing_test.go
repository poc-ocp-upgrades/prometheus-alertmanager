package cli

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"github.com/prometheus/alertmanager/client"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/dispatch"
)

type routingTestDefinition struct {
	alert			client.LabelSet
	expectedReceivers	[]string
	configFile		string
}

func checkResolvedReceivers(mainRoute *dispatch.Route, ls client.LabelSet, expectedReceivers []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resolvedReceivers, err := resolveAlertReceivers(mainRoute, &ls)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(expectedReceivers, resolvedReceivers) {
		return fmt.Errorf("Unexpected routing result want: `%s`, got: `%s`", strings.Join(expectedReceivers, ","), strings.Join(resolvedReceivers, ","))
	}
	return nil
}
func TestRoutingTest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []*routingTestDefinition{&routingTestDefinition{configFile: "testdata/conf.routing.yml", alert: client.LabelSet{"test": "1"}, expectedReceivers: []string{"test1"}}, &routingTestDefinition{configFile: "testdata/conf.routing.yml", alert: client.LabelSet{"test": "2"}, expectedReceivers: []string{"test1", "test2"}}, &routingTestDefinition{configFile: "testdata/conf.routing-reverted.yml", alert: client.LabelSet{"test": "2"}, expectedReceivers: []string{"test2", "test1"}}, &routingTestDefinition{configFile: "testdata/conf.routing.yml", alert: client.LabelSet{"test": "volovina"}, expectedReceivers: []string{"default"}}}
	for _, test := range tests {
		cfg, _, err := config.LoadFile(test.configFile)
		if err != nil {
			t.Fatalf("failed to load test configuration: %v", err)
		}
		mainRoute := dispatch.NewRoute(cfg.Route, nil)
		err = checkResolvedReceivers(mainRoute, test.alert, test.expectedReceivers)
		if err != nil {
			t.Fatalf("%v", err)
		}
		fmt.Println("  OK")
	}
}
