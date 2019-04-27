package silence

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Client{transport: transport, formats: formats}
}

type Client struct {
	transport	runtime.ClientTransport
	formats		strfmt.Registry
}

func (a *Client) DeleteSilence(params *DeleteSilenceParams) (*DeleteSilenceOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewDeleteSilenceParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "deleteSilence", Method: "DELETE", PathPattern: "/silence/{silenceID}", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &DeleteSilenceReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSilenceOK), nil
}
func (a *Client) GetSilence(params *GetSilenceParams) (*GetSilenceOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewGetSilenceParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "getSilence", Method: "GET", PathPattern: "/silence/{silenceID}", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &GetSilenceReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*GetSilenceOK), nil
}
func (a *Client) GetSilences(params *GetSilencesParams) (*GetSilencesOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewGetSilencesParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "getSilences", Method: "GET", PathPattern: "/silences", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &GetSilencesReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*GetSilencesOK), nil
}
func (a *Client) PostSilences(params *PostSilencesParams) (*PostSilencesOK, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params == nil {
		params = NewPostSilencesParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{ID: "postSilences", Method: "POST", PathPattern: "/silences", ProducesMediaTypes: []string{"application/json"}, ConsumesMediaTypes: []string{"application/json"}, Schemes: []string{"http"}, Params: params, Reader: &PostSilencesReader{formats: a.formats}, Context: params.Context, Client: params.HTTPClient})
	if err != nil {
		return nil, err
	}
	return result.(*PostSilencesOK), nil
}
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.transport = transport
}
