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

// CreateIPAssignmentReader is a Reader for the CreateIPAssignment structure.
type CreateIPAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIPAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIPAssignmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateIPAssignmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateIPAssignmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateIPAssignmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateIPAssignmentCreated creates a CreateIPAssignmentCreated with default headers values
func NewCreateIPAssignmentCreated() *CreateIPAssignmentCreated {
	return &CreateIPAssignmentCreated{}
}

/*CreateIPAssignmentCreated handles this case with default header values.

created
*/
type CreateIPAssignmentCreated struct {
	Payload *types.IPAssignment
}

func (o *CreateIPAssignmentCreated) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/ips][%d] createIpAssignmentCreated  %+v", 201, o.Payload)
}

func (o *CreateIPAssignmentCreated) GetPayload() *types.IPAssignment {
	return o.Payload
}

func (o *CreateIPAssignmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.IPAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIPAssignmentUnauthorized creates a CreateIPAssignmentUnauthorized with default headers values
func NewCreateIPAssignmentUnauthorized() *CreateIPAssignmentUnauthorized {
	return &CreateIPAssignmentUnauthorized{}
}

/*CreateIPAssignmentUnauthorized handles this case with default header values.

unauthorized
*/
type CreateIPAssignmentUnauthorized struct {
}

func (o *CreateIPAssignmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/ips][%d] createIpAssignmentUnauthorized ", 401)
}

func (o *CreateIPAssignmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateIPAssignmentNotFound creates a CreateIPAssignmentNotFound with default headers values
func NewCreateIPAssignmentNotFound() *CreateIPAssignmentNotFound {
	return &CreateIPAssignmentNotFound{}
}

/*CreateIPAssignmentNotFound handles this case with default header values.

not found
*/
type CreateIPAssignmentNotFound struct {
}

func (o *CreateIPAssignmentNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/ips][%d] createIpAssignmentNotFound ", 404)
}

func (o *CreateIPAssignmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateIPAssignmentUnprocessableEntity creates a CreateIPAssignmentUnprocessableEntity with default headers values
func NewCreateIPAssignmentUnprocessableEntity() *CreateIPAssignmentUnprocessableEntity {
	return &CreateIPAssignmentUnprocessableEntity{}
}

/*CreateIPAssignmentUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type CreateIPAssignmentUnprocessableEntity struct {
}

func (o *CreateIPAssignmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/ips][%d] createIpAssignmentUnprocessableEntity ", 422)
}

func (o *CreateIPAssignmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
