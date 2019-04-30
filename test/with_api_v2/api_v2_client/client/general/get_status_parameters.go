package general

import (
	"net/http"
	"time"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetStatusParams() *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusParams{timeout: cr.DefaultTimeout}
}
func NewGetStatusParamsWithTimeout(timeout time.Duration) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusParams{timeout: timeout}
}
func NewGetStatusParamsWithContext(ctx context.Context) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusParams{Context: ctx}
}
func NewGetStatusParamsWithHTTPClient(client *http.Client) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusParams{HTTPClient: client}
}

type GetStatusParams struct {
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *GetStatusParams) WithTimeout(timeout time.Duration) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *GetStatusParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *GetStatusParams) WithContext(ctx context.Context) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *GetStatusParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *GetStatusParams) WithHTTPClient(client *http.Client) *GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *GetStatusParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *GetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
