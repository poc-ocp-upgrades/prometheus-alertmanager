package general

import (
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetStatusParams() GetStatusParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return GetStatusParams{}
}

type GetStatusParams struct {
	HTTPRequest *http.Request `json:"-"`
}

func (o *GetStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
