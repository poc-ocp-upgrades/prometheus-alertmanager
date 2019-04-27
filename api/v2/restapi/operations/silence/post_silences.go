package silence

import (
	"net/http"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

type PostSilencesHandlerFunc func(PostSilencesParams) middleware.Responder

func (fn PostSilencesHandlerFunc) Handle(params PostSilencesParams) middleware.Responder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fn(params)
}

type PostSilencesHandler interface {
	Handle(PostSilencesParams) middleware.Responder
}

func NewPostSilences(ctx *middleware.Context, handler PostSilencesHandler) *PostSilences {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostSilences{Context: ctx, Handler: handler}
}

type PostSilences struct {
	Context	*middleware.Context
	Handler	PostSilencesHandler
}

func (o *PostSilences) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostSilencesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	res := o.Handler.Handle(Params)
	o.Context.Respond(rw, r, route.Produces, route, res)
}

type PostSilencesOKBody struct {
	SilenceID string `json:"silenceID,omitempty"`
}

func (o *PostSilencesOKBody) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (o *PostSilencesOKBody) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}
func (o *PostSilencesOKBody) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res PostSilencesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
