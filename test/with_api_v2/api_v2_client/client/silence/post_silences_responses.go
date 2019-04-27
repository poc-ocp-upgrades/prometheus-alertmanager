package silence

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	strfmt "github.com/go-openapi/strfmt"
)

type PostSilencesReader struct{ formats strfmt.Registry }

func (o *PostSilencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewPostSilencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSilencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewPostSilencesOK() *PostSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostSilencesOK{}
}

type PostSilencesOK struct{ Payload *PostSilencesOKBody }

func (o *PostSilencesOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[POST /silences][%d] postSilencesOK  %+v", 200, o.Payload)
}
func (o *PostSilencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = new(PostSilencesOKBody)
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewPostSilencesBadRequest() *PostSilencesBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostSilencesBadRequest{}
}

type PostSilencesBadRequest struct{ Payload string }

func (o *PostSilencesBadRequest) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[POST /silences][%d] postSilencesBadRequest  %+v", 400, o.Payload)
}
func (o *PostSilencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}

type PostSilencesOKBody struct {
	SilenceID string `json:"silenceID,omitempty"`
}

func (o *PostSilencesOKBody) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (o *PostSilencesOKBody) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}
func (o *PostSilencesOKBody) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res PostSilencesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
