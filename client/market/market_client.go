// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new market API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for market API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FindSpotMarketPrices(params *FindSpotMarketPricesParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketPricesOK, error)

	FindSpotMarketPricesHistory(params *FindSpotMarketPricesHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketPricesHistoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FindSpotMarketPrices gets current spot market prices

  Get Equinix Metal current spot market prices.
*/
func (a *Client) FindSpotMarketPrices(params *FindSpotMarketPricesParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSpotMarketPricesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findSpotMarketPrices",
		Method:             "GET",
		PathPattern:        "/market/spot/prices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSpotMarketPricesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSpotMarketPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSpotMarketPrices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindSpotMarketPricesHistory gets spot market prices for a given period of time

  Get spot market prices for a given plan and facility in a fixed period of time

*Note: In the `200` response, the property `datapoints` contains arrays of `[float, integer]`.*
*/
func (a *Client) FindSpotMarketPricesHistory(params *FindSpotMarketPricesHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*FindSpotMarketPricesHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSpotMarketPricesHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findSpotMarketPricesHistory",
		Method:             "GET",
		PathPattern:        "/market/spot/prices/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSpotMarketPricesHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSpotMarketPricesHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSpotMarketPricesHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
