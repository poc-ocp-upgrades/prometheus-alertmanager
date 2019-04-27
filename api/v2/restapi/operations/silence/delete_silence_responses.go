package silence

import (
	"net/http"
	"github.com/go-openapi/runtime"
)

const DeleteSilenceOKCode int = 200

type DeleteSilenceOK struct{}

func NewDeleteSilenceOK() *DeleteSilenceOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeleteSilenceOK{}
}
func (o *DeleteSilenceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.Header().Del(runtime.HeaderContentType)
	rw.WriteHeader(200)
}

const DeleteSilenceInternalServerErrorCode int = 500

type DeleteSilenceInternalServerError struct {
	Payload string `json:"body,omitempty"`
}

func NewDeleteSilenceInternalServerError() *DeleteSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeleteSilenceInternalServerError{}
}
func (o *DeleteSilenceInternalServerError) WithPayload(payload string) *DeleteSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *DeleteSilenceInternalServerError) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *DeleteSilenceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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
