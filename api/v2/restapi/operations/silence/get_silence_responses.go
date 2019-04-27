package silence

import (
	"net/http"
	"github.com/go-openapi/runtime"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

const GetSilenceOKCode int = 200

type GetSilenceOK struct {
	Payload *models.GettableSilence `json:"body,omitempty"`
}

func NewGetSilenceOK() *GetSilenceOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceOK{}
}
func (o *GetSilenceOK) WithPayload(payload *models.GettableSilence) *GetSilenceOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetSilenceOK) SetPayload(payload *models.GettableSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetSilenceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err)
		}
	}
}

const GetSilenceNotFoundCode int = 404

type GetSilenceNotFound struct{}

func NewGetSilenceNotFound() *GetSilenceNotFound {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceNotFound{}
}
func (o *GetSilenceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.Header().Del(runtime.HeaderContentType)
	rw.WriteHeader(404)
}

const GetSilenceInternalServerErrorCode int = 500

type GetSilenceInternalServerError struct {
	Payload string `json:"body,omitempty"`
}

func NewGetSilenceInternalServerError() *GetSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceInternalServerError{}
}
func (o *GetSilenceInternalServerError) WithPayload(payload string) *GetSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetSilenceInternalServerError) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetSilenceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
