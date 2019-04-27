package models

import (
	"strconv"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type GettableAlert struct {
	Annotations	LabelSet		`json:"annotations"`
	EndsAt		*strfmt.DateTime	`json:"endsAt"`
	Fingerprint	*string			`json:"fingerprint"`
	Receivers	[]*Receiver		`json:"receivers"`
	StartsAt	*strfmt.DateTime	`json:"startsAt"`
	Status		*AlertStatus		`json:"status"`
	UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	Alert
}

func (m *GettableAlert) UnmarshalJSON(raw []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dataAO0 struct {
		Annotations	LabelSet		`json:"annotations"`
		EndsAt		*strfmt.DateTime	`json:"endsAt"`
		Fingerprint	*string			`json:"fingerprint"`
		Receivers	[]*Receiver		`json:"receivers"`
		StartsAt	*strfmt.DateTime	`json:"startsAt"`
		Status		*AlertStatus		`json:"status"`
		UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}
	m.Annotations = dataAO0.Annotations
	m.EndsAt = dataAO0.EndsAt
	m.Fingerprint = dataAO0.Fingerprint
	m.Receivers = dataAO0.Receivers
	m.StartsAt = dataAO0.StartsAt
	m.Status = dataAO0.Status
	m.UpdatedAt = dataAO0.UpdatedAt
	var aO1 Alert
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Alert = aO1
	return nil
}
func (m GettableAlert) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_parts := make([][]byte, 0, 2)
	var dataAO0 struct {
		Annotations	LabelSet		`json:"annotations"`
		EndsAt		*strfmt.DateTime	`json:"endsAt"`
		Fingerprint	*string			`json:"fingerprint"`
		Receivers	[]*Receiver		`json:"receivers"`
		StartsAt	*strfmt.DateTime	`json:"startsAt"`
		Status		*AlertStatus		`json:"status"`
		UpdatedAt	*strfmt.DateTime	`json:"updatedAt"`
	}
	dataAO0.Annotations = m.Annotations
	dataAO0.EndsAt = m.EndsAt
	dataAO0.Fingerprint = m.Fingerprint
	dataAO0.Receivers = m.Receivers
	dataAO0.StartsAt = m.StartsAt
	dataAO0.Status = m.Status
	dataAO0.UpdatedAt = m.UpdatedAt
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
func (m *GettableAlert) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateUpdatedAt(formats); err != nil {
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
func (m *GettableAlert) validateAnnotations(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}
	return nil
}
func (m *GettableAlert) validateEndsAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("endsAt", "body", m.EndsAt); err != nil {
		return err
	}
	if err := validate.FormatOf("endsAt", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *GettableAlert) validateFingerprint(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}
	return nil
}
func (m *GettableAlert) validateReceivers(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("receivers", "body", m.Receivers); err != nil {
		return err
	}
	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) {
			continue
		}
		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}
	return nil
}
func (m *GettableAlert) validateStartsAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("startsAt", "body", m.StartsAt); err != nil {
		return err
	}
	if err := validate.FormatOf("startsAt", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *GettableAlert) validateStatus(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (m *GettableAlert) validateUpdatedAt(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (m *GettableAlert) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *GettableAlert) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res GettableAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
