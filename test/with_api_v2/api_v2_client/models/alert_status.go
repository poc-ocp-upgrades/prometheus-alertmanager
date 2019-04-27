package models

import (
	"encoding/json"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type AlertStatus struct {
	InhibitedBy	[]string	`json:"inhibitedBy"`
	SilencedBy	[]string	`json:"silencedBy"`
	State		*string		`json:"state"`
}

func (m *AlertStatus) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateInhibitedBy(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateSilencedBy(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *AlertStatus) validateInhibitedBy(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("inhibitedBy", "body", m.InhibitedBy); err != nil {
		return err
	}
	return nil
}
func (m *AlertStatus) validateSilencedBy(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("silencedBy", "body", m.SilencedBy); err != nil {
		return err
	}
	return nil
}

var alertStatusTypeStatePropEnum []interface{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []string
	if err := json.Unmarshal([]byte(`["unprocessed","active","suppressed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertStatusTypeStatePropEnum = append(alertStatusTypeStatePropEnum, v)
	}
}

const (
	AlertStatusStateUnprocessed	string	= "unprocessed"
	AlertStatusStateActive		string	= "active"
	AlertStatusStateSuppressed	string	= "suppressed"
)

func (m *AlertStatus) validateStateEnum(path, location string, value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Enum(path, location, value, alertStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}
func (m *AlertStatus) validateState(formats strfmt.Registry) error {
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
func (m *AlertStatus) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *AlertStatus) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res AlertStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
