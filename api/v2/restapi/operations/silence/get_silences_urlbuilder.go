package silence

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"github.com/go-openapi/swag"
)

type GetSilencesURL struct {
	Filter		[]string
	_basePath	string
	_		struct{}
}

func (o *GetSilencesURL) WithBasePath(bp string) *GetSilencesURL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.SetBasePath(bp)
	return o
}
func (o *GetSilencesURL) SetBasePath(bp string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o._basePath = bp
}
func (o *GetSilencesURL) Build() (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result url.URL
	var _path = "/silences"
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)
	qs := make(url.Values)
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
	result.RawQuery = qs.Encode()
	return &result, nil
}
func (o *GetSilencesURL) Must(u *url.URL, err error) *url.URL {
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
func (o *GetSilencesURL) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.Build()).String()
}
func (o *GetSilencesURL) BuildFull(scheme, host string) (*url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetSilencesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetSilencesURL")
	}
	base, err := o.Build()
	if err != nil {
		return nil, err
	}
	base.Scheme = scheme
	base.Host = host
	return base, nil
}
func (o *GetSilencesURL) StringFull(scheme, host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.Must(o.BuildFull(scheme, host)).String()
}
