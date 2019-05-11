package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type PostableAlert struct {
	Annotations	LabelSet		`json:"annotations,omitempty"`
	EndsAt		strfmt.DateTime	`json:"endsAt,omitempty"`
	StartsAt	strfmt.DateTime	`json:"startsAt,omitempty"`
	Alert
}

func (m *PostableAlert) UnmarshalJSON(raw []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dataAO0 struct {
		Annotations	LabelSet		`json:"annotations,omitempty"`
		EndsAt		strfmt.DateTime	`json:"endsAt,omitempty"`
		StartsAt	strfmt.DateTime	`json:"startsAt,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}
	m.Annotations = dataAO0.Annotations
	m.EndsAt = dataAO0.EndsAt
	m.StartsAt = dataAO0.StartsAt
	var aO1 Alert
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Alert = aO1
	return nil
}
func (m PostableAlert) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_parts := make([][]byte, 0, 2)
	var dataAO0 struct {
		Annotations	LabelSet		`json:"annotations,omitempty"`
		EndsAt		strfmt.DateTime	`json:"endsAt,omitempty"`
		StartsAt	strfmt.DateTime	`json:"startsAt,omitempty"`
	}
	dataAO0.Annotations = m.Annotations
	dataAO0.EndsAt = m.EndsAt
	dataAO0.StartsAt = m.StartsAt
	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	aO1, err := swag.WriteJSON(m.Alert)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}
func (m *PostableAlert) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.Alert.Validate(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *PostableAlert) validateAnnotations(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if swag.IsZero(m.Annotations) {
		return nil
	}
	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}
	return nil
}
func (m *PostableAlert) validateEndsAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if swag.IsZero(m.EndsAt) {
		return nil
	}
	if err := validate.FormatOf("endsAt", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *PostableAlert) validateStartsAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if swag.IsZero(m.StartsAt) {
		return nil
	}
	if err := validate.FormatOf("startsAt", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *PostableAlert) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *PostableAlert) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res PostableAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
