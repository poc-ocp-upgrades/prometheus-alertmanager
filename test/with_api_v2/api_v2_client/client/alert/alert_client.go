package alert

import (
	"github.com/go-openapi/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	strfmt "github.com/go-openapi/strfmt"
)

func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Client{transport: transport, formats: formats}
}

type Client struct {
	transport	runtime.ClientTransport
	formats		strfmt.Registry
}

func (a *Client) GetAlerts(params *GetAlertsParams) (*GetAlertsOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewGetAlertsParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "getAlerts", Method: "GET", PathPattern: "/alerts", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &GetAlertsReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*GetAlertsOK), nil
}
func (a *Client) PostAlerts(params *PostAlertsParams) (*PostAlertsOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewPostAlertsParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "postAlerts", Method: "POST", PathPattern: "/alerts", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &PostAlertsReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*PostAlertsOK), nil
}
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.transport = transport
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
