package silence

import (
	"net/http"
	"github.com/go-openapi/runtime"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

const GetSilencesOKCode int = 200

type GetSilencesOK struct {
	Payload models.GettableSilences `json:"body,omitempty"`
}

func NewGetSilencesOK() *GetSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilencesOK{}
}
func (o *GetSilencesOK) WithPayload(payload models.GettableSilences) *GetSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetSilencesOK) SetPayload(payload models.GettableSilences) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetSilencesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GettableSilences, 0, 50)
	}
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}

const GetSilencesInternalServerErrorCode int = 500

type GetSilencesInternalServerError struct {
	Payload string `json:"body,omitempty"`
}

func NewGetSilencesInternalServerError() *GetSilencesInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilencesInternalServerError{}
}
func (o *GetSilencesInternalServerError) WithPayload(payload string) *GetSilencesInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetSilencesInternalServerError) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetSilencesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
