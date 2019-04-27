package cli

import (
	"context"
	"errors"
	"github.com/prometheus/client_golang/api"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/prometheus/alertmanager/client"
)

type silenceExpireCmd struct{ ids []string }

func configureSilenceExpireCmd(cc *kingpin.CmdClause) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		c		= &silenceExpireCmd{}
		expireCmd	= cc.Command("expire", "expire an alertmanager silence")
	)
	expireCmd.Arg("silence-ids", "Ids of silences to expire").StringsVar(&c.ids)
	expireCmd.Action(execWithTimeout(c.expire))
}
func (c *silenceExpireCmd) expire(ctx context.Context, _ *kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.ids) < 1 {
		return errors.New("no silence IDs specified")
	}
	apiClient, err := api.NewClient(api.Config{Address: alertmanagerURL.String()})
	if err != nil {
		return err
	}
	silenceAPI := client.NewSilenceAPI(apiClient)
	for _, id := range c.ids {
		err := silenceAPI.Expire(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}
