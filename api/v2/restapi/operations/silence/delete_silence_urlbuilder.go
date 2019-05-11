package silence

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
	"github.com/go-openapi/strfmt"
)

type DeleteSilenceURL struct {
	SilenceID	strfmt.UUID
	_basePath	string
	_			struct{}
}

func (o *DeleteSilenceURL) WithBasePath(bp string) *DeleteSilenceURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *DeleteSilenceURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *DeleteSilenceURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/silence/{silenceID}"
	silenceID := o.SilenceID.String()
	if silenceID != "" {
		_path = strings.Replace(_path, "{silenceID}", silenceID, -1)
	} else {
		return nil, errors.New("SilenceID is required on DeleteSilenceURL")
	}
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	return &result, nil
}
func (o *DeleteSilenceURL) Must(u *url.URL, err error) *url.URL {
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
func (o *DeleteSilenceURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *DeleteSilenceURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteSilenceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteSilenceURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *DeleteSilenceURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
