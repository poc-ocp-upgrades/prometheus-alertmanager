package alert

import (
	"net/http"
	"github.com/go-openapi/runtime"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

const GetAlertsOKCode int = 200

type GetAlertsOK struct {
	Payload models.GettableAlerts `json:"body,omitempty"`
}

func NewGetAlertsOK() *GetAlertsOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsOK{}
}
func (o *GetAlertsOK) WithPayload(payload models.GettableAlerts) *GetAlertsOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetAlertsOK) SetPayload(payload models.GettableAlerts) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetAlertsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GettableAlerts, 0, 50)
	}
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}

const GetAlertsBadRequestCode int = 400

type GetAlertsBadRequest struct {
	Payload string `json:"body,omitempty"`
}

func NewGetAlertsBadRequest() *GetAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsBadRequest{}
}
func (o *GetAlertsBadRequest) WithPayload(payload string) *GetAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetAlertsBadRequest) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetAlertsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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

const GetAlertsInternalServerErrorCode int = 500

type GetAlertsInternalServerError struct {
	Payload string `json:"body,omitempty"`
}

func NewGetAlertsInternalServerError() *GetAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsInternalServerError{}
}
func (o *GetAlertsInternalServerError) WithPayload(payload string) *GetAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetAlertsInternalServerError) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetAlertsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
