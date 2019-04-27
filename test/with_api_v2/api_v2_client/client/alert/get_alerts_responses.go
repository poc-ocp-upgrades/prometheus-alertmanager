package alert

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type GetAlertsReader struct{ formats strfmt.Registry }

func (o *GetAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewGetAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewGetAlertsOK() *GetAlertsOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsOK{}
}

type GetAlertsOK struct{ Payload models.GettableAlerts }

func (o *GetAlertsOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /alerts][%d] getAlertsOK  %+v", 200, o.Payload)
}
func (o *GetAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewGetAlertsBadRequest() *GetAlertsBadRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsBadRequest{}
}

type GetAlertsBadRequest struct{ Payload string }

func (o *GetAlertsBadRequest) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /alerts][%d] getAlertsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
func NewGetAlertsInternalServerError() *GetAlertsInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetAlertsInternalServerError{}
}

type GetAlertsInternalServerError struct{ Payload string }

func (o *GetAlertsInternalServerError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /alerts][%d] getAlertsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
