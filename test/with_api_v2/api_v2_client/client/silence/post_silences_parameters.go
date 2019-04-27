package silence

import (
	"net/http"
	"time"
	"golang.org/x/net/context"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

func NewPostSilencesParams() *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostSilencesParams{timeout: cr.DefaultTimeout}
}
func NewPostSilencesParamsWithTimeout(timeout time.Duration) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostSilencesParams{timeout: timeout}
}
func NewPostSilencesParamsWithContext(ctx context.Context) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostSilencesParams{Context: ctx}
}
func NewPostSilencesParamsWithHTTPClient(client *http.Client) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostSilencesParams{HTTPClient: client}
}

type PostSilencesParams struct {
	Silence		*models.PostableSilence
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *PostSilencesParams) WithTimeout(timeout time.Duration) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *PostSilencesParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *PostSilencesParams) WithContext(ctx context.Context) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *PostSilencesParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *PostSilencesParams) WithHTTPClient(client *http.Client) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *PostSilencesParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *PostSilencesParams) WithSilence(silence *models.PostableSilence) *PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetSilence(silence)
	return o
}
func (o *PostSilencesParams) SetSilence(silence *models.PostableSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Silence = silence
}
func (o *PostSilencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Silence != nil {
		if err := r.SetBodyParam(o.Silence); err != nil {
			return err
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
