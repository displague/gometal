// Code generated by go-swagger; DO NOT EDIT.

package spot_market_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new spot market request API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for spot market request API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSpotMarketRequest(params *DeleteSpotMarketRequestParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSpotMarketRequestNoContent, error)

	FindSpotMarketRequestByID(params *FindSpotMarketRequestByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketRequestByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteSpotMarketRequest deletes the spot market request

  Deletes the spot market request.
*/
func (a *Client) DeleteSpotMarketRequest(params *DeleteSpotMarketRequestParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSpotMarketRequestNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpotMarketRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSpotMarketRequest",
		Method:             "DELETE",
		PathPattern:        "/spot-market-requests/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSpotMarketRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSpotMarketRequestNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSpotMarketRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindSpotMarketRequestByID retrieves a spot market request

  Returns a single spot market request
*/
func (a *Client) FindSpotMarketRequestByID(params *FindSpotMarketRequestByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketRequestByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSpotMarketRequestByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findSpotMarketRequestById",
		Method:             "GET",
		PathPattern:        "/spot-market-requests/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSpotMarketRequestByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSpotMarketRequestByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSpotMarketRequestById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
