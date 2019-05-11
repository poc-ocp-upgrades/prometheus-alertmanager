package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type GettableSilence struct {
	ID			*string				`json:"id"`
	Status		*SilenceStatus		`json:"status"`
	UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	Silence
}

func (m *GettableSilence) UnmarshalJSON(raw []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dataAO0 struct {
		ID			*string				`json:"id"`
		Status		*SilenceStatus		`json:"status"`
		UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}
	m.ID = dataAO0.ID
	m.Status = dataAO0.Status
	m.UpdatedAt = dataAO0.UpdatedAt
	var aO1 Silence
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Silence = aO1
	return nil
}
func (m GettableSilence) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_parts := make([][]byte, 0, 2)
	var dataAO0 struct {
		ID			*string				`json:"id"`
		Status		*SilenceStatus		`json:"status"`
		UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	}
	dataAO0.ID = m.ID
	dataAO0.Status = m.Status
	dataAO0.UpdatedAt = m.UpdatedAt
	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	aO1, err := swag.WriteJSON(m.Silence)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}
func (m *GettableSilence) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.Silence.Validate(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *GettableSilence) validateID(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}
	return nil
}
func (m *GettableSilence) validateStatus(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}
	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}
	return nil
}
func (m *GettableSilence) validateUpdatedAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}
	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *GettableSilence) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *GettableSilence) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res GettableSilence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
