// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AssignNativeVlan(params *AssignNativeVlanParams, authInfo runtime.ClientAuthInfoWriter) (*AssignNativeVlanOK, error)

	AssignPort(params *AssignPortParams, authInfo runtime.ClientAuthInfoWriter) (*AssignPortOK, error)

	BondPort(params *BondPortParams, authInfo runtime.ClientAuthInfoWriter) (*BondPortOK, error)

	ConvertLayer2(params *ConvertLayer2Params, authInfo runtime.ClientAuthInfoWriter) (*ConvertLayer2OK, error)

	ConvertLayer3(params *ConvertLayer3Params, authInfo runtime.ClientAuthInfoWriter) (*ConvertLayer3OK, error)

	DeleteNativeVlan(params *DeleteNativeVlanParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNativeVlanOK, error)

	DisbondPort(params *DisbondPortParams, authInfo runtime.ClientAuthInfoWriter) (*DisbondPortOK, error)

	UnassignPort(params *UnassignPortParams, authInfo runtime.ClientAuthInfoWriter) (*UnassignPortOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AssignNativeVlan assigns a native v l a n

  Assigns a virtual network to this port as a "native VLAN"
*/
func (a *Client) AssignNativeVlan(params *AssignNativeVlanParams, authInfo runtime.ClientAuthInfoWriter) (*AssignNativeVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignNativeVlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "assignNativeVlan",
		Method:             "POST",
		PathPattern:        "/ports/{id}/native-vlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AssignNativeVlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssignNativeVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignNativeVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssignPort assigns a port to virtual network

  Assign a port for a hardware to virtual network.
*/
func (a *Client) AssignPort(params *AssignPortParams, authInfo runtime.ClientAuthInfoWriter) (*AssignPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "assignPort",
		Method:             "POST",
		PathPattern:        "/ports/{id}/assign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AssignPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssignPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BondPort enablings bonding

  Enabling bonding for one or all ports
*/
func (a *Client) BondPort(params *BondPortParams, authInfo runtime.ClientAuthInfoWriter) (*BondPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBondPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bondPort",
		Method:             "POST",
		PathPattern:        "/ports/{id}/bond",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BondPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BondPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bondPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConvertLayer2 converts to layer 2

  Converts a bond port to Layer 2. IP assignments of the port will be removed.
*/
func (a *Client) ConvertLayer2(params *ConvertLayer2Params, authInfo runtime.ClientAuthInfoWriter) (*ConvertLayer2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConvertLayer2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "convertLayer2",
		Method:             "POST",
		PathPattern:        "/ports/{id}/convert/layer-2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConvertLayer2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConvertLayer2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for convertLayer2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConvertLayer3 converts to layer 3

  Converts a bond port to Layer 3. VLANs must first be unassigned.
*/
func (a *Client) ConvertLayer3(params *ConvertLayer3Params, authInfo runtime.ClientAuthInfoWriter) (*ConvertLayer3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConvertLayer3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "convertLayer3",
		Method:             "POST",
		PathPattern:        "/ports/{id}/convert/layer-3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConvertLayer3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConvertLayer3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for convertLayer3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNativeVlan removes native v l a n

  Removes the native VLAN from this port
*/
func (a *Client) DeleteNativeVlan(params *DeleteNativeVlanParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNativeVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNativeVlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNativeVlan",
		Method:             "DELETE",
		PathPattern:        "/ports/{id}/native-vlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNativeVlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNativeVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNativeVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisbondPort disablings bonding

  Disabling bonding for one or all ports
*/
func (a *Client) DisbondPort(params *DisbondPortParams, authInfo runtime.ClientAuthInfoWriter) (*DisbondPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisbondPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disbondPort",
		Method:             "POST",
		PathPattern:        "/ports/{id}/disbond",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisbondPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisbondPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disbondPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnassignPort unassigns a port

  Unassign a port for a hardware.
*/
func (a *Client) UnassignPort(params *UnassignPortParams, authInfo runtime.ClientAuthInfoWriter) (*UnassignPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnassignPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unassignPort",
		Method:             "POST",
		PathPattern:        "/ports/{id}/unassign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnassignPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnassignPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unassignPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}