// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BgpConfig bgp config
//
// swagger:model BgpConfig
type BgpConfig struct {

	// Autonomous System Number
	Asn int64 `json:"asn,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// In a Local BGP deployment, a customer uses an internal ASN to control routes within a single Equinix Metal datacenter. This means that the routes are never advertised to the global Internet. Global BGP, on the other hand, requires a customer to have a registered ASN and IP space.
	//
	// Enum: [global local]
	DeploymentType string `json:"deployment_type,omitempty"`

	// href
	Href string `json:"href,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The maximum number of route filters allowed per server
	MaxPrefix int64 `json:"max_prefix,omitempty"`

	// (Optional) Password for BGP session in plaintext (not a checksum)
	Md5 string `json:"md5,omitempty"`

	// project
	Project *Href `json:"project,omitempty"`

	// The IP block ranges associated to the ASN (Populated in Global BGP only)
	Ranges []*GlobalBgpRange `json:"ranges"`

	// requested at
	// Format: date-time
	RequestedAt strfmt.DateTime `json:"requested_at,omitempty"`

	// Specifies AS-MACRO (aka AS-SET) to use when building client route filters
	RouteObject string `json:"route_object,omitempty"`

	// The direct connections between neighboring routers that want to exchange routing information.
	Sessions []*BgpSession `json:"sessions"`

	// status requested is valid only for global deployment
	// Enum: [requested enabled disabled]
	Status string `json:"status,omitempty"`
}

// Validate validates this bgp config
func (m *BgpConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessions(formats); err != nil {
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

func (m *BgpConfig) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var bgpConfigTypeDeploymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["global","local"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bgpConfigTypeDeploymentTypePropEnum = append(bgpConfigTypeDeploymentTypePropEnum, v)
	}
}

const (

	// BgpConfigDeploymentTypeGlobal captures enum value "global"
	BgpConfigDeploymentTypeGlobal string = "global"

	// BgpConfigDeploymentTypeLocal captures enum value "local"
	BgpConfigDeploymentTypeLocal string = "local"
)

// prop value enum
func (m *BgpConfig) validateDeploymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bgpConfigTypeDeploymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BgpConfig) validateDeploymentType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentTypeEnum("deployment_type", "body", m.DeploymentType); err != nil {
		return err
	}

	return nil
}

func (m *BgpConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BgpConfig) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *BgpConfig) validateRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	for i := 0; i < len(m.Ranges); i++ {
		if swag.IsZero(m.Ranges[i]) { // not required
			continue
		}

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BgpConfig) validateRequestedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_at", "body", "date-time", m.RequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BgpConfig) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var bgpConfigTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["requested","enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bgpConfigTypeStatusPropEnum = append(bgpConfigTypeStatusPropEnum, v)
	}
}

const (

	// BgpConfigStatusRequested captures enum value "requested"
	BgpConfigStatusRequested string = "requested"

	// BgpConfigStatusEnabled captures enum value "enabled"
	BgpConfigStatusEnabled string = "enabled"

	// BgpConfigStatusDisabled captures enum value "disabled"
	BgpConfigStatusDisabled string = "disabled"
)

// prop value enum
func (m *BgpConfig) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bgpConfigTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BgpConfig) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BgpConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpConfig) UnmarshalBinary(b []byte) error {
	var res BgpConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}