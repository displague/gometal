// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindDeviceCustomdataReader is a Reader for the FindDeviceCustomdata structure.
type FindDeviceCustomdataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDeviceCustomdataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDeviceCustomdataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindDeviceCustomdataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindDeviceCustomdataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindDeviceCustomdataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindDeviceCustomdataOK creates a FindDeviceCustomdataOK with default headers values
func NewFindDeviceCustomdataOK() *FindDeviceCustomdataOK {
	return &FindDeviceCustomdataOK{}
}

/*FindDeviceCustomdataOK handles this case with default header values.

ok
*/
type FindDeviceCustomdataOK struct {
}

func (o *FindDeviceCustomdataOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/customdata][%d] findDeviceCustomdataOK ", 200)
}

func (o *FindDeviceCustomdataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDeviceCustomdataUnauthorized creates a FindDeviceCustomdataUnauthorized with default headers values
func NewFindDeviceCustomdataUnauthorized() *FindDeviceCustomdataUnauthorized {
	return &FindDeviceCustomdataUnauthorized{}
}

/*FindDeviceCustomdataUnauthorized handles this case with default header values.

unauthorized
*/
type FindDeviceCustomdataUnauthorized struct {
}

func (o *FindDeviceCustomdataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/customdata][%d] findDeviceCustomdataUnauthorized ", 401)
}

func (o *FindDeviceCustomdataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDeviceCustomdataForbidden creates a FindDeviceCustomdataForbidden with default headers values
func NewFindDeviceCustomdataForbidden() *FindDeviceCustomdataForbidden {
	return &FindDeviceCustomdataForbidden{}
}

/*FindDeviceCustomdataForbidden handles this case with default header values.

forbidden
*/
type FindDeviceCustomdataForbidden struct {
}

func (o *FindDeviceCustomdataForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/customdata][%d] findDeviceCustomdataForbidden ", 403)
}

func (o *FindDeviceCustomdataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDeviceCustomdataNotFound creates a FindDeviceCustomdataNotFound with default headers values
func NewFindDeviceCustomdataNotFound() *FindDeviceCustomdataNotFound {
	return &FindDeviceCustomdataNotFound{}
}

/*FindDeviceCustomdataNotFound handles this case with default header values.

not found
*/
type FindDeviceCustomdataNotFound struct {
}

func (o *FindDeviceCustomdataNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/customdata][%d] findDeviceCustomdataNotFound ", 404)
}

func (o *FindDeviceCustomdataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}