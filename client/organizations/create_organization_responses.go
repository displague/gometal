// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateOrganizationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationCreated creates a CreateOrganizationCreated with default headers values
func NewCreateOrganizationCreated() *CreateOrganizationCreated {
	return &CreateOrganizationCreated{}
}

/*CreateOrganizationCreated handles this case with default header values.

created
*/
type CreateOrganizationCreated struct {
	Payload *types.Organization
}

func (o *CreateOrganizationCreated) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationCreated) GetPayload() *types.Organization {
	return o.Payload
}

func (o *CreateOrganizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationUnauthorized creates a CreateOrganizationUnauthorized with default headers values
func NewCreateOrganizationUnauthorized() *CreateOrganizationUnauthorized {
	return &CreateOrganizationUnauthorized{}
}

/*CreateOrganizationUnauthorized handles this case with default header values.

unauthorized
*/
type CreateOrganizationUnauthorized struct {
}

func (o *CreateOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationUnauthorized ", 401)
}

func (o *CreateOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/*CreateOrganizationNotFound handles this case with default header values.

not found
*/
type CreateOrganizationNotFound struct {
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationNotFound ", 404)
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationUnprocessableEntity creates a CreateOrganizationUnprocessableEntity with default headers values
func NewCreateOrganizationUnprocessableEntity() *CreateOrganizationUnprocessableEntity {
	return &CreateOrganizationUnprocessableEntity{}
}

/*CreateOrganizationUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type CreateOrganizationUnprocessableEntity struct {
}

func (o *CreateOrganizationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrganizationUnprocessableEntity ", 422)
}

func (o *CreateOrganizationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
