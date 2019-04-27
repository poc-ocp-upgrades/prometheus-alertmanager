package silence

import (
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
	strfmt "github.com/go-openapi/strfmt"
)

func NewDeleteSilenceParams() DeleteSilenceParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return DeleteSilenceParams{}
}

type DeleteSilenceParams struct {
	HTTPRequest	*http.Request	`json:"-"`
	SilenceID	strfmt.UUID
}

func (o *DeleteSilenceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	rSilenceID, rhkSilenceID, _ := route.Params.GetOK("silenceID")
	if err := o.bindSilenceID(rSilenceID, rhkSilenceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (o *DeleteSilenceParams) bindSilenceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("silenceID", "path", "strfmt.UUID", raw)
	}
	o.SilenceID = *(value.(*strfmt.UUID))
	if err := o.validateSilenceID(formats); err != nil {
		return err
	}
	return nil
}
func (o *DeleteSilenceParams) validateSilenceID(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.FormatOf("silenceID", "path", "uuid", o.SilenceID.String(), formats); err != nil {
		return err
	}
	return nil
}
