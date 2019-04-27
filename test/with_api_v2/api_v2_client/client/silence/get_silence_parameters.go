package silence

import (
	"net/http"
	"time"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetSilenceParams() *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilenceParams{timeout: cr.DefaultTimeout}
}
func NewGetSilenceParamsWithTimeout(timeout time.Duration) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilenceParams{timeout: timeout}
}
func NewGetSilenceParamsWithContext(ctx context.Context) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilenceParams{Context: ctx}
}
func NewGetSilenceParamsWithHTTPClient(client *http.Client) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilenceParams{HTTPClient: client}
}

type GetSilenceParams struct {
	SilenceID	strfmt.UUID
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *GetSilenceParams) WithTimeout(timeout time.Duration) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *GetSilenceParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *GetSilenceParams) WithContext(ctx context.Context) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *GetSilenceParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *GetSilenceParams) WithHTTPClient(client *http.Client) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *GetSilenceParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *GetSilenceParams) WithSilenceID(silenceID strfmt.UUID) *GetSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetSilenceID(silenceID)
	return o
}
func (o *GetSilenceParams) SetSilenceID(silenceID strfmt.UUID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SilenceID = silenceID
}
func (o *GetSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetPathParam("silenceID", o.SilenceID.String()); err != nil {
		return err
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
