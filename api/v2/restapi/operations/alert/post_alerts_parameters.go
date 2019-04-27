package alert

import (
	"io"
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

func NewPostAlertsParams() PostAlertsParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return PostAlertsParams{}
}

type PostAlertsParams struct {
	HTTPRequest	*http.Request	`json:"-"`
	Alerts		models.PostableAlerts
}

func (o *PostAlertsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	o.HTTPRequest = r
	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostableAlerts
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("alerts", "body"))
			} else {
				res = append(res, errors.NewParseError("alerts", "body", "", err))
			}
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}
			if len(res) == 0 {
				o.Alerts = body
			}
		}
	} else {
		res = append(res, errors.Required("alerts", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
