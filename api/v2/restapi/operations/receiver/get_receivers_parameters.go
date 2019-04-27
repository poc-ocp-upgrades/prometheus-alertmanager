package receiver

import (
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetReceiversParams() GetReceiversParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return GetReceiversParams{}
}

type GetReceiversParams struct {
	HTTPRequest *http.Request `json:"-"`
}

func (o *GetReceiversParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
