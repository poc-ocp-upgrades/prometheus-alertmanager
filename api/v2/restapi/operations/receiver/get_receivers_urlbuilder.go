package receiver

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

type GetReceiversURL struct{ _basePath string }

func (o *GetReceiversURL) WithBasePath(bp string) *GetReceiversURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *GetReceiversURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *GetReceiversURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/receivers"
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	return &result, nil
}
func (o *GetReceiversURL) Must(u *url.URL, err error) *url.URL {
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
func (o *GetReceiversURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *GetReceiversURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetReceiversURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetReceiversURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *GetReceiversURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
