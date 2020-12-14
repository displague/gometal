// Code generated by go-swagger; DO NOT EDIT.

package user_verification_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user verification tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user verification tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ConsumeVerificationRequest(params *ConsumeVerificationRequestParams, authInfo runtime.ClientAuthInfoWriter) (*ConsumeVerificationRequestOK, error)

	CreateValidationRequest(params *CreateValidationRequestParams, authInfo runtime.ClientAuthInfoWriter) (*CreateValidationRequestCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConsumeVerificationRequest verifies a user using an email verification token

  Consumes an email verification token and verifies the user associated with it.
*/
func (a *Client) ConsumeVerificationRequest(params *ConsumeVerificationRequestParams, authInfo runtime.ClientAuthInfoWriter) (*ConsumeVerificationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConsumeVerificationRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "consumeVerificationRequest",
		Method:             "PUT",
		PathPattern:        "/verify-email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConsumeVerificationRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConsumeVerificationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for consumeVerificationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateValidationRequest creates an email verification request

  Creates an email verification request
*/
func (a *Client) CreateValidationRequest(params *CreateValidationRequestParams, authInfo runtime.ClientAuthInfoWriter) (*CreateValidationRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateValidationRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createValidationRequest",
		Method:             "POST",
		PathPattern:        "/verify-email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateValidationRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateValidationRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createValidationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}