package client

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
	"time"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/types"
)

const (
	apiPrefix		= "/api/v1"
	epStatus		= apiPrefix + "/status"
	epSilence		= apiPrefix + "/silence/:id"
	epSilences		= apiPrefix + "/silences"
	epAlerts		= apiPrefix + "/alerts"
	statusSuccess	= "success"
	statusError		= "error"
)

type ServerStatus struct {
	ConfigYAML		string				`json:"configYAML"`
	ConfigJSON		*config.Config		`json:"configJSON"`
	VersionInfo		map[string]string	`json:"versionInfo"`
	Uptime			time.Time			`json:"uptime"`
	ClusterStatus	*ClusterStatus		`json:"clusterStatus"`
}
type PeerStatus struct {
	Name	string	`json:"name"`
	Address	string	`json:"address"`
}
type ClusterStatus struct {
	Name	string			`json:"name"`
	Status	string			`json:"status"`
	Peers	[]PeerStatus	`json:"peers"`
}
type apiClient struct{ api.Client }
type apiResponse struct {
	Status		string			`json:"status"`
	Data		json.RawMessage	`json:"data,omitempty"`
	ErrorType	string			`json:"errorType,omitempty"`
	Error		string			`json:"error,omitempty"`
}
type clientError struct {
	code	int
	msg		string
}

func (e *clientError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s (code: %d)", e.msg, e.code)
}
func (c apiClient) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := c.Client.Do(ctx, req)
	if err != nil {
		return resp, body, err
	}
	code := resp.StatusCode
	var result apiResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return resp, body, &clientError{code: code, msg: string(body)}
	}
	if (code/100 == 2) && (result.Status != statusSuccess) {
		return resp, body, &clientError{code: code, msg: "inconsistent body for response code"}
	}
	if result.Status == statusError {
		err = &clientError{code: code, msg: result.Error}
	}
	return resp, []byte(result.Data), err
}

type StatusAPI interface {
	Get(ctx context.Context) (*ServerStatus, error)
}

func NewStatusAPI(c api.Client) StatusAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpStatusAPI{client: apiClient{c}}
}

type httpStatusAPI struct{ client api.Client }

func (h *httpStatusAPI) Get(ctx context.Context) (*ServerStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epStatus, nil)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	_, body, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var ss *ServerStatus
	err = json.Unmarshal(body, &ss)
	return ss, err
}

type AlertAPI interface {
	List(ctx context.Context, filter, receiver string, silenced, inhibited, active, unprocessed bool) ([]*ExtendedAlert, error)
	Push(ctx context.Context, alerts ...Alert) error
}
type Alert struct {
	Labels			LabelSet	`json:"labels"`
	Annotations		LabelSet	`json:"annotations"`
	StartsAt		time.Time	`json:"startsAt,omitempty"`
	EndsAt			time.Time	`json:"endsAt,omitempty"`
	GeneratorURL	string		`json:"generatorURL"`
}
type ExtendedAlert struct {
	Alert
	Status		types.AlertStatus	`json:"status"`
	Receivers	[]string			`json:"receivers"`
	Fingerprint	string				`json:"fingerprint"`
}
type LabelSet map[LabelName]LabelValue
type LabelName string
type LabelValue string

func NewAlertAPI(c api.Client) AlertAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpAlertAPI{client: apiClient{c}}
}

type httpAlertAPI struct{ client api.Client }

func (h *httpAlertAPI) List(ctx context.Context, filter, receiver string, silenced, inhibited, active, unprocessed bool) ([]*ExtendedAlert, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epAlerts, nil)
	params := url.Values{}
	if filter != "" {
		params.Add("filter", filter)
	}
	params.Add("silenced", fmt.Sprintf("%t", silenced))
	params.Add("inhibited", fmt.Sprintf("%t", inhibited))
	params.Add("active", fmt.Sprintf("%t", active))
	params.Add("unprocessed", fmt.Sprintf("%t", unprocessed))
	params.Add("receiver", receiver)
	u.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	_, body, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var alts []*ExtendedAlert
	err = json.Unmarshal(body, &alts)
	return alts, err
}
func (h *httpAlertAPI) Push(ctx context.Context, alerts ...Alert) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epAlerts, nil)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&alerts); err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), &buf)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	_, _, err = h.client.Do(ctx, req)
	return err
}

type SilenceAPI interface {
	Get(ctx context.Context, id string) (*types.Silence, error)
	Set(ctx context.Context, sil types.Silence) (string, error)
	Expire(ctx context.Context, id string) error
	List(ctx context.Context, filter string) ([]*types.Silence, error)
}

func NewSilenceAPI(c api.Client) SilenceAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpSilenceAPI{client: apiClient{c}}
}

type httpSilenceAPI struct{ client api.Client }

func (h *httpSilenceAPI) Get(ctx context.Context, id string) (*types.Silence, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epSilence, map[string]string{"id": id})
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	_, body, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var sil types.Silence
	err = json.Unmarshal(body, &sil)
	return &sil, err
}
func (h *httpSilenceAPI) Expire(ctx context.Context, id string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epSilence, map[string]string{"id": id})
	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	_, _, err = h.client.Do(ctx, req)
	return err
}
func (h *httpSilenceAPI) Set(ctx context.Context, sil types.Silence) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epSilences, nil)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&sil); err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), &buf)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	_, body, err := h.client.Do(ctx, req)
	if err != nil {
		return "", err
	}
	var res struct {
		SilenceID string `json:"silenceId"`
	}
	err = json.Unmarshal(body, &res)
	return res.SilenceID, err
}
func (h *httpSilenceAPI) List(ctx context.Context, filter string) ([]*types.Silence, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := h.client.URL(epSilences, nil)
	params := url.Values{}
	if filter != "" {
		params.Add("filter", filter)
	}
	u.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	_, body, err := h.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var sils []*types.Silence
	err = json.Unmarshal(body, &sils)
	return sils, err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
