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

// UpdateDeviceReader is a Reader for the UpdateDevice structure.
type UpdateDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDeviceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeviceOK creates a UpdateDeviceOK with default headers values
func NewUpdateDeviceOK() *UpdateDeviceOK {
	return &UpdateDeviceOK{}
}

/*UpdateDeviceOK handles this case with default header values.

ok
*/
type UpdateDeviceOK struct {
	Payload *types.Device
}

func (o *UpdateDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}][%d] updateDeviceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceOK) GetPayload() *types.Device {
	return o.Payload
}

func (o *UpdateDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceUnauthorized creates a UpdateDeviceUnauthorized with default headers values
func NewUpdateDeviceUnauthorized() *UpdateDeviceUnauthorized {
	return &UpdateDeviceUnauthorized{}
}

/*UpdateDeviceUnauthorized handles this case with default header values.

unauthorized
*/
type UpdateDeviceUnauthorized struct {
}

func (o *UpdateDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}][%d] updateDeviceUnauthorized ", 401)
}

func (o *UpdateDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeviceForbidden creates a UpdateDeviceForbidden with default headers values
func NewUpdateDeviceForbidden() *UpdateDeviceForbidden {
	return &UpdateDeviceForbidden{}
}

/*UpdateDeviceForbidden handles this case with default header values.

forbidden
*/
type UpdateDeviceForbidden struct {
}

func (o *UpdateDeviceForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}][%d] updateDeviceForbidden ", 403)
}

func (o *UpdateDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeviceNotFound creates a UpdateDeviceNotFound with default headers values
func NewUpdateDeviceNotFound() *UpdateDeviceNotFound {
	return &UpdateDeviceNotFound{}
}

/*UpdateDeviceNotFound handles this case with default header values.

not found
*/
type UpdateDeviceNotFound struct {
}

func (o *UpdateDeviceNotFound) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}][%d] updateDeviceNotFound ", 404)
}

func (o *UpdateDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeviceUnprocessableEntity creates a UpdateDeviceUnprocessableEntity with default headers values
func NewUpdateDeviceUnprocessableEntity() *UpdateDeviceUnprocessableEntity {
	return &UpdateDeviceUnprocessableEntity{}
}

/*UpdateDeviceUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type UpdateDeviceUnprocessableEntity struct {
}

func (o *UpdateDeviceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /devices/{id}][%d] updateDeviceUnprocessableEntity ", 422)
}

func (o *UpdateDeviceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
