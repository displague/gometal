// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HardwareLocation hardware location
//
// swagger:model HardwareLocation
type HardwareLocation struct {

	// cage
	Cage string `json:"cage,omitempty"`

	// facility
	Facility string `json:"facility,omitempty"`

	// rack
	Rack string `json:"rack,omitempty"`

	// row
	Row string `json:"row,omitempty"`

	// switch
	Switch string `json:"switch,omitempty"`
}

// Validate validates this hardware location
func (m *HardwareLocation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HardwareLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwareLocation) UnmarshalBinary(b []byte) error {
	var res HardwareLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
