package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"github.com/prometheus/alertmanager/client"
	"github.com/prometheus/alertmanager/dispatch"
	"github.com/xlab/treeprint"
	"gopkg.in/alecthomas/kingpin.v2"
)

const routingTestHelp = `Test alert routing

Will return receiver names which the alert with given labels resolves to.
If the labelset resolves to multiple receivers, they are printed out in order as defined in the routing tree.

Routing is loaded from a local configuration file or a running Alertmanager configuration.
Specifying --config.file takes precedence over --alertmanager.url.

Example:

./amtool config routes test --config.file=doc/examples/simple.yml --verify.receivers=team-DB-pager service=database

`

func configureRoutingTestCmd(cc *kingpin.CmdClause, c *routingShow) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var routingTestCmd = cc.Command("test", routingTestHelp)
	routingTestCmd.Flag("verify.receivers", "Checks if specified receivers matches resolved receivers. The command fails if the labelset does not route to the specified receivers.").StringVar(&c.expectedReceivers)
	routingTestCmd.Flag("tree", "Prints out matching routes tree.").BoolVar(&c.debugTree)
	routingTestCmd.Arg("labels", "List of labels to be tested against the configured routes.").StringsVar(&c.labels)
	routingTestCmd.Action(execWithTimeout(c.routingTestAction))
}
func resolveAlertReceivers(mainRoute *dispatch.Route, labels *client.LabelSet) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		finalRoutes	[]*dispatch.Route
		receivers	[]string
	)
	finalRoutes = mainRoute.Match(convertClientToCommonLabelSet(*labels))
	for _, r := range finalRoutes {
		receivers = append(receivers, r.RouteOpts.Receiver)
	}
	return receivers, nil
}
func printMatchingTree(mainRoute *dispatch.Route, ls client.LabelSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tree := treeprint.New()
	getMatchingTree(mainRoute, tree, ls)
	fmt.Println("Matching routes:")
	fmt.Println(tree.String())
	fmt.Print("\n")
}
func (c *routingShow) routingTestAction(ctx context.Context, _ *kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, err := loadAlertmanagerConfig(ctx, alertmanagerURL, c.configFile)
	if err != nil {
		kingpin.Fatalf("%v\n", err)
		return err
	}
	mainRoute := dispatch.NewRoute(cfg.Route, nil)
	ls, err := parseLabels(c.labels)
	if err != nil {
		kingpin.Fatalf("Failed to parse labels: %v\n", err)
	}
	if c.debugTree {
		printMatchingTree(mainRoute, ls)
	}
	receivers, err := resolveAlertReceivers(mainRoute, &ls)
	receiversSlug := strings.Join(receivers, ",")
	fmt.Printf("%s\n", receiversSlug)
	if c.expectedReceivers != "" && c.expectedReceivers != receiversSlug {
		fmt.Printf("WARNING: Expected receivers did not match resolved receivers.\n")
		os.Exit(1)
	}
	return err
}