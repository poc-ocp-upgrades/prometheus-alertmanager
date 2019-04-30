package cli

import (
	"context"
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/prometheus/alertmanager/cli/format"
)

const configHelp = `View current config.

The amount of output is controlled by the output selection flag:
	- Simple: Print just the running config
	- Extended: Print the running config as well as uptime and all version info
	- Json: Print entire config object as json
`

func configureConfigCmd(app *kingpin.Application) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configCmd := app.Command("config", configHelp)
	configCmd.Command("show", configHelp).Default().Action(execWithTimeout(queryConfig)).PreAction(requireAlertManagerURL)
	configureRoutingCmd(configCmd)
}
func queryConfig(ctx context.Context, _ *kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status, err := getRemoteAlertmanagerConfigStatus(ctx, alertmanagerURL)
	if err != nil {
		return err
	}
	formatter, found := format.Formatters[output]
	if !found {
		return errors.New("unknown output formatter")
	}
	return formatter.FormatConfig(status)
}
