package silence

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

type PostSilencesURL struct{ _basePath string }

func (o *PostSilencesURL) WithBasePath(bp string) *PostSilencesURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *PostSilencesURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *PostSilencesURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/silences"
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	return &result, nil
}
func (o *PostSilencesURL) Must(u *url.URL, err error) *url.URL {
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
func (o *PostSilencesURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *PostSilencesURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostSilencesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostSilencesURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *PostSilencesURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
