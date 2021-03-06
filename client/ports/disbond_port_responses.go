// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// DisbondPortReader is a Reader for the DisbondPort structure.
type DisbondPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisbondPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisbondPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDisbondPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisbondPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisbondPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDisbondPortUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDisbondPortOK creates a DisbondPortOK with default headers values
func NewDisbondPortOK() *DisbondPortOK {
	return &DisbondPortOK{}
}

/*DisbondPortOK handles this case with default header values.

ok
*/
type DisbondPortOK struct {
	Payload *types.Port
}

func (o *DisbondPortOK) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/disbond][%d] disbondPortOK  %+v", 200, o.Payload)
}

func (o *DisbondPortOK) GetPayload() *types.Port {
	return o.Payload
}

func (o *DisbondPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Port)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisbondPortUnauthorized creates a DisbondPortUnauthorized with default headers values
func NewDisbondPortUnauthorized() *DisbondPortUnauthorized {
	return &DisbondPortUnauthorized{}
}

/*DisbondPortUnauthorized handles this case with default header values.

unauthorized
*/
type DisbondPortUnauthorized struct {
}

func (o *DisbondPortUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/disbond][%d] disbondPortUnauthorized ", 401)
}

func (o *DisbondPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisbondPortForbidden creates a DisbondPortForbidden with default headers values
func NewDisbondPortForbidden() *DisbondPortForbidden {
	return &DisbondPortForbidden{}
}

/*DisbondPortForbidden handles this case with default header values.

forbidden
*/
type DisbondPortForbidden struct {
}

func (o *DisbondPortForbidden) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/disbond][%d] disbondPortForbidden ", 403)
}

func (o *DisbondPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisbondPortNotFound creates a DisbondPortNotFound with default headers values
func NewDisbondPortNotFound() *DisbondPortNotFound {
	return &DisbondPortNotFound{}
}

/*DisbondPortNotFound handles this case with default header values.

not found
*/
type DisbondPortNotFound struct {
}

func (o *DisbondPortNotFound) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/disbond][%d] disbondPortNotFound ", 404)
}

func (o *DisbondPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisbondPortUnprocessableEntity creates a DisbondPortUnprocessableEntity with default headers values
func NewDisbondPortUnprocessableEntity() *DisbondPortUnprocessableEntity {
	return &DisbondPortUnprocessableEntity{}
}

/*DisbondPortUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type DisbondPortUnprocessableEntity struct {
}

func (o *DisbondPortUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/disbond][%d] disbondPortUnprocessableEntity ", 422)
}

func (o *DisbondPortUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
