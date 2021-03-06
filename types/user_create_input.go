// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserCreateInput user create input
//
// swagger:model UserCreateInput
type UserCreateInput struct {

	// avatar
	// Format: binary
	Avatar io.ReadCloser `json:"avatar,omitempty"`

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// company url
	CompanyURL string `json:"company_url,omitempty"`

	// customdata
	Customdata interface{} `json:"customdata,omitempty"`

	// emails
	// Required: true
	Emails []*EmailInput `json:"emails"`

	// first name
	// Required: true
	FirstName *string `json:"first_name"`

	// last name
	// Required: true
	LastName *string `json:"last_name"`

	// level
	Level string `json:"level,omitempty"`

	// locked
	Locked bool `json:"locked,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// social accounts
	SocialAccounts interface{} `json:"social_accounts,omitempty"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// two factor auth
	TwoFactorAuth string `json:"two_factor_auth,omitempty"`

	// verified at
	// Format: date-time
	VerifiedAt strfmt.DateTime `json:"verified_at,omitempty"`
}

// Validate validates this user create input
func (m *UserCreateInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCreateInput) validateEmails(formats strfmt.Registry) error {

	if err := validate.Required("emails", "body", m.Emails); err != nil {
		return err
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserCreateInput) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *UserCreateInput) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *UserCreateInput) validateVerifiedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.VerifiedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("verified_at", "body", "date-time", m.VerifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCreateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCreateInput) UnmarshalBinary(b []byte) error {
	var res UserCreateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
