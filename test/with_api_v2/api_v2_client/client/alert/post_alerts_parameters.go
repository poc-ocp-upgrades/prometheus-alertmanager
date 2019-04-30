package alert

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

func NewPostAlertsParams() *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostAlertsParams{timeout: cr.DefaultTimeout}
}
func NewPostAlertsParamsWithTimeout(timeout time.Duration) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostAlertsParams{timeout: timeout}
}
func NewPostAlertsParamsWithContext(ctx context.Context) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostAlertsParams{Context: ctx}
}
func NewPostAlertsParamsWithHTTPClient(client *http.Client) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ()
	return &PostAlertsParams{HTTPClient: client}
}

type PostAlertsParams struct {
	Alerts		models.PostableAlerts
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *PostAlertsParams) WithTimeout(timeout time.Duration) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *PostAlertsParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *PostAlertsParams) WithContext(ctx context.Context) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *PostAlertsParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *PostAlertsParams) WithHTTPClient(client *http.Client) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *PostAlertsParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *PostAlertsParams) WithAlerts(alerts models.PostableAlerts) *PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetAlerts(alerts)
	return o
}
func (o *PostAlertsParams) SetAlerts(alerts models.PostableAlerts) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Alerts = alerts
}
func (o *PostAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Alerts != nil {
		if err := r.SetBodyParam(o.Alerts); err != nil {
			return err
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
