package general

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type GetStatusReader struct{ formats strfmt.Registry }

func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewGetStatusOK() *GetStatusOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetStatusOK{}
}

type GetStatusOK struct{ Payload *models.AlertmanagerStatus }

func (o *GetStatusOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}
func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o.Payload = new(models.AlertmanagerStatus)
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
