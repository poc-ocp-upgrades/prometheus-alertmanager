package silence

import (
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	strfmt "github.com/go-openapi/strfmt"
)

func NewGetSilencesParams() GetSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return GetSilencesParams{}
}

type GetSilencesParams struct {
	HTTPRequest	*http.Request	`json:"-"`
	Filter		[]string
}

func (o *GetSilencesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	qs := runtime.Values(r.URL.Query())
	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (o *GetSilencesParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
