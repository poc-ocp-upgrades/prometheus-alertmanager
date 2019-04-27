package silence

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type GetSilencesReader struct{ formats strfmt.Registry }

func (o *GetSilencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewGetSilencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSilencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewGetSilencesOK() *GetSilencesOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilencesOK{}
}

type GetSilencesOK struct{ Payload models.GettableSilences }

func (o *GetSilencesOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /silences][%d] getSilencesOK  %+v", 200, o.Payload)
}
func (o *GetSilencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewGetSilencesInternalServerError() *GetSilencesInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetSilencesInternalServerError{}
}

type GetSilencesInternalServerError struct{ Payload string }

func (o *GetSilencesInternalServerError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /silences][%d] getSilencesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSilencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
