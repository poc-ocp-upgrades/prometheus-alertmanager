package models

import (
	"encoding/json"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type SilenceStatus struct {
	State *string `json:"state"`
}

func (m *SilenceStatus) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var silenceStatusTypeStatePropEnum []interface{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []string
	if err := json.Unmarshal([]byte(`["expired","active","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		silenceStatusTypeStatePropEnum = append(silenceStatusTypeStatePropEnum, v)
	}
}

const (
	SilenceStatusStateExpired	string	= "expired"
	SilenceStatusStateActive	string	= "active"
	SilenceStatusStatePending	string	= "pending"
)

func (m *SilenceStatus) validateStateEnum(path, location string, value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Enum(path, location, value, silenceStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}
func (m *SilenceStatus) validateState(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}
	return nil
}
func (m *SilenceStatus) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *SilenceStatus) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res SilenceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
