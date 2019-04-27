package alert

import (
	"net/http"
	"github.com/go-openapi/runtime"
)

const PostAlertsOKCode int = 200

type PostAlertsOK struct{}

func NewPostAlertsOK() *PostAlertsOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsOK{}
}
func (o *PostAlertsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.Header().Del(runtime.HeaderContentType)
	rw.WriteHeader(200)
}

const PostAlertsBadRequestCode int = 400

type PostAlertsBadRequest struct {
	Payload string `json:"body,omitempty"`
}

func NewPostAlertsBadRequest() *PostAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsBadRequest{}
}
func (o *PostAlertsBadRequest) WithPayload(payload string) *PostAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *PostAlertsBadRequest) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *PostAlertsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}

const PostAlertsInternalServerErrorCode int = 500

type PostAlertsInternalServerError struct {
	Payload string `json:"body,omitempty"`
}

func NewPostAlertsInternalServerError() *PostAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsInternalServerError{}
}
func (o *PostAlertsInternalServerError) WithPayload(payload string) *PostAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *PostAlertsInternalServerError) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *PostAlertsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}
