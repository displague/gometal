// Code generated by go-swagger; DO NOT EDIT.

package two_factor_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new two factor auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for two factor auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DisableTfaApp(params *DisableTfaAppParams, authInfo runtime.ClientAuthInfoWriter) (*DisableTfaAppNoContent, error)

	DisableTfaSms(params *DisableTfaSmsParams, authInfo runtime.ClientAuthInfoWriter) (*DisableTfaSmsNoContent, error)

	EnableTfaApp(params *EnableTfaAppParams, authInfo runtime.ClientAuthInfoWriter) (*EnableTfaAppOK, error)

	EnableTfaSms(params *EnableTfaSmsParams, authInfo runtime.ClientAuthInfoWriter) (*EnableTfaSmsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DisableTfaApp disables two factor authentication

  Disables two factor authentication.
*/
func (a *Client) DisableTfaApp(params *DisableTfaAppParams, authInfo runtime.ClientAuthInfoWriter) (*DisableTfaAppNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableTfaAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disableTfaApp",
		Method:             "DELETE",
		PathPattern:        "/user/otp/app",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableTfaAppReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableTfaAppNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableTfaApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisableTfaSms disables two factor authentication

  Disables two factor authentication.
*/
func (a *Client) DisableTfaSms(params *DisableTfaSmsParams, authInfo runtime.ClientAuthInfoWriter) (*DisableTfaSmsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableTfaSmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disableTfaSms",
		Method:             "DELETE",
		PathPattern:        "/user/otp/sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableTfaSmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableTfaSmsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableTfaSms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnableTfaApp enables two factor auth using app

  Enables two factor authentication using authenticator app.
*/
func (a *Client) EnableTfaApp(params *EnableTfaAppParams, authInfo runtime.ClientAuthInfoWriter) (*EnableTfaAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableTfaAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "enableTfaApp",
		Method:             "POST",
		PathPattern:        "/user/otp/app",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableTfaAppReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableTfaAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableTfaApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnableTfaSms enables two factor auth using sms

  Enables two factor authentication with sms.
*/
func (a *Client) EnableTfaSms(params *EnableTfaSmsParams, authInfo runtime.ClientAuthInfoWriter) (*EnableTfaSmsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableTfaSmsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "enableTfaSms",
		Method:             "POST",
		PathPattern:        "/user/otp/sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableTfaSmsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableTfaSmsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableTfaSms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
