package config

import (
	"encoding/json"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
	commoncfg "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestLoadEmptyString(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var in string
	_, err := Load(in)
	expected := "no route provided in config"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%v", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%v\ngot:\n%v", expected, err.Error())
	}
}
func TestDefaultReceiverExists(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
   group_wait: 30s
`
	_, err := Load(in)
	expected := "root route must specify a default receiver"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%v", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%v\ngot:\n%v", expected, err.Error())
	}
}
func TestReceiverNameIsUnique(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X

receivers:
- name: 'team-X'
- name: 'team-X'
`
	_, err := Load(in)
	expected := "notification config name \"team-X\" is not unique"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestReceiverExists(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X

receivers:
- name: 'team-Y'
`
	_, err := Load(in)
	expected := "undefined receiver \"team-X\" used in route"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestReceiverHasName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:

receivers:
- name: ''
`
	_, err := Load(in)
	expected := "missing name in receiver"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestGroupByHasNoDuplicatedLabels(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
  group_by: ['alertname', 'cluster', 'service', 'cluster']

receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "duplicated label \"cluster\" in group_by"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestWildcardGroupByWithOtherGroupByLabels(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
  group_by: ['alertname', 'cluster', '...']
  receiver: team-X-mails
receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "cannot have wildcard group_by (`...`) and other other labels at the same time"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestGroupByInvalidLabel(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
  group_by: ['-invalid-']
  receiver: team-X-mails
receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "invalid label name \"-invalid-\" in group_by list"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestRootRouteExists(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "no routes provided"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestRootRouteHasNoMatcher(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
  receiver: 'team-X'
  match:
    severity: critical

receivers:
- name: 'team-X'
`
	_, err := Load(in)
	expected := "root route must not have any matchers"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestContinueErrorInRouteRoot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X-mails
    continue: true

receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "cannot have continue in root route"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestGroupIntervalIsGreaterThanZero(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X-mails
    group_interval: 0s

receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "group_interval cannot be zero"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestRepeatIntervalIsGreaterThanZero(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X-mails
    repeat_interval: 0s

receivers:
- name: 'team-X-mails'
`
	_, err := Load(in)
	expected := "repeat_interval cannot be zero"
	if err == nil {
		t.Fatalf("no error returned, expected:\n%q", expected)
	}
	if err.Error() != expected {
		t.Errorf("\nexpected:\n%q\ngot:\n%q", expected, err.Error())
	}
}
func TestHideConfigSecrets(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Fatalf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	s := c.String()
	if strings.Count(s, "<secret>") != 14 || strings.Contains(s, "mysecret") {
		t.Fatal("config's String method reveals authentication credentials.")
	}
}
func TestJSONMarshal(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	_, err = json.Marshal(c)
	if err != nil {
		t.Fatal("JSON Marshaling failed:", err)
	}
}
func TestJSONMarshalSecret(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	test := struct{ S Secret }{S: Secret("test")}
	c, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "{\"S\":\"\\u003csecret\\u003e\"}", string(c), "Secret not properly elided.")
}
func TestMarshalSecretURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urlp, err := url.Parse("http://example.com/")
	if err != nil {
		t.Fatal(err)
	}
	u := &SecretURL{urlp}
	c, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "\"\\u003csecret\\u003e\"", string(c), "SecretURL not properly elided in JSON.")
	out := &SecretURL{}
	err = json.Unmarshal(c, out)
	if err != nil {
		t.Fatal(err)
	}
	c, err = yaml.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "<secret>\n", string(c), "SecretURL not properly elided in YAML.")
	out = &SecretURL{}
	err = yaml.Unmarshal(c, &out)
	if err != nil {
		t.Fatal(err)
	}
}
func TestUnmarshalSecretURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := []byte(`"http://example.com/se cret"`)
	var u SecretURL
	err := json.Unmarshal(b, &u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "http://example.com/se%20cret", u.String(), "SecretURL not properly unmarshalled in JSON.")
	err = yaml.Unmarshal(b, &u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "http://example.com/se%20cret", u.String(), "SecretURL not properly unmarshalled in YAML.")
}
func TestMarshalURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urlp, err := url.Parse("http://example.com/")
	if err != nil {
		t.Fatal(err)
	}
	u := &URL{urlp}
	c, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "\"http://example.com/\"", string(c), "URL not properly marshalled in JSON.")
	c, err = yaml.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "http://example.com/\n", string(c), "URL not properly marshalled in YAML.")
}
func TestUnmarshalURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := []byte(`"http://example.com/a b"`)
	var u URL
	err := json.Unmarshal(b, &u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "http://example.com/a%20b", u.String(), "URL not properly unmarshalled in JSON.")
	err = json.Unmarshal(b, &u)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "http://example.com/a%20b", u.String(), "URL not properly unmarshalled in YAML.")
}
func TestUnmarshalInvalidURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, b := range [][]byte{[]byte(`"://example.com"`), []byte(`"http:example.com"`), []byte(`"telnet://example.com"`)} {
		var u URL
		err := json.Unmarshal(b, &u)
		if err == nil {
			t.Errorf("Expected an error unmarshalling %q from JSON", string(b))
		}
		err = yaml.Unmarshal(b, &u)
		if err == nil {
			t.Errorf("Expected an error unmarshalling %q from YAML", string(b))
		}
		t.Logf("%s", err)
	}
}
func TestUnmarshalRelativeURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := []byte(`"/home"`)
	var u URL
	err := json.Unmarshal(b, &u)
	if err == nil {
		t.Errorf("Expected an error parsing URL")
	}
	err = yaml.Unmarshal(b, &u)
	if err == nil {
		t.Errorf("Expected an error parsing URL")
	}
}
func TestJSONUnmarshal(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	_, err = json.Marshal(c)
	if err != nil {
		t.Fatal("JSON Marshaling failed:", err)
	}
}
func TestMarshalIdempotency(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	marshaled, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal("YAML Marshaling failed:", err)
	}
	c = new(Config)
	if err := yaml.Unmarshal(marshaled, c); err != nil {
		t.Fatal("YAML Unmarshaling failed:", err)
	}
}
func TestGroupByAllNotMarshaled(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := `
route:
    receiver: team-X-mails
    group_by: [...]

receivers:
- name: 'team-X-mails'
`
	c, err := Load(in)
	if err != nil {
		t.Fatal("load failed:", err)
	}
	dat, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal("YAML Marshaling failed:", err)
	}
	if strings.Contains(string(dat), "groupbyall") {
		t.Fatal("groupbyall found in config file")
	}
}
func TestEmptyFieldsAndRegex(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	boolFoo := true
	var regexpFoo Regexp
	regexpFoo.Regexp, _ = regexp.Compile("^(?:^(foo1|foo2|baz)$)$")
	var expectedConf = Config{Global: &GlobalConfig{HTTPConfig: &commoncfg.HTTPClientConfig{}, ResolveTimeout: model.Duration(5 * time.Minute), SMTPSmarthost: "localhost:25", SMTPFrom: "alertmanager@example.org", HipchatAuthToken: "mysecret", HipchatAPIURL: mustParseURL("https://hipchat.foobar.org/"), SlackAPIURL: (*SecretURL)(mustParseURL("http://slack.example.com/")), SMTPRequireTLS: true, PagerdutyURL: mustParseURL("https://events.pagerduty.com/v2/enqueue"), OpsGenieAPIURL: mustParseURL("https://api.opsgenie.com/"), WeChatAPIURL: mustParseURL("https://qyapi.weixin.qq.com/cgi-bin/"), VictorOpsAPIURL: mustParseURL("https://alert.victorops.com/integrations/generic/20131114/alert/")}, Templates: []string{"/etc/alertmanager/template/*.tmpl"}, Route: &Route{Receiver: "team-X-mails", GroupBy: []model.LabelName{"alertname", "cluster", "service"}, GroupByStr: []string{"alertname", "cluster", "service"}, GroupByAll: false, Routes: []*Route{{Receiver: "team-X-mails", MatchRE: map[string]Regexp{"service": regexpFoo}}}}, Receivers: []*Receiver{{Name: "team-X-mails", EmailConfigs: []*EmailConfig{{To: "team-X+alerts@example.org", From: "alertmanager@example.org", Smarthost: "localhost:25", HTML: "{{ template \"email.default.html\" . }}", RequireTLS: &boolFoo}}}}}
	_, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	config, _, err := LoadFile("testdata/conf.empty-fields.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.empty-fields.yml", err)
	}
	configGot, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal("YAML Marshaling failed:", err)
	}
	configExp, err := yaml.Marshal(expectedConf)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if !reflect.DeepEqual(configGot, configExp) {
		t.Fatalf("%s: unexpected config result: \n\n%s\n expected\n\n%s", "testdata/conf.empty-fields.yml", configGot, configExp)
	}
}
func TestSMTPHello(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.good.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.good.yml", err)
	}
	const refValue = "host.example.org"
	var hostName = c.Global.SMTPHello
	if hostName != refValue {
		t.Errorf("Invalid SMTP Hello hostname: %s\nExpected: %s", hostName, refValue)
	}
}
func TestGroupByAll(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, _, err := LoadFile("testdata/conf.group-by-all.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.group-by-all.yml", err)
	}
	if !c.Route.GroupByAll {
		t.Errorf("Invalid group by all param: expected to by true")
	}
}
func TestVictorOpsDefaultAPIKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conf, _, err := LoadFile("testdata/conf.victorops-default-apikey.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.victorops-default-apikey.yml", err)
	}
	var defaultKey = conf.Global.VictorOpsAPIKey
	if defaultKey != conf.Receivers[0].VictorOpsConfigs[0].APIKey {
		t.Errorf("Invalid victorops key: %s\nExpected: %s", conf.Receivers[0].VictorOpsConfigs[0].APIKey, defaultKey)
	}
	if defaultKey == conf.Receivers[1].VictorOpsConfigs[0].APIKey {
		t.Errorf("Invalid victorops key: %s\nExpected: %s", conf.Receivers[0].VictorOpsConfigs[0].APIKey, "qwe456")
	}
}
func TestVictorOpsNoAPIKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, err := LoadFile("testdata/conf.victorops-no-apikey.yml")
	if err == nil {
		t.Errorf("Expected an error parsing %s: %s", "testdata/conf.victorops-no-apikey.yml", err)
	}
	if err.Error() != "no global VictorOps API Key set" {
		t.Errorf("Expected: %s\nGot: %s", "no global VictorOps API Key set", err.Error())
	}
}
func TestOpsGenieDefaultAPIKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conf, _, err := LoadFile("testdata/conf.opsgenie-default-apikey.yml")
	if err != nil {
		t.Errorf("Error parsing %s: %s", "testdata/conf.opsgenie-default-apikey.yml", err)
	}
	var defaultKey = conf.Global.OpsGenieAPIKey
	if defaultKey != conf.Receivers[0].OpsGenieConfigs[0].APIKey {
		t.Errorf("Invalid OpsGenie key: %s\nExpected: %s", conf.Receivers[0].OpsGenieConfigs[0].APIKey, defaultKey)
	}
	if defaultKey == conf.Receivers[1].OpsGenieConfigs[0].APIKey {
		t.Errorf("Invalid OpsGenie key: %s\nExpected: %s", conf.Receivers[0].OpsGenieConfigs[0].APIKey, "qwe456")
	}
}
func TestOpsGenieNoAPIKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, err := LoadFile("testdata/conf.opsgenie-no-apikey.yml")
	if err == nil {
		t.Errorf("Expected an error parsing %s: %s", "testdata/conf.opsgenie-no-apikey.yml", err)
	}
	if err.Error() != "no global OpsGenie API Key set" {
		t.Errorf("Expected: %s\nGot: %s", "no global OpsGenie API Key set", err.Error())
	}
}
