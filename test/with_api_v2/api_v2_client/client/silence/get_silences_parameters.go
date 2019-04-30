package silence

import (
	"net/http"
	"time"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetSilencesParams() *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilencesParams{timeout: cr.DefaultTimeout}
}
func NewGetSilencesParamsWithTimeout(timeout time.Duration) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilencesParams{timeout: timeout}
}
func NewGetSilencesParamsWithContext(ctx context.Context) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilencesParams{Context: ctx}
}
func NewGetSilencesParamsWithHTTPClient(client *http.Client) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &GetSilencesParams{HTTPClient: client}
}

type GetSilencesParams struct {
	Filter		[]string
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *GetSilencesParams) WithTimeout(timeout time.Duration) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *GetSilencesParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *GetSilencesParams) WithContext(ctx context.Context) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *GetSilencesParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *GetSilencesParams) WithHTTPClient(client *http.Client) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *GetSilencesParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *GetSilencesParams) WithFilter(filter []string) *GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetFilter(filter)
	return o
}
func (o *GetSilencesParams) SetFilter(filter []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Filter = filter
}
func (o *GetSilencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	valuesFilter := o.Filter
	joinedFilter := swag.JoinByFormat(valuesFilter, "")
	if err := r.SetQueryParam("filter", joinedFilter...); err != nil {
		return err
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
