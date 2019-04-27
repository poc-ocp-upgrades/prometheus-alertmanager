package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
	"github.com/go-kit/kit/log"
	commoncfg "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/template"
	"github.com/prometheus/alertmanager/types"
)

func getContextWithCancelingURL(h ...func(w http.ResponseWriter, r *http.Request)) (context.Context, *url.URL, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	i := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if i < len(h) {
			h[i](w, r)
		} else {
			cancel()
			<-done
		}
		i++
	}))
	u, _ := url.Parse(srv.URL)
	return ctx, u, func() {
		close(done)
		srv.Close()
	}
}
func assertNotifyLeaksNoSecret(t *testing.T, ctx context.Context, n Notifier, secret ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Helper()
	require.NotEmpty(t, secret)
	ctx = WithGroupKey(ctx, "1")
	ok, err := n.Notify(ctx, []*types.Alert{&types.Alert{Alert: model.Alert{Labels: model.LabelSet{"lbl1": "val1"}, StartsAt: time.Now(), EndsAt: time.Now().Add(time.Hour)}}}...)
	require.Error(t, err)
	require.Contains(t, err.Error(), context.Canceled.Error())
	for _, s := range secret {
		require.NotContains(t, err.Error(), s)
	}
	require.True(t, ok)
}
func TestWebhookRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("failed to parse URL: %v", err)
	}
	notifier := &Webhook{conf: &config.WebhookConfig{URL: &config.URL{u}}}
	for statusCode, expected := range retryTests(defaultRetryCodes()) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestPagerDutyRetryV1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(PagerDuty)
	retryCodes := append(defaultRetryCodes(), http.StatusForbidden)
	for statusCode, expected := range retryTests(retryCodes) {
		resp := &http.Response{StatusCode: statusCode}
		actual, _ := notifier.retryV1(resp)
		require.Equal(t, expected, actual, fmt.Sprintf("retryv1 - error on status %d", statusCode))
	}
}
func TestPagerDutyRetryV2(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(PagerDuty)
	retryCodes := append(defaultRetryCodes(), http.StatusTooManyRequests)
	for statusCode, expected := range retryTests(retryCodes) {
		actual, _ := notifier.retryV2(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("retryv2 - error on status %d", statusCode))
	}
}
func TestPagerDutyRedactedURLV1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	key := "01234567890123456789012345678901"
	notifier := NewPagerDuty(&config.PagerdutyConfig{ServiceKey: config.Secret(key), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	notifier.apiV1 = u.String()
	assertNotifyLeaksNoSecret(t, ctx, notifier, key)
}
func TestPagerDutyRedactedURLV2(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	key := "01234567890123456789012345678901"
	notifier := NewPagerDuty(&config.PagerdutyConfig{URL: &config.URL{u}, RoutingKey: config.Secret(key), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, key)
}
func TestSlackRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(Slack)
	for statusCode, expected := range retryTests(defaultRetryCodes()) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestSlackRedactedURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	notifier := NewSlack(&config.SlackConfig{APIURL: &config.SecretURL{URL: u}, HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, u.String())
}
func TestHipchatRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(Hipchat)
	retryCodes := append(defaultRetryCodes(), http.StatusTooManyRequests)
	for statusCode, expected := range retryTests(retryCodes) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestHipchatRedactedURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	token := "secret_token"
	notifier := NewHipchat(&config.HipchatConfig{APIURL: &config.URL{URL: u}, AuthToken: config.Secret(token), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, token)
}
func TestOpsGenieRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(OpsGenie)
	retryCodes := append(defaultRetryCodes(), http.StatusTooManyRequests)
	for statusCode, expected := range retryTests(retryCodes) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestOpsGenieRedactedURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	key := "key"
	notifier := NewOpsGenie(&config.OpsGenieConfig{APIURL: &config.URL{URL: u}, APIKey: config.Secret(key), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, key)
}
func TestVictorOpsRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(VictorOps)
	for statusCode, expected := range retryTests(defaultRetryCodes()) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestVictorOpsRedactedURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	secret := "secret"
	notifier := NewVictorOps(&config.VictorOpsConfig{APIURL: &config.URL{URL: u}, APIKey: config.Secret(secret), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, secret)
}
func TestPushoverRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifier := new(Pushover)
	for statusCode, expected := range retryTests(defaultRetryCodes()) {
		actual, _ := notifier.retry(statusCode)
		require.Equal(t, expected, actual, fmt.Sprintf("error on status %d", statusCode))
	}
}
func TestPushoverRedactedURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	key, token := "user_key", "token"
	notifier := NewPushover(&config.PushoverConfig{UserKey: config.Secret(key), Token: config.Secret(token), HTTPConfig: &commoncfg.HTTPClientConfig{}}, createTmpl(t), log.NewNopLogger())
	notifier.apiURL = u.String()
	assertNotifyLeaksNoSecret(t, ctx, notifier, key, token)
}
func retryTests(retryCodes []int) map[int]bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := map[int]bool{http.StatusContinue: false, http.StatusSwitchingProtocols: false, http.StatusProcessing: false, http.StatusOK: false, http.StatusCreated: false, http.StatusAccepted: false, http.StatusNonAuthoritativeInfo: false, http.StatusNoContent: false, http.StatusResetContent: false, http.StatusPartialContent: false, http.StatusMultiStatus: false, http.StatusAlreadyReported: false, http.StatusIMUsed: false, http.StatusMultipleChoices: false, http.StatusMovedPermanently: false, http.StatusFound: false, http.StatusSeeOther: false, http.StatusNotModified: false, http.StatusUseProxy: false, http.StatusTemporaryRedirect: false, http.StatusPermanentRedirect: false, http.StatusBadRequest: false, http.StatusUnauthorized: false, http.StatusPaymentRequired: false, http.StatusForbidden: false, http.StatusNotFound: false, http.StatusMethodNotAllowed: false, http.StatusNotAcceptable: false, http.StatusProxyAuthRequired: false, http.StatusRequestTimeout: false, http.StatusConflict: false, http.StatusGone: false, http.StatusLengthRequired: false, http.StatusPreconditionFailed: false, http.StatusRequestEntityTooLarge: false, http.StatusRequestURITooLong: false, http.StatusUnsupportedMediaType: false, http.StatusRequestedRangeNotSatisfiable: false, http.StatusExpectationFailed: false, http.StatusTeapot: false, http.StatusUnprocessableEntity: false, http.StatusLocked: false, http.StatusFailedDependency: false, http.StatusUpgradeRequired: false, http.StatusPreconditionRequired: false, http.StatusTooManyRequests: false, http.StatusRequestHeaderFieldsTooLarge: false, http.StatusUnavailableForLegalReasons: false, http.StatusInternalServerError: false, http.StatusNotImplemented: false, http.StatusBadGateway: false, http.StatusServiceUnavailable: false, http.StatusGatewayTimeout: false, http.StatusHTTPVersionNotSupported: false, http.StatusVariantAlsoNegotiates: false, http.StatusInsufficientStorage: false, http.StatusLoopDetected: false, http.StatusNotExtended: false, http.StatusNetworkAuthenticationRequired: false}
	for _, statusCode := range retryCodes {
		tests[statusCode] = true
	}
	return tests
}
func defaultRetryCodes() []int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []int{http.StatusInternalServerError, http.StatusNotImplemented, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout, http.StatusHTTPVersionNotSupported, http.StatusVariantAlsoNegotiates, http.StatusInsufficientStorage, http.StatusLoopDetected, http.StatusNotExtended, http.StatusNetworkAuthenticationRequired}
}
func createTmpl(t *testing.T) *template.Template {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tmpl, err := template.FromGlobs()
	require.NoError(t, err)
	tmpl.ExternalURL, _ = url.Parse("http://am")
	return tmpl
}
func readBody(t *testing.T, r *http.Request) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	body, err := ioutil.ReadAll(r.Body)
	require.NoError(t, err)
	return string(body)
}
func TestOpsGenie(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse("https://opsgenie/api")
	if err != nil {
		t.Fatalf("failed to parse URL: %v", err)
	}
	logger := log.NewNopLogger()
	tmpl := createTmpl(t)
	conf := &config.OpsGenieConfig{NotifierConfig: config.NotifierConfig{VSendResolved: true}, Message: `{{ .CommonLabels.Message }}`, Description: `{{ .CommonLabels.Description }}`, Source: `{{ .CommonLabels.Source }}`, Teams: `{{ .CommonLabels.Teams }}`, Tags: `{{ .CommonLabels.Tags }}`, Note: `{{ .CommonLabels.Note }}`, Priority: `{{ .CommonLabels.Priority }}`, APIKey: `{{ .ExternalURL }}`, APIURL: &config.URL{u}}
	notifier := NewOpsGenie(conf, tmpl, logger)
	ctx := context.Background()
	ctx = WithGroupKey(ctx, "1")
	expectedURL, _ := url.Parse("https://opsgenie/apiv2/alerts")
	alert1 := &types.Alert{Alert: model.Alert{StartsAt: time.Now(), EndsAt: time.Now().Add(time.Hour)}}
	expectedBody := `{"alias":"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b","message":"","details":{},"source":""}
`
	req, retry, err := notifier.createRequest(ctx, alert1)
	require.NoError(t, err)
	require.Equal(t, true, retry)
	require.Equal(t, expectedURL, req.URL)
	require.Equal(t, "GenieKey http://am", req.Header.Get("Authorization"))
	require.Equal(t, expectedBody, readBody(t, req))
	alert2 := &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"Message": "message", "Description": "description", "Source": "http://prometheus", "Teams": "TeamA,TeamB,", "Tags": "tag1,tag2", "Note": "this is a note", "Priotity": "P1"}, StartsAt: time.Now(), EndsAt: time.Now().Add(time.Hour)}}
	expectedBody = `{"alias":"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b","message":"message","description":"description","details":{},"source":"http://prometheus","teams":[{"name":"TeamA"},{"name":"TeamB"}],"tags":["tag1","tag2"],"note":"this is a note"}
`
	req, retry, err = notifier.createRequest(ctx, alert2)
	require.NoError(t, err)
	require.Equal(t, true, retry)
	require.Equal(t, expectedBody, readBody(t, req))
	conf.APIKey = "{{ kaput "
	_, _, err = notifier.createRequest(ctx, alert2)
	require.Error(t, err)
	require.Equal(t, err.Error(), "templating error: template: :1: function \"kaput\" not defined")
}
func TestEmailConfigNoAuthMechs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	email := &Email{conf: &config.EmailConfig{AuthUsername: "test"}, tmpl: &template.Template{}, logger: log.NewNopLogger()}
	_, err := email.auth("")
	require.Error(t, err)
	require.Equal(t, err.Error(), "unknown auth mechanism: ")
}
func TestEmailConfigMissingAuthParam(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	conf := &config.EmailConfig{AuthUsername: "test"}
	email := &Email{conf: conf, tmpl: &template.Template{}, logger: log.NewNopLogger()}
	_, err := email.auth("CRAM-MD5")
	require.Error(t, err)
	require.Equal(t, err.Error(), "missing secret for CRAM-MD5 auth mechanism")
	_, err = email.auth("PLAIN")
	require.Error(t, err)
	require.Equal(t, err.Error(), "missing password for PLAIN auth mechanism")
	_, err = email.auth("LOGIN")
	require.Error(t, err)
	require.Equal(t, err.Error(), "missing password for LOGIN auth mechanism")
	_, err = email.auth("PLAIN LOGIN")
	require.Error(t, err)
	require.Equal(t, err.Error(), "missing password for PLAIN auth mechanism; missing password for LOGIN auth mechanism")
}
func TestEmailNoUsernameStillOk(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	email := &Email{conf: &config.EmailConfig{}, tmpl: &template.Template{}, logger: log.NewNopLogger()}
	a, err := email.auth("CRAM-MD5")
	require.NoError(t, err)
	require.Nil(t, a)
}
func TestVictorOpsCustomFields(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger := log.NewNopLogger()
	tmpl := createTmpl(t)
	url, err := url.Parse("http://nowhere.com")
	require.NoError(t, err, "unexpected error parsing mock url")
	conf := &config.VictorOpsConfig{APIKey: `12345`, APIURL: &config.URL{url}, EntityDisplayName: `{{ .CommonLabels.Message }}`, StateMessage: `{{ .CommonLabels.Message }}`, RoutingKey: `test`, MessageType: ``, MonitoringTool: `AM`, CustomFields: map[string]string{"Field_A": "{{ .CommonLabels.Message }}"}}
	notifier := NewVictorOps(conf, tmpl, logger)
	ctx := context.Background()
	ctx = WithGroupKey(ctx, "1")
	alert := &types.Alert{Alert: model.Alert{Labels: model.LabelSet{"Message": "message"}, StartsAt: time.Now(), EndsAt: time.Now().Add(time.Hour)}}
	msg, err := notifier.createVictorOpsPayload(ctx, alert)
	require.NoError(t, err)
	var m map[string]string
	err = json.Unmarshal(msg.Bytes(), &m)
	require.NoError(t, err)
	require.Equal(t, "message", m["Field_A"])
}
func TestWechatRedactedURLOnInitialAuthentication(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, u, fn := getContextWithCancelingURL()
	defer fn()
	secret := "secret_key"
	notifier := NewWechat(&config.WechatConfig{APIURL: &config.URL{URL: u}, HTTPConfig: &commoncfg.HTTPClientConfig{}, CorpID: "corpid", APISecret: config.Secret(secret)}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, secret)
}
func TestWechatRedactedURLOnNotify(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	secret, token := "secret", "token"
	ctx, u, fn := getContextWithCancelingURL(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"access_token":"%s"}`, token)
	})
	defer fn()
	notifier := NewWechat(&config.WechatConfig{APIURL: &config.URL{URL: u}, HTTPConfig: &commoncfg.HTTPClientConfig{}, CorpID: "corpid", APISecret: config.Secret(secret)}, createTmpl(t), log.NewNopLogger())
	assertNotifyLeaksNoSecret(t, ctx, notifier, secret, token)
}
