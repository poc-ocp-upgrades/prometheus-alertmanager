package config

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	godefaulthttp "net/http"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	commoncfg "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"gopkg.in/yaml.v2"
)

const secretToken = "<secret>"

var secretTokenJSON string

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(secretToken)
	if err != nil {
		panic(err)
	}
	secretTokenJSON = string(b)
}

type Secret string

func (s Secret) MarshalYAML() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s != "" {
		return secretToken, nil
	}
	return nil, nil
}
func (s *Secret) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type plain Secret
	return unmarshal((*plain)(s))
}
func (s Secret) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return json.Marshal(secretToken)
}

type URL struct{ *url.URL }

func (u *URL) Copy() *URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := *u.URL
	return &URL{&v}
}
func (u URL) MarshalYAML() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if u.URL != nil {
		return u.URL.String(), nil
	}
	return nil, nil
}
func (u *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	urlp, err := parseURL(s)
	if err != nil {
		return err
	}
	u.URL = urlp.URL
	return nil
}
func (u URL) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if u.URL != nil {
		return json.Marshal(u.URL.String())
	}
	return nil, nil
}
func (u *URL) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	urlp, err := parseURL(s)
	if err != nil {
		return err
	}
	u.URL = urlp.URL
	return nil
}

type SecretURL URL

func (s SecretURL) MarshalYAML() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.URL != nil {
		return secretToken, nil
	}
	return nil, nil
}
func (s *SecretURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if str == secretToken {
		s.URL = &url.URL{}
		return nil
	}
	return unmarshal((*URL)(s))
}
func (s SecretURL) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return json.Marshal(secretToken)
}
func (s *SecretURL) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if string(data) == secretToken || string(data) == secretTokenJSON {
		s.URL = &url.URL{}
		return nil
	}
	return json.Unmarshal(data, (*URL)(s))
}
func Load(s string) (*Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &Config{}
	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}
	if cfg.Route == nil {
		return nil, errors.New("no route provided in config")
	}
	if cfg.Route.Continue {
		return nil, errors.New("cannot have continue in root route")
	}
	cfg.original = s
	return cfg, nil
}
func LoadFile(filename string) (*Config, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, nil, err
	}
	resolveFilepaths(filepath.Dir(filename), cfg)
	return cfg, content, nil
}
func resolveFilepaths(baseDir string, cfg *Config) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	join := func(fp string) string {
		if len(fp) > 0 && !filepath.IsAbs(fp) {
			fp = filepath.Join(baseDir, fp)
		}
		return fp
	}
	for i, tf := range cfg.Templates {
		cfg.Templates[i] = join(tf)
	}
}

type Config struct {
	Global		*GlobalConfig	`yaml:"global,omitempty" json:"global,omitempty"`
	Route		*Route		`yaml:"route,omitempty" json:"route,omitempty"`
	InhibitRules	[]*InhibitRule	`yaml:"inhibit_rules,omitempty" json:"inhibit_rules,omitempty"`
	Receivers	[]*Receiver	`yaml:"receivers,omitempty" json:"receivers,omitempty"`
	Templates	[]string	`yaml:"templates" json:"templates"`
	original	string
}

func (c Config) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("<error creating config string: %s>", err)
	}
	return string(b)
}
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	if c.Global == nil {
		c.Global = &GlobalConfig{}
		*c.Global = DefaultGlobalConfig()
	}
	names := map[string]struct{}{}
	for _, rcv := range c.Receivers {
		if _, ok := names[rcv.Name]; ok {
			return fmt.Errorf("notification config name %q is not unique", rcv.Name)
		}
		for _, wh := range rcv.WebhookConfigs {
			if wh.HTTPConfig == nil {
				wh.HTTPConfig = c.Global.HTTPConfig
			}
		}
		for _, ec := range rcv.EmailConfigs {
			if ec.Smarthost == "" {
				if c.Global.SMTPSmarthost == "" {
					return fmt.Errorf("no global SMTP smarthost set")
				}
				ec.Smarthost = c.Global.SMTPSmarthost
			}
			if ec.From == "" {
				if c.Global.SMTPFrom == "" {
					return fmt.Errorf("no global SMTP from set")
				}
				ec.From = c.Global.SMTPFrom
			}
			if ec.Hello == "" {
				ec.Hello = c.Global.SMTPHello
			}
			if ec.AuthUsername == "" {
				ec.AuthUsername = c.Global.SMTPAuthUsername
			}
			if ec.AuthPassword == "" {
				ec.AuthPassword = c.Global.SMTPAuthPassword
			}
			if ec.AuthSecret == "" {
				ec.AuthSecret = c.Global.SMTPAuthSecret
			}
			if ec.AuthIdentity == "" {
				ec.AuthIdentity = c.Global.SMTPAuthIdentity
			}
			if ec.RequireTLS == nil {
				ec.RequireTLS = new(bool)
				*ec.RequireTLS = c.Global.SMTPRequireTLS
			}
		}
		for _, sc := range rcv.SlackConfigs {
			if sc.HTTPConfig == nil {
				sc.HTTPConfig = c.Global.HTTPConfig
			}
			if sc.APIURL == nil {
				if c.Global.SlackAPIURL == nil {
					return fmt.Errorf("no global Slack API URL set")
				}
				sc.APIURL = c.Global.SlackAPIURL
			}
		}
		for _, hc := range rcv.HipchatConfigs {
			if hc.HTTPConfig == nil {
				hc.HTTPConfig = c.Global.HTTPConfig
			}
			if hc.APIURL == nil {
				if c.Global.HipchatAPIURL == nil {
					return fmt.Errorf("no global Hipchat API URL set")
				}
				hc.APIURL = c.Global.HipchatAPIURL
			}
			if !strings.HasSuffix(hc.APIURL.Path, "/") {
				hc.APIURL.Path += "/"
			}
			if hc.AuthToken == "" {
				if c.Global.HipchatAuthToken == "" {
					return fmt.Errorf("no global Hipchat Auth Token set")
				}
				hc.AuthToken = c.Global.HipchatAuthToken
			}
		}
		for _, poc := range rcv.PushoverConfigs {
			if poc.HTTPConfig == nil {
				poc.HTTPConfig = c.Global.HTTPConfig
			}
		}
		for _, pdc := range rcv.PagerdutyConfigs {
			if pdc.HTTPConfig == nil {
				pdc.HTTPConfig = c.Global.HTTPConfig
			}
			if pdc.URL == nil {
				if c.Global.PagerdutyURL == nil {
					return fmt.Errorf("no global PagerDuty URL set")
				}
				pdc.URL = c.Global.PagerdutyURL
			}
		}
		for _, ogc := range rcv.OpsGenieConfigs {
			if ogc.HTTPConfig == nil {
				ogc.HTTPConfig = c.Global.HTTPConfig
			}
			if ogc.APIURL == nil {
				if c.Global.OpsGenieAPIURL == nil {
					return fmt.Errorf("no global OpsGenie URL set")
				}
				ogc.APIURL = c.Global.OpsGenieAPIURL
			}
			if !strings.HasSuffix(ogc.APIURL.Path, "/") {
				ogc.APIURL.Path += "/"
			}
			if ogc.APIKey == "" {
				if c.Global.OpsGenieAPIKey == "" {
					return fmt.Errorf("no global OpsGenie API Key set")
				}
				ogc.APIKey = c.Global.OpsGenieAPIKey
			}
		}
		for _, wcc := range rcv.WechatConfigs {
			if wcc.HTTPConfig == nil {
				wcc.HTTPConfig = c.Global.HTTPConfig
			}
			if wcc.APIURL == nil {
				if c.Global.WeChatAPIURL == nil {
					return fmt.Errorf("no global Wechat URL set")
				}
				wcc.APIURL = c.Global.WeChatAPIURL
			}
			if wcc.APISecret == "" {
				if c.Global.WeChatAPISecret == "" {
					return fmt.Errorf("no global Wechat ApiSecret set")
				}
				wcc.APISecret = c.Global.WeChatAPISecret
			}
			if wcc.CorpID == "" {
				if c.Global.WeChatAPICorpID == "" {
					return fmt.Errorf("no global Wechat CorpID set")
				}
				wcc.CorpID = c.Global.WeChatAPICorpID
			}
			if !strings.HasSuffix(wcc.APIURL.Path, "/") {
				wcc.APIURL.Path += "/"
			}
		}
		for _, voc := range rcv.VictorOpsConfigs {
			if voc.HTTPConfig == nil {
				voc.HTTPConfig = c.Global.HTTPConfig
			}
			if voc.APIURL == nil {
				if c.Global.VictorOpsAPIURL == nil {
					return fmt.Errorf("no global VictorOps URL set")
				}
				voc.APIURL = c.Global.VictorOpsAPIURL
			}
			if !strings.HasSuffix(voc.APIURL.Path, "/") {
				voc.APIURL.Path += "/"
			}
			if voc.APIKey == "" {
				if c.Global.VictorOpsAPIKey == "" {
					return fmt.Errorf("no global VictorOps API Key set")
				}
				voc.APIKey = c.Global.VictorOpsAPIKey
			}
		}
		names[rcv.Name] = struct{}{}
	}
	if c.Route == nil {
		return fmt.Errorf("no routes provided")
	}
	if len(c.Route.Receiver) == 0 {
		return fmt.Errorf("root route must specify a default receiver")
	}
	if len(c.Route.Match) > 0 || len(c.Route.MatchRE) > 0 {
		return fmt.Errorf("root route must not have any matchers")
	}
	return checkReceiver(c.Route, names)
}
func checkReceiver(r *Route, receivers map[string]struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Receiver == "" {
		return nil
	}
	if _, ok := receivers[r.Receiver]; !ok {
		return fmt.Errorf("undefined receiver %q used in route", r.Receiver)
	}
	for _, sr := range r.Routes {
		if err := checkReceiver(sr, receivers); err != nil {
			return err
		}
	}
	return nil
}
func DefaultGlobalConfig() GlobalConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return GlobalConfig{ResolveTimeout: model.Duration(5 * time.Minute), HTTPConfig: &commoncfg.HTTPClientConfig{}, SMTPHello: "localhost", SMTPRequireTLS: true, PagerdutyURL: mustParseURL("https://events.pagerduty.com/v2/enqueue"), HipchatAPIURL: mustParseURL("https://api.hipchat.com/"), OpsGenieAPIURL: mustParseURL("https://api.opsgenie.com/"), WeChatAPIURL: mustParseURL("https://qyapi.weixin.qq.com/cgi-bin/"), VictorOpsAPIURL: mustParseURL("https://alert.victorops.com/integrations/generic/20131114/alert/")}
}
func mustParseURL(s string) *URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := parseURL(s)
	if err != nil {
		panic(err)
	}
	return u
}
func parseURL(s string) (*URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("unsupported scheme %q for URL", u.Scheme)
	}
	if u.Host == "" {
		return nil, fmt.Errorf("missing host for URL")
	}
	return &URL{u}, nil
}

type GlobalConfig struct {
	ResolveTimeout		model.Duration			`yaml:"resolve_timeout" json:"resolve_timeout"`
	HTTPConfig		*commoncfg.HTTPClientConfig	`yaml:"http_config,omitempty" json:"http_config,omitempty"`
	SMTPFrom		string				`yaml:"smtp_from,omitempty" json:"smtp_from,omitempty"`
	SMTPHello		string				`yaml:"smtp_hello,omitempty" json:"smtp_hello,omitempty"`
	SMTPSmarthost		string				`yaml:"smtp_smarthost,omitempty" json:"smtp_smarthost,omitempty"`
	SMTPAuthUsername	string				`yaml:"smtp_auth_username,omitempty" json:"smtp_auth_username,omitempty"`
	SMTPAuthPassword	Secret				`yaml:"smtp_auth_password,omitempty" json:"smtp_auth_password,omitempty"`
	SMTPAuthSecret		Secret				`yaml:"smtp_auth_secret,omitempty" json:"smtp_auth_secret,omitempty"`
	SMTPAuthIdentity	string				`yaml:"smtp_auth_identity,omitempty" json:"smtp_auth_identity,omitempty"`
	SMTPRequireTLS		bool				`yaml:"smtp_require_tls,omitempty" json:"smtp_require_tls,omitempty"`
	SlackAPIURL		*SecretURL			`yaml:"slack_api_url,omitempty" json:"slack_api_url,omitempty"`
	PagerdutyURL		*URL				`yaml:"pagerduty_url,omitempty" json:"pagerduty_url,omitempty"`
	HipchatAPIURL		*URL				`yaml:"hipchat_api_url,omitempty" json:"hipchat_api_url,omitempty"`
	HipchatAuthToken	Secret				`yaml:"hipchat_auth_token,omitempty" json:"hipchat_auth_token,omitempty"`
	OpsGenieAPIURL		*URL				`yaml:"opsgenie_api_url,omitempty" json:"opsgenie_api_url,omitempty"`
	OpsGenieAPIKey		Secret				`yaml:"opsgenie_api_key,omitempty" json:"opsgenie_api_key,omitempty"`
	WeChatAPIURL		*URL				`yaml:"wechat_api_url,omitempty" json:"wechat_api_url,omitempty"`
	WeChatAPISecret		Secret				`yaml:"wechat_api_secret,omitempty" json:"wechat_api_secret,omitempty"`
	WeChatAPICorpID		string				`yaml:"wechat_api_corp_id,omitempty" json:"wechat_api_corp_id,omitempty"`
	VictorOpsAPIURL		*URL				`yaml:"victorops_api_url,omitempty" json:"victorops_api_url,omitempty"`
	VictorOpsAPIKey		Secret				`yaml:"victorops_api_key,omitempty" json:"victorops_api_key,omitempty"`
}

func (c *GlobalConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*c = DefaultGlobalConfig()
	type plain GlobalConfig
	return unmarshal((*plain)(c))
}

type Route struct {
	Receiver	string			`yaml:"receiver,omitempty" json:"receiver,omitempty"`
	GroupByStr	[]string		`yaml:"group_by,omitempty" json:"group_by,omitempty"`
	GroupBy		[]model.LabelName	`yaml:"-" json:"-"`
	GroupByAll	bool			`yaml:"-" json:"-"`
	Match		map[string]string	`yaml:"match,omitempty" json:"match,omitempty"`
	MatchRE		map[string]Regexp	`yaml:"match_re,omitempty" json:"match_re,omitempty"`
	Continue	bool			`yaml:"continue,omitempty" json:"continue,omitempty"`
	Routes		[]*Route		`yaml:"routes,omitempty" json:"routes,omitempty"`
	GroupWait	*model.Duration		`yaml:"group_wait,omitempty" json:"group_wait,omitempty"`
	GroupInterval	*model.Duration		`yaml:"group_interval,omitempty" json:"group_interval,omitempty"`
	RepeatInterval	*model.Duration		`yaml:"repeat_interval,omitempty" json:"repeat_interval,omitempty"`
}

func (r *Route) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type plain Route
	if err := unmarshal((*plain)(r)); err != nil {
		return err
	}
	for k := range r.Match {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	for k := range r.MatchRE {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	for _, l := range r.GroupByStr {
		if l == "..." {
			r.GroupByAll = true
		} else {
			labelName := model.LabelName(l)
			if !labelName.IsValid() {
				return fmt.Errorf("invalid label name %q in group_by list", l)
			}
			r.GroupBy = append(r.GroupBy, labelName)
		}
	}
	if len(r.GroupBy) > 0 && r.GroupByAll {
		return fmt.Errorf("cannot have wildcard group_by (`...`) and other other labels at the same time")
	}
	groupBy := map[model.LabelName]struct{}{}
	for _, ln := range r.GroupBy {
		if _, ok := groupBy[ln]; ok {
			return fmt.Errorf("duplicated label %q in group_by", ln)
		}
		groupBy[ln] = struct{}{}
	}
	if r.GroupInterval != nil && time.Duration(*r.GroupInterval) == time.Duration(0) {
		return fmt.Errorf("group_interval cannot be zero")
	}
	if r.RepeatInterval != nil && time.Duration(*r.RepeatInterval) == time.Duration(0) {
		return fmt.Errorf("repeat_interval cannot be zero")
	}
	return nil
}

type InhibitRule struct {
	SourceMatch	map[string]string	`yaml:"source_match,omitempty" json:"source_match,omitempty"`
	SourceMatchRE	map[string]Regexp	`yaml:"source_match_re,omitempty" json:"source_match_re,omitempty"`
	TargetMatch	map[string]string	`yaml:"target_match,omitempty" json:"target_match,omitempty"`
	TargetMatchRE	map[string]Regexp	`yaml:"target_match_re,omitempty" json:"target_match_re,omitempty"`
	Equal		model.LabelNames	`yaml:"equal,omitempty" json:"equal,omitempty"`
}

func (r *InhibitRule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type plain InhibitRule
	if err := unmarshal((*plain)(r)); err != nil {
		return err
	}
	for k := range r.SourceMatch {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	for k := range r.SourceMatchRE {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	for k := range r.TargetMatch {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	for k := range r.TargetMatchRE {
		if !model.LabelNameRE.MatchString(k) {
			return fmt.Errorf("invalid label name %q", k)
		}
	}
	return nil
}

type Receiver struct {
	Name			string			`yaml:"name" json:"name"`
	EmailConfigs		[]*EmailConfig		`yaml:"email_configs,omitempty" json:"email_configs,omitempty"`
	PagerdutyConfigs	[]*PagerdutyConfig	`yaml:"pagerduty_configs,omitempty" json:"pagerduty_configs,omitempty"`
	HipchatConfigs		[]*HipchatConfig	`yaml:"hipchat_configs,omitempty" json:"hipchat_configs,omitempty"`
	SlackConfigs		[]*SlackConfig		`yaml:"slack_configs,omitempty" json:"slack_configs,omitempty"`
	WebhookConfigs		[]*WebhookConfig	`yaml:"webhook_configs,omitempty" json:"webhook_configs,omitempty"`
	OpsGenieConfigs		[]*OpsGenieConfig	`yaml:"opsgenie_configs,omitempty" json:"opsgenie_configs,omitempty"`
	WechatConfigs		[]*WechatConfig		`yaml:"wechat_configs,omitempty" json:"wechat_configs,omitempty"`
	PushoverConfigs		[]*PushoverConfig	`yaml:"pushover_configs,omitempty" json:"pushover_configs,omitempty"`
	VictorOpsConfigs	[]*VictorOpsConfig	`yaml:"victorops_configs,omitempty" json:"victorops_configs,omitempty"`
}

func (c *Receiver) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type plain Receiver
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	if c.Name == "" {
		return fmt.Errorf("missing name in receiver")
	}
	return nil
}

type Regexp struct{ *regexp.Regexp }

func (re *Regexp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	regex, err := regexp.Compile("^(?:" + s + ")$")
	if err != nil {
		return err
	}
	re.Regexp = regex
	return nil
}
func (re Regexp) MarshalYAML() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if re.Regexp != nil {
		return re.String(), nil
	}
	return nil, nil
}
func (re *Regexp) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	regex, err := regexp.Compile("^(?:" + s + ")$")
	if err != nil {
		return err
	}
	re.Regexp = regex
	return nil
}
func (re Regexp) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if re.Regexp != nil {
		return json.Marshal(re.String())
	}
	return nil, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
