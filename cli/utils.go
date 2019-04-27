package cli

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"github.com/prometheus/alertmanager/client"
	amconfig "github.com/prometheus/alertmanager/config"
	"github.com/prometheus/client_golang/api"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/prometheus/alertmanager/pkg/parse"
	"github.com/prometheus/alertmanager/types"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

type ByAlphabetical []labels.Matcher

func (s ByAlphabetical) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s)
}
func (s ByAlphabetical) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s[i], s[j] = s[j], s[i]
}
func (s ByAlphabetical) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s[i].Name != s[j].Name {
		return s[i].Name < s[j].Name
	} else if s[i].Type != s[j].Type {
		return s[i].Type < s[j].Type
	} else if s[i].Value != s[j].Value {
		return s[i].Value < s[j].Value
	}
	return false
}
func GetAlertmanagerURL(p string) url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	amURL := *alertmanagerURL
	amURL.Path = path.Join(alertmanagerURL.Path, p)
	return amURL
}
func parseMatchers(inputMatchers []string) ([]labels.Matcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	matchers := make([]labels.Matcher, 0)
	for _, v := range inputMatchers {
		name, value, matchType, err := parse.Input(v)
		if err != nil {
			return []labels.Matcher{}, err
		}
		matchers = append(matchers, labels.Matcher{Type: matchType, Name: name, Value: value})
	}
	return matchers, nil
}
func getRemoteAlertmanagerConfigStatus(ctx context.Context, alertmanagerURL *url.URL) (*client.ServerStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := api.NewClient(api.Config{Address: alertmanagerURL.String()})
	if err != nil {
		return nil, err
	}
	statusAPI := client.NewStatusAPI(c)
	status, err := statusAPI.Get(ctx)
	if err != nil {
		return nil, err
	}
	return status, nil
}
func checkRoutingConfigInputFlags(alertmanagerURL *url.URL, configFile string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if alertmanagerURL != nil && configFile != "" {
		fmt.Fprintln(os.Stderr, "Warning: --config.file flag overrides the --alertmanager.url.")
	}
	if alertmanagerURL == nil && configFile == "" {
		kingpin.Fatalf("You have to specify one of --config.file or --alertmanager.url flags.")
	}
}
func loadAlertmanagerConfig(ctx context.Context, alertmanagerURL *url.URL, configFile string) (*amconfig.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkRoutingConfigInputFlags(alertmanagerURL, configFile)
	if configFile != "" {
		cfg, _, err := amconfig.LoadFile(configFile)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}
	if alertmanagerURL != nil {
		status, err := getRemoteAlertmanagerConfigStatus(ctx, alertmanagerURL)
		if err != nil {
			return nil, err
		}
		return status.ConfigJSON, nil
	}
	return nil, errors.New("failed to get Alertmanager configuration")
}
func convertClientToCommonLabelSet(cls client.LabelSet) model.LabelSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mls := make(model.LabelSet, len(cls))
	for ln, lv := range cls {
		mls[model.LabelName(ln)] = model.LabelValue(lv)
	}
	return mls
}
func parseLabels(inputLabels []string) (client.LabelSet, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	labelSet := make(client.LabelSet, len(inputLabels))
	for _, l := range inputLabels {
		name, value, matchType, err := parse.Input(l)
		if err != nil {
			return client.LabelSet{}, err
		}
		if matchType != labels.MatchEqual {
			return client.LabelSet{}, errors.New("labels must be specified as key=value pairs")
		}
		labelSet[client.LabelName(name)] = client.LabelValue(value)
	}
	return labelSet, nil
}
func TypeMatchers(matchers []labels.Matcher) (types.Matchers, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	typeMatchers := types.Matchers{}
	for _, matcher := range matchers {
		typeMatcher, err := TypeMatcher(matcher)
		if err != nil {
			return types.Matchers{}, err
		}
		typeMatchers = append(typeMatchers, &typeMatcher)
	}
	return typeMatchers, nil
}
func TypeMatcher(matcher labels.Matcher) (types.Matcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	typeMatcher := types.NewMatcher(model.LabelName(matcher.Name), matcher.Value)
	switch matcher.Type {
	case labels.MatchEqual:
		typeMatcher.IsRegex = false
	case labels.MatchRegexp:
		typeMatcher.IsRegex = true
	default:
		return types.Matcher{}, fmt.Errorf("invalid match type for creation operation: %s", matcher.Type)
	}
	return *typeMatcher, nil
}
func execWithTimeout(fn func(context.Context, *kingpin.ParseContext) error) func(*kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(x *kingpin.ParseContext) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return fn(ctx, x)
	}
}
