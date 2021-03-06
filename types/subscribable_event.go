// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubscribableEvent subscribable event
//
// swagger:model SubscribableEvent
type SubscribableEvent struct {

	// event name
	EventName string `json:"event_name,omitempty"`

	// event slug
	EventSlug string `json:"event_slug,omitempty"`

	// event type
	EventType string `json:"event_type,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this subscribable event
func (m *SubscribableEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscribableEvent) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscribableEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscribableEvent) UnmarshalBinary(b []byte) error {
	var res SubscribableEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
