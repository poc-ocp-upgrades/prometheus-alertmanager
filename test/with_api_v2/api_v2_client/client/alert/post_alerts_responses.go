package alert

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

type PostAlertsReader struct{ formats strfmt.Registry }

func (o *PostAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewPostAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewPostAlertsOK() *PostAlertsOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsOK{}
}

type PostAlertsOK struct{}

func (o *PostAlertsOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[POST /alerts][%d] postAlertsOK ", 200)
}
func (o *PostAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func NewPostAlertsBadRequest() *PostAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsBadRequest{}
}

type PostAlertsBadRequest struct{ Payload string }

func (o *PostAlertsBadRequest) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[POST /alerts][%d] postAlertsBadRequest  %+v", 400, o.Payload)
}
func (o *PostAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewPostAlertsInternalServerError() *PostAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PostAlertsInternalServerError{}
}

type PostAlertsInternalServerError struct{ Payload string }

func (o *PostAlertsInternalServerError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[POST /alerts][%d] postAlertsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
