// Code generated by go-swagger; DO NOT EDIT.

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vpn API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vpn API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FindCurrentUserVPNConfig(params *FindCurrentUserVPNConfigParams, authInfo runtime.ClientAuthInfoWriter) (*FindCurrentUserVPNConfigOK, error)

	TurnOffCurrentUserVPN(params *TurnOffCurrentUserVPNParams, authInfo runtime.ClientAuthInfoWriter) (*TurnOffCurrentUserVPNNoContent, error)

	TurnOnCurrentUserVPN(params *TurnOnCurrentUserVPNParams, authInfo runtime.ClientAuthInfoWriter) (*TurnOnCurrentUserVPNCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FindCurrentUserVPNConfig retrieves the client vpn config for current user

  Returns the client vpn config for the currently logged-in user.
*/
func (a *Client) FindCurrentUserVPNConfig(params *FindCurrentUserVPNConfigParams, authInfo runtime.ClientAuthInfoWriter) (*FindCurrentUserVPNConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindCurrentUserVPNConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findCurrentUserVpnConfig",
		Method:             "GET",
		PathPattern:        "/user/vpn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindCurrentUserVPNConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindCurrentUserVPNConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findCurrentUserVpnConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TurnOffCurrentUserVPN turns off vpn for the current user

  Turns off vpn for the currently logged-in user.
*/
func (a *Client) TurnOffCurrentUserVPN(params *TurnOffCurrentUserVPNParams, authInfo runtime.ClientAuthInfoWriter) (*TurnOffCurrentUserVPNNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTurnOffCurrentUserVPNParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "turnOffCurrentUserVpn",
		Method:             "DELETE",
		PathPattern:        "/user/vpn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TurnOffCurrentUserVPNReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TurnOffCurrentUserVPNNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for turnOffCurrentUserVpn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TurnOnCurrentUserVPN turns on vpn for the current user

  Turns on vpn for the currently logged-in user.
*/
func (a *Client) TurnOnCurrentUserVPN(params *TurnOnCurrentUserVPNParams, authInfo runtime.ClientAuthInfoWriter) (*TurnOnCurrentUserVPNCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTurnOnCurrentUserVPNParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "turnOnCurrentUserVpn",
		Method:             "POST",
		PathPattern:        "/user/vpn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TurnOnCurrentUserVPNReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TurnOnCurrentUserVPNCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for turnOnCurrentUserVpn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
