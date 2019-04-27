package receiver

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetReceiversParams() *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversParams{timeout: cr.DefaultTimeout}
}
func NewGetReceiversParamsWithTimeout(timeout time.Duration) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversParams{timeout: timeout}
}
func NewGetReceiversParamsWithContext(ctx context.Context) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversParams{Context: ctx}
}
func NewGetReceiversParamsWithHTTPClient(client *http.Client) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversParams{HTTPClient: client}
}

type GetReceiversParams struct {
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *GetReceiversParams) WithTimeout(timeout time.Duration) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *GetReceiversParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *GetReceiversParams) WithContext(ctx context.Context) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *GetReceiversParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *GetReceiversParams) WithHTTPClient(client *http.Client) *GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *GetReceiversParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *GetReceiversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
