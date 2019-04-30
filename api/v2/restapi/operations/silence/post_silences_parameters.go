package silence

import (
	"io"
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

func NewPostSilencesParams() PostSilencesParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PostSilencesParams{}
}

type PostSilencesParams struct {
	HTTPRequest	*http.Request	`json:"-"`
	Silence		*models.PostableSilence
}

func (o *PostSilencesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostableSilence
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("silence", "body"))
			} else {
				res = append(res, errors.NewParseError("silence", "body", "", err))
			}
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}
			if len(res) == 0 {
				o.Silence = &body
			}
		}
	} else {
		res = append(res, errors.Required("silence", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
