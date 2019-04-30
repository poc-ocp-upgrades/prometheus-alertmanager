package silence

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type GetSilenceReader struct{ formats strfmt.Registry }

func (o *GetSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewGetSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSilenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSilenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewGetSilenceOK() *GetSilenceOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceOK{}
}

type GetSilenceOK struct{ Payload *models.GettableSilence }

func (o *GetSilenceOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceOK  %+v", 200, o.Payload)
}
func (o *GetSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = new(models.GettableSilence)
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewGetSilenceNotFound() *GetSilenceNotFound {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceNotFound{}
}

type GetSilenceNotFound struct{}

func (o *GetSilenceNotFound) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceNotFound ", 404)
}
func (o *GetSilenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func NewGetSilenceInternalServerError() *GetSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilenceInternalServerError{}
}

type GetSilenceInternalServerError struct{ Payload string }

func (o *GetSilenceInternalServerError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSilenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
