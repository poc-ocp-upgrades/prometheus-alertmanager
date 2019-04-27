package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type VersionInfo struct {
	Branch		*string	`json:"branch"`
	BuildDate	*string	`json:"buildDate"`
	BuildUser	*string	`json:"buildUser"`
	GoVersion	*string	`json:"goVersion"`
	Revision	*string	`json:"revision"`
	Version		*string	`json:"version"`
}

func (m *VersionInfo) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateBuildDate(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateBuildUser(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateGoVersion(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *VersionInfo) validateBranch(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("branch", "body", m.Branch); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) validateBuildDate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("buildDate", "body", m.BuildDate); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) validateBuildUser(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("buildUser", "body", m.BuildUser); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) validateGoVersion(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("goVersion", "body", m.GoVersion); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) validateRevision(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) validateVersion(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}
	return nil
}
func (m *VersionInfo) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *VersionInfo) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res VersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
