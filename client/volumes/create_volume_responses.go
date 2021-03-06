// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// CreateVolumeReader is a Reader for the CreateVolume structure.
type CreateVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVolumeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVolumeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateVolumeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVolumeCreated creates a CreateVolumeCreated with default headers values
func NewCreateVolumeCreated() *CreateVolumeCreated {
	return &CreateVolumeCreated{}
}

/*CreateVolumeCreated handles this case with default header values.

created
*/
type CreateVolumeCreated struct {
	Payload *types.Volume
}

func (o *CreateVolumeCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/storage][%d] createVolumeCreated  %+v", 201, o.Payload)
}

func (o *CreateVolumeCreated) GetPayload() *types.Volume {
	return o.Payload
}

func (o *CreateVolumeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVolumeUnauthorized creates a CreateVolumeUnauthorized with default headers values
func NewCreateVolumeUnauthorized() *CreateVolumeUnauthorized {
	return &CreateVolumeUnauthorized{}
}

/*CreateVolumeUnauthorized handles this case with default header values.

unauthorized
*/
type CreateVolumeUnauthorized struct {
}

func (o *CreateVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/storage][%d] createVolumeUnauthorized ", 401)
}

func (o *CreateVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVolumeForbidden creates a CreateVolumeForbidden with default headers values
func NewCreateVolumeForbidden() *CreateVolumeForbidden {
	return &CreateVolumeForbidden{}
}

/*CreateVolumeForbidden handles this case with default header values.

forbidden
*/
type CreateVolumeForbidden struct {
}

func (o *CreateVolumeForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/storage][%d] createVolumeForbidden ", 403)
}

func (o *CreateVolumeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVolumeNotFound creates a CreateVolumeNotFound with default headers values
func NewCreateVolumeNotFound() *CreateVolumeNotFound {
	return &CreateVolumeNotFound{}
}

/*CreateVolumeNotFound handles this case with default header values.

not found
*/
type CreateVolumeNotFound struct {
}

func (o *CreateVolumeNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/storage][%d] createVolumeNotFound ", 404)
}

func (o *CreateVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVolumeUnprocessableEntity creates a CreateVolumeUnprocessableEntity with default headers values
func NewCreateVolumeUnprocessableEntity() *CreateVolumeUnprocessableEntity {
	return &CreateVolumeUnprocessableEntity{}
}

/*CreateVolumeUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type CreateVolumeUnprocessableEntity struct {
}

func (o *CreateVolumeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/storage][%d] createVolumeUnprocessableEntity ", 422)
}

func (o *CreateVolumeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
