// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// GetBGPNeighborDataReader is a Reader for the GetBGPNeighborData structure.
type GetBGPNeighborDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBGPNeighborDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBGPNeighborDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBGPNeighborDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBGPNeighborDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBGPNeighborDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBGPNeighborDataOK creates a GetBGPNeighborDataOK with default headers values
func NewGetBGPNeighborDataOK() *GetBGPNeighborDataOK {
	return &GetBGPNeighborDataOK{}
}

/*GetBGPNeighborDataOK handles this case with default header values.

ok
*/
type GetBGPNeighborDataOK struct {
	Payload *types.BGPSessionNeighbors
}

func (o *GetBGPNeighborDataOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bgp/neighbors][%d] getBgpNeighborDataOK  %+v", 200, o.Payload)
}

func (o *GetBGPNeighborDataOK) GetPayload() *types.BGPSessionNeighbors {
	return o.Payload
}

func (o *GetBGPNeighborDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.BGPSessionNeighbors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBGPNeighborDataUnauthorized creates a GetBGPNeighborDataUnauthorized with default headers values
func NewGetBGPNeighborDataUnauthorized() *GetBGPNeighborDataUnauthorized {
	return &GetBGPNeighborDataUnauthorized{}
}

/*GetBGPNeighborDataUnauthorized handles this case with default header values.

unauthorized
*/
type GetBGPNeighborDataUnauthorized struct {
}

func (o *GetBGPNeighborDataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bgp/neighbors][%d] getBgpNeighborDataUnauthorized ", 401)
}

func (o *GetBGPNeighborDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBGPNeighborDataForbidden creates a GetBGPNeighborDataForbidden with default headers values
func NewGetBGPNeighborDataForbidden() *GetBGPNeighborDataForbidden {
	return &GetBGPNeighborDataForbidden{}
}

/*GetBGPNeighborDataForbidden handles this case with default header values.

forbidden
*/
type GetBGPNeighborDataForbidden struct {
}

func (o *GetBGPNeighborDataForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bgp/neighbors][%d] getBgpNeighborDataForbidden ", 403)
}

func (o *GetBGPNeighborDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBGPNeighborDataNotFound creates a GetBGPNeighborDataNotFound with default headers values
func NewGetBGPNeighborDataNotFound() *GetBGPNeighborDataNotFound {
	return &GetBGPNeighborDataNotFound{}
}

/*GetBGPNeighborDataNotFound handles this case with default header values.

not found
*/
type GetBGPNeighborDataNotFound struct {
}

func (o *GetBGPNeighborDataNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bgp/neighbors][%d] getBgpNeighborDataNotFound ", 404)
}

func (o *GetBGPNeighborDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
