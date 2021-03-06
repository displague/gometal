// Code generated by go-swagger; DO NOT EDIT.

package ip_addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ip addresses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ip addresses API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteIPAddress(params *DeleteIPAddressParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPAddressNoContent, error)

	FindIPAddressByID(params *FindIPAddressByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAddressByIDOK, error)

	FindIPAddressCustomdata(params *FindIPAddressCustomdataParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAddressCustomdataOK, error)

	FindIPAvailabilities(params *FindIPAvailabilitiesParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAvailabilitiesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteIPAddress unassigns an ip address

  Note! This call can be used to un-assign an IP assignment or delete an IP reservation. Un-assign an IP address record. Use the assignment UUID you get after attaching the IP. This will remove the relationship between an IP and the device and will make the IP address available to be assigned to another device. Delete and IP reservation. Use the reservation UUID you get after adding the IP to the project. This will permanently delete the IP block reservation from the project.
*/
func (a *Client) DeleteIPAddress(params *DeleteIPAddressParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteIPAddressNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteIPAddress",
		Method:             "DELETE",
		PathPattern:        "/ips/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteIPAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIPAddressNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteIPAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindIPAddressByID retrieves an ip address

  Returns a single ip address if the user has access.
*/
func (a *Client) FindIPAddressByID(params *FindIPAddressByIDParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAddressByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPAddressByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPAddressById",
		Method:             "GET",
		PathPattern:        "/ips/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPAddressByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindIPAddressByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findIPAddressById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindIPAddressCustomdata retrieves the custom metadata of an IP reservation or IP assignment

  Provides the custom metadata stored for this IP Reservation or IP Assignment in json format
*/
func (a *Client) FindIPAddressCustomdata(params *FindIPAddressCustomdataParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAddressCustomdataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPAddressCustomdataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPAddressCustomdata",
		Method:             "GET",
		PathPattern:        "/ips/{id}/customdata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPAddressCustomdataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindIPAddressCustomdataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findIPAddressCustomdata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindIPAvailabilities retrieves all available subnets of a particular reservation

  Provides a list of IP resevations for a single project.
*/
func (a *Client) FindIPAvailabilities(params *FindIPAvailabilitiesParams, authInfo runtime.ClientAuthInfoWriter) (*FindIPAvailabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindIPAvailabilitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findIPAvailabilities",
		Method:             "GET",
		PathPattern:        "/ips/{id}/available",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindIPAvailabilitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindIPAvailabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findIPAvailabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
