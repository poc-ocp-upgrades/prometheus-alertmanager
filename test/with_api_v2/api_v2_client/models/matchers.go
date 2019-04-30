package models

import (
	"strconv"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type Matchers []*Matcher

func (m Matchers) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	iMatchersSize := int64(len(m))
	if err := validate.MinItems("", "body", iMatchersSize, 1); err != nil {
		return err
	}
	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) {
			continue
		}
		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
