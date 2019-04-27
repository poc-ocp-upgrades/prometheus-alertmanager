package alert

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
)

type PostAlertsHandlerFunc func(PostAlertsParams) middleware.Responder

func (fn PostAlertsHandlerFunc) Handle(params PostAlertsParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type PostAlertsHandler interface {
	Handle(PostAlertsParams) middleware.Responder
}

func NewPostAlerts(ctx *middleware.Context, handler PostAlertsHandler) *PostAlerts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlerts{Context: ctx, Handler: handler}
}

type PostAlerts struct {
	Context	*middleware.Context
	Handler	PostAlertsHandler
}

func (o *PostAlerts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAlertsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}
