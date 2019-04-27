package alert

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

type PostAlertsURL struct{ _basePath string }

func (o *PostAlertsURL) WithBasePath(bp string) *PostAlertsURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *PostAlertsURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *PostAlertsURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/alerts"
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	return &result, nil
}
func (o *PostAlertsURL) Must(u *url.URL, err error) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (o *PostAlertsURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *PostAlertsURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostAlertsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostAlertsURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *PostAlertsURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
