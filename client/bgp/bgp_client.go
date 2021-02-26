// Code generated by go-swagger; DO NOT EDIT.

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bgp API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bgp API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteBGPSession(params *DeleteBGPSessionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBGPSessionNoContent, error)

	FindBGPSessionByID(params *FindBGPSessionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindBGPSessionByIDOK, error)

	UpdateBGPSession(params *UpdateBGPSessionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBGPSessionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteBGPSession deletes the BGP session

  Deletes the BGP session.
*/
func (a *Client) DeleteBGPSession(params *DeleteBGPSessionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBGPSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBGPSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBgpSession",
		Method:             "DELETE",
		PathPattern:        "/bgp/sessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBGPSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBGPSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBgpSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindBGPSessionByID retrieves a BGP session

  Returns a BGP session
*/
func (a *Client) FindBGPSessionByID(params *FindBGPSessionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindBGPSessionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindBGPSessionByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findBgpSessionById",
		Method:             "GET",
		PathPattern:        "/bgp/sessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindBGPSessionByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindBGPSessionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findBgpSessionById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBGPSession updates the BGP session

  Updates the BGP session by either enabling or disabling the default route functionality.
*/
func (a *Client) UpdateBGPSession(params *UpdateBGPSessionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBGPSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBGPSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateBgpSession",
		Method:             "PUT",
		PathPattern:        "/bgp/sessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBGPSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBGPSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBgpSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
