package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type AlertmanagerStatus struct {
	Cluster		*ClusterStatus		`json:"cluster"`
	Config		*AlertmanagerConfig	`json:"config"`
	Uptime		*strfmt.DateTime	`json:"uptime"`
	VersionInfo	*VersionInfo		`json:"versionInfo"`
}

func (m *AlertmanagerStatus) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateUptime(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateVersionInfo(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *AlertmanagerStatus) validateCluster(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}
	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}
	return nil
}
func (m *AlertmanagerStatus) validateConfig(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}
	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}
	return nil
}
func (m *AlertmanagerStatus) validateUptime(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("uptime", "body", m.Uptime); err != nil {
		return err
	}
	if err := validate.FormatOf("uptime", "body", "date-time", m.Uptime.String(), formats); err != nil {
		return err
	}
	return nil
}
func (m *AlertmanagerStatus) validateVersionInfo(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("versionInfo", "body", m.VersionInfo); err != nil {
		return err
	}
	if m.VersionInfo != nil {
		if err := m.VersionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionInfo")
			}
			return err
		}
	}
	return nil
}
func (m *AlertmanagerStatus) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *AlertmanagerStatus) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res AlertmanagerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
