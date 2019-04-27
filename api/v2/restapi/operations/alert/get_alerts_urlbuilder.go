package alert

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"github.com/go-openapi/swag"
)

type GetAlertsURL struct {
	Active		*bool
	Filter		[]string
	Inhibited	*bool
	Receiver	*string
	Silenced	*bool
	Unprocessed	*bool
	_basePath	string
	_		struct{}
}

func (o *GetAlertsURL) WithBasePath(bp string) *GetAlertsURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *GetAlertsURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *GetAlertsURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/alerts"
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	qs := make(url.Values)
	var active string
	if o.Active != nil {
		active = swag.FormatBool(*o.Active)
	}
	if active != "" {
		qs.Set("active", active)
	}
	var filterIR []string
	for _, filterI := range o.Filter {
		filterIS := filterI
		if filterIS != "" {
			filterIR = append(filterIR, filterIS)
		}
	}
	filter := swag.JoinByFormat(filterIR, "")
	if len(filter) > 0 {
		qsv := filter[0]
		if qsv != "" {
			qs.Set("filter", qsv)
		}
	}
	var inhibited string
	if o.Inhibited != nil {
		inhibited = swag.FormatBool(*o.Inhibited)
	}
	if inhibited != "" {
		qs.Set("inhibited", inhibited)
	}
	var receiver string
	if o.Receiver != nil {
		receiver = *o.Receiver
	}
	if receiver != "" {
		qs.Set("receiver", receiver)
	}
	var silenced string
	if o.Silenced != nil {
		silenced = swag.FormatBool(*o.Silenced)
	}
	if silenced != "" {
		qs.Set("silenced", silenced)
	}
	var unprocessed string
	if o.Unprocessed != nil {
		unprocessed = swag.FormatBool(*o.Unprocessed)
	}
	if unprocessed != "" {
		qs.Set("unprocessed", unprocessed)
	}
	result.RawQuery = qs.Encode()
	return &result, nil
}
func (o *GetAlertsURL) Must(u *url.URL, err error) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}
func (o *GetAlertsURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *GetAlertsURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAlertsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAlertsURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *GetAlertsURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
