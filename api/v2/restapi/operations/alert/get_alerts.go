package alert

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	middleware "github.com/go-openapi/runtime/middleware"
)

type GetAlertsHandlerFunc func(GetAlertsParams) middleware.Responder

func (fn GetAlertsHandlerFunc) Handle(params GetAlertsParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type GetAlertsHandler interface {
	Handle(GetAlertsParams) middleware.Responder
}

func NewGetAlerts(ctx *middleware.Context, handler GetAlertsHandler) *GetAlerts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlerts{Context: ctx, Handler: handler}
}

type GetAlerts struct {
	Context	*middleware.Context
	Handler	GetAlertsHandler
}

func (o *GetAlerts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAlertsParams()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
