package receiver

import (
	"github.com/go-openapi/runtime"
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

func (a *Client) GetReceivers(params *GetReceiversParams) (*GetReceiversOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewGetReceiversParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "getReceivers", Method: "GET", PathPattern: "/receivers", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &GetReceiversReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*GetReceiversOK), nil
}
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.transport = transport
}
