package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type PeerStatus struct {
	Address	*string	`json:"address"`
	Name	*string	`json:"name"`
}

func (m *PeerStatus) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *PeerStatus) validateAddress(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}
	return nil
}
func (m *PeerStatus) validateName(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}
	return nil
}
func (m *PeerStatus) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *PeerStatus) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res PeerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
