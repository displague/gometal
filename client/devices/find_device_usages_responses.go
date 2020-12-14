// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/models"
)

// FindDeviceUsagesReader is a Reader for the FindDeviceUsages structure.
type FindDeviceUsagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDeviceUsagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDeviceUsagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindDeviceUsagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindDeviceUsagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindDeviceUsagesOK creates a FindDeviceUsagesOK with default headers values
func NewFindDeviceUsagesOK() *FindDeviceUsagesOK {
	return &FindDeviceUsagesOK{}
}

/*FindDeviceUsagesOK handles this case with default header values.

ok
*/
type FindDeviceUsagesOK struct {
	Payload *models.DeviceUsageList
}

func (o *FindDeviceUsagesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/usages][%d] findDeviceUsagesOK  %+v", 200, o.Payload)
}

func (o *FindDeviceUsagesOK) GetPayload() *models.DeviceUsageList {
	return o.Payload
}

func (o *FindDeviceUsagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceUsageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindDeviceUsagesUnauthorized creates a FindDeviceUsagesUnauthorized with default headers values
func NewFindDeviceUsagesUnauthorized() *FindDeviceUsagesUnauthorized {
	return &FindDeviceUsagesUnauthorized{}
}

/*FindDeviceUsagesUnauthorized handles this case with default header values.

unauthorized
*/
type FindDeviceUsagesUnauthorized struct {
}

func (o *FindDeviceUsagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/usages][%d] findDeviceUsagesUnauthorized ", 401)
}

func (o *FindDeviceUsagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDeviceUsagesNotFound creates a FindDeviceUsagesNotFound with default headers values
func NewFindDeviceUsagesNotFound() *FindDeviceUsagesNotFound {
	return &FindDeviceUsagesNotFound{}
}

/*FindDeviceUsagesNotFound handles this case with default header values.

not found
*/
type FindDeviceUsagesNotFound struct {
}

func (o *FindDeviceUsagesNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/usages][%d] findDeviceUsagesNotFound ", 404)
}

func (o *FindDeviceUsagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}