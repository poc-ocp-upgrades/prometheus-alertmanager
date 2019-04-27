package alert

import (
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetAlertsParams() GetAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		activeDefault		= bool(true)
		inhibitedDefault	= bool(true)
		silencedDefault		= bool(true)
		unprocessedDefault	= bool(true)
	)
	return GetAlertsParams{Active: &activeDefault, Inhibited: &inhibitedDefault, Silenced: &silencedDefault, Unprocessed: &unprocessedDefault}
}

type GetAlertsParams struct {
	HTTPRequest	*http.Request	`json:"-"`
	Active		*bool
	Filter		[]string
	Inhibited	*bool
	Receiver	*string
	Silenced	*bool
	Unprocessed	*bool
}

func (o *GetAlertsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	qs := runtime.Values(r.URL.Query())
	qActive, qhkActive, _ := qs.GetOK("active")
	if err := o.bindActive(qActive, qhkActive, route.Formats); err != nil {
		res = append(res, err)
	}
	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}
	qInhibited, qhkInhibited, _ := qs.GetOK("inhibited")
	if err := o.bindInhibited(qInhibited, qhkInhibited, route.Formats); err != nil {
		res = append(res, err)
	}
	qReceiver, qhkReceiver, _ := qs.GetOK("receiver")
	if err := o.bindReceiver(qReceiver, qhkReceiver, route.Formats); err != nil {
		res = append(res, err)
	}
	qSilenced, qhkSilenced, _ := qs.GetOK("silenced")
	if err := o.bindSilenced(qSilenced, qhkSilenced, route.Formats); err != nil {
		res = append(res, err)
	}
	qUnprocessed, qhkUnprocessed, _ := qs.GetOK("unprocessed")
	if err := o.bindUnprocessed(qUnprocessed, qhkUnprocessed, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (o *GetAlertsParams) bindActive(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" {
		return nil
	}
	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("active", "query", "bool", raw)
	}
	o.Active = &value
	return nil
}
func (o *GetAlertsParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var qvFilter string
	if len(rawData) > 0 {
		qvFilter = rawData[len(rawData)-1]
	}
	filterIC := swag.SplitByFormat(qvFilter, "")
	if len(filterIC) == 0 {
		return nil
	}
	var filterIR []string
	for _, filterIV := range filterIC {
		filterI := filterIV
		filterIR = append(filterIR, filterI)
	}
	o.Filter = filterIR
	return nil
}
func (o *GetAlertsParams) bindInhibited(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" {
		return nil
	}
	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("inhibited", "query", "bool", raw)
	}
	o.Inhibited = &value
	return nil
}
func (o *GetAlertsParams) bindReceiver(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" {
		return nil
	}
	o.Receiver = &raw
	return nil
}
func (o *GetAlertsParams) bindSilenced(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" {
		return nil
	}
	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("silenced", "query", "bool", raw)
	}
	o.Silenced = &value
	return nil
}
func (o *GetAlertsParams) bindUnprocessed(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" {
		return nil
	}
	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("unprocessed", "query", "bool", raw)
	}
	o.Unprocessed = &value
	return nil
}
