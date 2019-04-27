package receiver

import (
	"net/http"
	"github.com/go-openapi/runtime"
	models "github.com/prometheus/alertmanager/api/v2/models"
)

const GetReceiversOKCode int = 200

type GetReceiversOK struct {
	Payload []*models.Receiver `json:"body,omitempty"`
}

func NewGetReceiversOK() *GetReceiversOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversOK{}
}
func (o *GetReceiversOK) WithPayload(payload []*models.Receiver) *GetReceiversOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *GetReceiversOK) SetPayload(payload []*models.Receiver) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *GetReceiversOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Receiver, 0, 50)
	}
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}
