package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type Silence struct {
	Comment		*string			`json:"comment"`
	CreatedBy	*string			`json:"createdBy"`
	EndsAt		*strfmt.DateTime	`json:"endsAt"`
	Matchers	Matchers		`json:"matchers"`
	StartsAt	*strfmt.DateTime	`json:"startsAt"`
}

func (m *Silence) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateMatchers(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *Silence) validateComment(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}
	return nil
}
func (m *Silence) validateCreatedBy(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}
	return nil
}
func (m *Silence) validateEndsAt(formats strfmt.Registry) error {
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
func (m *Silence) validateMatchers(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("matchers", "body", m.Matchers); err != nil {
		return err
	}
	if err := m.Matchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matchers")
		}
		return err
	}
	return nil
}
func (m *Silence) validateStartsAt(formats strfmt.Registry) error {
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
func (m *Silence) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *Silence) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res Silence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
