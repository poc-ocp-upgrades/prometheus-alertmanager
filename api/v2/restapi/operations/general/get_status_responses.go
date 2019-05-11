package general

import (
	"net/http"
	"github.com/go-openapi/runtime"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

const GetStatusOKCode int = 200

type GetStatusOK struct {
	Payload *models.AlertmanagerStatus `json:"body,omitempty"`
}

func NewGetStatusOK() *GetStatusOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusOK{}
}
func (o *GetStatusOK) WithPayload(payload *models.AlertmanagerStatus) *GetStatusOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetStatusOK) SetPayload(payload *models.AlertmanagerStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
