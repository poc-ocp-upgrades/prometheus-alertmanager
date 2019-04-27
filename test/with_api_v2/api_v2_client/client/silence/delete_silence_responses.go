package silence

import (
	"fmt"
	"io"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

type DeleteSilenceReader struct{ formats strfmt.Registry }

func (o *DeleteSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch response.Code() {
	case 200:
		result := NewDeleteSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteSilenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}
func NewDeleteSilenceOK() *DeleteSilenceOK {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeleteSilenceOK{}
}

type DeleteSilenceOK struct{}

func (o *DeleteSilenceOK) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceOK ", 200)
}
func (o *DeleteSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func NewDeleteSilenceInternalServerError() *DeleteSilenceInternalServerError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeleteSilenceInternalServerError{}
}

type DeleteSilenceInternalServerError struct{ Payload string }

func (o *DeleteSilenceInternalServerError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("[DELETE /silence/{silenceID}][%d] deleteSilenceInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteSilenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}
	return nil
}
