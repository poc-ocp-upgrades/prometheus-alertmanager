package cli

import (
	"bytes"
	"context"
	"fmt"
	"github.com/xlab/treeprint"
	"github.com/prometheus/alertmanager/client"
	"github.com/prometheus/alertmanager/dispatch"
	"gopkg.in/alecthomas/kingpin.v2"
)

type routingShow struct {
	configFile		string
	labels			[]string
	expectedReceivers	string
	debugTree		bool
}

const (
	routingHelp	= `Prints alert routing tree

Will print whole routing tree in form of ASCII tree view.

Routing is loaded from a local configuration file or a running Alertmanager configuration.
Specifying --config.file takes precedence over --alertmanager.url.

Example:

./amtool config routes [show] --config.file=doc/examples/simple.yml

`
	branchSlugSeparator	= "  "
)

func configureRoutingCmd(app *kingpin.CmdClause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		c		= &routingShow{}
		routingCmd	= app.Command("routes", routingHelp)
		routingShowCmd	= routingCmd.Command("show", routingHelp).Default()
		configFlag	= routingCmd.Flag("config.file", "Config file to be tested.")
	)
	configFlag.ExistingFileVar(&c.configFile)
	routingShowCmd.Action(execWithTimeout(c.routingShowAction))
	configureRoutingTestCmd(routingCmd, c)
}
func (c *routingShow) routingShowAction(ctx context.Context, _ *kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, err := loadAlertmanagerConfig(ctx, alertmanagerURL, c.configFile)
	if err != nil {
		kingpin.Fatalf("%s", err)
		return err
	}
	route := dispatch.NewRoute(cfg.Route, nil)
	tree := treeprint.New()
	convertRouteToTree(route, tree)
	fmt.Println("Routing tree:")
	fmt.Println(tree.String())
	return nil
}
func getRouteTreeSlug(route *dispatch.Route, showContinue bool, showReceiver bool) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var branchSlug bytes.Buffer
	if route.Matchers.Len() == 0 {
		branchSlug.WriteString("default-route")
	} else {
		branchSlug.WriteString(route.Matchers.String())
	}
	if route.Continue && showContinue {
		branchSlug.WriteString(branchSlugSeparator)
		branchSlug.WriteString("continue: true")
	}
	if showReceiver {
		branchSlug.WriteString(branchSlugSeparator)
		branchSlug.WriteString("receiver: ")
		branchSlug.WriteString(route.RouteOpts.Receiver)
	}
	return branchSlug.String()
}
func convertRouteToTree(route *dispatch.Route, tree treeprint.Tree) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	branch := tree.AddBranch(getRouteTreeSlug(route, true, true))
	for _, r := range route.Routes {
		convertRouteToTree(r, branch)
	}
}
func getMatchingTree(route *dispatch.Route, tree treeprint.Tree, lset client.LabelSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	final := true
	branch := tree.AddBranch(getRouteTreeSlug(route, false, false))
	for _, r := range route.Routes {
		if r.Matchers.Match(convertClientToCommonLabelSet(lset)) {
			getMatchingTree(r, branch, lset)
			final = false
			if !r.Continue {
				break
			}
		}
	}
	if final {
		branch.SetValue(getRouteTreeSlug(route, false, true))
	}
}
