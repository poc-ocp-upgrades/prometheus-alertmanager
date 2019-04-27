package models

import (
	"encoding/json"
	"strconv"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

type ClusterStatus struct {
	Name	*string		`json:"name"`
	Peers	[]*PeerStatus	`json:"peers"`
	Status	*string		`json:"status"`
}

func (m *ClusterStatus) Validate(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []error
	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}
	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
func (m *ClusterStatus) validateName(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}
	return nil
}
func (m *ClusterStatus) validatePeers(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("peers", "body", m.Peers); err != nil {
		return err
	}
	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) {
			continue
		}
		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}
	return nil
}

var clusterStatusTypeStatusPropEnum []interface{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res []string
	if err := json.Unmarshal([]byte(`["ready","settling","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterStatusTypeStatusPropEnum = append(clusterStatusTypeStatusPropEnum, v)
	}
}

const (
	ClusterStatusStatusReady	string	= "ready"
	ClusterStatusStatusSettling	string	= "settling"
	ClusterStatusStatusDisabled	string	= "disabled"
)

func (m *ClusterStatus) validateStatusEnum(path, location string, value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Enum(path, location, value, clusterStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}
func (m *ClusterStatus) validateStatus(formats strfmt.Registry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}
	return nil
}
func (m *ClusterStatus) MarshalBinary() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}
func (m *ClusterStatus) UnmarshalBinary(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res ClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
