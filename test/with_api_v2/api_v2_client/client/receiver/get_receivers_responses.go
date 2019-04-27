package receiver

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type GetReceiversReader struct{ formats strfmt.Registry }

func (o *GetReceiversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewGetReceiversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewGetReceiversOK() *GetReceiversOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &GetReceiversOK{}
}

type GetReceiversOK struct{ Payload []*models.Receiver }

func (o *GetReceiversOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[GET /receivers][%d] getReceiversOK  %+v", 200, o.Payload)
}
func (o *GetReceiversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
