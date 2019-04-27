package silence

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
)

type GetSilenceHandlerFunc func(GetSilenceParams) middleware.Responder

func (fn GetSilenceHandlerFunc) Handle(params GetSilenceParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type GetSilenceHandler interface {
	Handle(GetSilenceParams) middleware.Responder
}

func NewGetSilence(ctx *middleware.Context, handler GetSilenceHandler) *GetSilence {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilence{Context: ctx, Handler: handler}
}

type GetSilence struct {
	Context	*middleware.Context
	Handler	GetSilenceHandler
}

func (o *GetSilence) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSilenceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}
