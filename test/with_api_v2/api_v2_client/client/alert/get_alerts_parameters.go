package alert

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

func NewGetAlertsParams() *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		activeDefault		= bool(true)
		inhibitedDefault	= bool(true)
		silencedDefault		= bool(true)
		unprocessedDefault	= bool(true)
	)
	return &GetAlertsParams{Active: &activeDefault, Inhibited: &inhibitedDefault, Silenced: &silencedDefault, Unprocessed: &unprocessedDefault, timeout: cr.DefaultTimeout}
}
func NewGetAlertsParamsWithTimeout(timeout time.Duration) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		activeDefault		= bool(true)
		inhibitedDefault	= bool(true)
		silencedDefault		= bool(true)
		unprocessedDefault	= bool(true)
	)
	return &GetAlertsParams{Active: &activeDefault, Inhibited: &inhibitedDefault, Silenced: &silencedDefault, Unprocessed: &unprocessedDefault, timeout: timeout}
}
func NewGetAlertsParamsWithContext(ctx context.Context) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		activeDefault		= bool(true)
		inhibitedDefault	= bool(true)
		silencedDefault		= bool(true)
		unprocessedDefault	= bool(true)
	)
	return &GetAlertsParams{Active: &activeDefault, Inhibited: &inhibitedDefault, Silenced: &silencedDefault, Unprocessed: &unprocessedDefault, Context: ctx}
}
func NewGetAlertsParamsWithHTTPClient(client *http.Client) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		activeDefault		= bool(true)
		inhibitedDefault	= bool(true)
		silencedDefault		= bool(true)
		unprocessedDefault	= bool(true)
	)
	return &GetAlertsParams{Active: &activeDefault, Inhibited: &inhibitedDefault, Silenced: &silencedDefault, Unprocessed: &unprocessedDefault, HTTPClient: client}
}

type GetAlertsParams struct {
	Active		*bool
	Filter		[]string
	Inhibited	*bool
	Receiver	*string
	Silenced	*bool
	Unprocessed	*bool
	timeout		time.Duration
	Context		context.Context
	HTTPClient	*http.Client
}

func (o *GetAlertsParams) WithTimeout(timeout time.Duration) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetTimeout(timeout)
	return o
}
func (o *GetAlertsParams) SetTimeout(timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.timeout = timeout
}
func (o *GetAlertsParams) WithContext(ctx context.Context) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetContext(ctx)
	return o
}
func (o *GetAlertsParams) SetContext(ctx context.Context) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context = ctx
}
func (o *GetAlertsParams) WithHTTPClient(client *http.Client) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetHTTPClient(client)
	return o
}
func (o *GetAlertsParams) SetHTTPClient(client *http.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.HTTPClient = client
}
func (o *GetAlertsParams) WithActive(active *bool) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetActive(active)
	return o
}
func (o *GetAlertsParams) SetActive(active *bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Active = active
}
func (o *GetAlertsParams) WithFilter(filter []string) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetFilter(filter)
	return o
}
func (o *GetAlertsParams) SetFilter(filter []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Filter = filter
}
func (o *GetAlertsParams) WithInhibited(inhibited *bool) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetInhibited(inhibited)
	return o
}
func (o *GetAlertsParams) SetInhibited(inhibited *bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Inhibited = inhibited
}
func (o *GetAlertsParams) WithReceiver(receiver *string) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetReceiver(receiver)
	return o
}
func (o *GetAlertsParams) SetReceiver(receiver *string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Receiver = receiver
}
func (o *GetAlertsParams) WithSilenced(silenced *bool) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetSilenced(silenced)
	return o
}
func (o *GetAlertsParams) SetSilenced(silenced *bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Silenced = silenced
}
func (o *GetAlertsParams) WithUnprocessed(unprocessed *bool) *GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetUnprocessed(unprocessed)
	return o
}
func (o *GetAlertsParams) SetUnprocessed(unprocessed *bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Unprocessed = unprocessed
}
func (o *GetAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Active != nil {
		var qrActive bool
		if o.Active != nil {
			qrActive = *o.Active
		}
		qActive := swag.FormatBool(qrActive)
		if qActive != "" {
			if err := r.SetQueryParam("active", qActive); err != nil {
				return err
			}
		}
	}
	valuesFilter := o.Filter
	joinedFilter := swag.JoinByFormat(valuesFilter, "")
	if err := r.SetQueryParam("filter", joinedFilter...); err != nil {
		return err
	}
	if o.Inhibited != nil {
		var qrInhibited bool
		if o.Inhibited != nil {
			qrInhibited = *o.Inhibited
		}
		qInhibited := swag.FormatBool(qrInhibited)
		if qInhibited != "" {
			if err := r.SetQueryParam("inhibited", qInhibited); err != nil {
				return err
			}
		}
	}
	if o.Receiver != nil {
		var qrReceiver string
		if o.Receiver != nil {
			qrReceiver = *o.Receiver
		}
		qReceiver := qrReceiver
		if qReceiver != "" {
			if err := r.SetQueryParam("receiver", qReceiver); err != nil {
				return err
			}
		}
	}
	if o.Silenced != nil {
		var qrSilenced bool
		if o.Silenced != nil {
			qrSilenced = *o.Silenced
		}
		qSilenced := swag.FormatBool(qrSilenced)
		if qSilenced != "" {
			if err := r.SetQueryParam("silenced", qSilenced); err != nil {
				return err
			}
		}
	}
	if o.Unprocessed != nil {
		var qrUnprocessed bool
		if o.Unprocessed != nil {
			qrUnprocessed = *o.Unprocessed
		}
		qUnprocessed := swag.FormatBool(qrUnprocessed)
		if qUnprocessed != "" {
			if err := r.SetQueryParam("unprocessed", qUnprocessed); err != nil {
				return err
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
