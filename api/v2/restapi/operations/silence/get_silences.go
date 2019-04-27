package silence

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
)

type GetSilencesHandlerFunc func(GetSilencesParams) middleware.Responder

func (fn GetSilencesHandlerFunc) Handle(params GetSilencesParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type GetSilencesHandler interface {
	Handle(GetSilencesParams) middleware.Responder
}

func NewGetSilences(ctx *middleware.Context, handler GetSilencesHandler) *GetSilences {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilences{Context: ctx, Handler: handler}
}

type GetSilences struct {
	Context	*middleware.Context
	Handler	GetSilencesHandler
}

func (o *GetSilences) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSilencesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}
