package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

type PostableSilence struct {
	ID	string	`json:"id,omitempty"`
	Silence
}

func (m *PostableSilence) UnmarshalJSON(raw []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dataAO0 struct {
		ID string `json:"id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}
	m.ID = dataAO0.ID
	var aO1 Silence
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Silence = aO1
	return nil
}
func (m PostableSilence) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_parts := make([][]byte, 0, 2)
	var dataAO0 struct {
		ID string `json:"id,omitempty"`
	}
	dataAO0.ID = m.ID
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
func (m *PostableSilence) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.Silence.Validate(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *PostableSilence) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *PostableSilence) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res PostableSilence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
