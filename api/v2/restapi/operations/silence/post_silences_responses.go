package silence

import (
	"net/http"
	"github.com/go-openapi/runtime"
)

const PostSilencesOKCode int = 200

type PostSilencesOK struct {
	Payload *PostSilencesOKBody `json:"body,omitempty"`
}

func NewPostSilencesOK() *PostSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostSilencesOK{}
}
func (o *PostSilencesOK) WithPayload(payload *PostSilencesOKBody) *PostSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *PostSilencesOK) SetPayload(payload *PostSilencesOKBody) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *PostSilencesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
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

const PostSilencesBadRequestCode int = 400

type PostSilencesBadRequest struct {
	Payload string `json:"body,omitempty"`
}

func NewPostSilencesBadRequest() *PostSilencesBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostSilencesBadRequest{}
}
func (o *PostSilencesBadRequest) WithPayload(payload string) *PostSilencesBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
	return o
}
func (o *PostSilencesBadRequest) SetPayload(payload string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = payload
}
func (o *PostSilencesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err)
	}
}
