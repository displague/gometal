// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindInstanceBandwidthReader is a Reader for the FindInstanceBandwidth structure.
type FindInstanceBandwidthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindInstanceBandwidthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindInstanceBandwidthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewFindInstanceBandwidthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindInstanceBandwidthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindInstanceBandwidthOK creates a FindInstanceBandwidthOK with default headers values
func NewFindInstanceBandwidthOK() *FindInstanceBandwidthOK {
	return &FindInstanceBandwidthOK{}
}

/*FindInstanceBandwidthOK handles this case with default header values.

ok
*/
type FindInstanceBandwidthOK struct {
}

func (o *FindInstanceBandwidthOK) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bandwidth][%d] findInstanceBandwidthOK ", 200)
}

func (o *FindInstanceBandwidthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindInstanceBandwidthForbidden creates a FindInstanceBandwidthForbidden with default headers values
func NewFindInstanceBandwidthForbidden() *FindInstanceBandwidthForbidden {
	return &FindInstanceBandwidthForbidden{}
}

/*FindInstanceBandwidthForbidden handles this case with default header values.

forbidden
*/
type FindInstanceBandwidthForbidden struct {
}

func (o *FindInstanceBandwidthForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bandwidth][%d] findInstanceBandwidthForbidden ", 403)
}

func (o *FindInstanceBandwidthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindInstanceBandwidthNotFound creates a FindInstanceBandwidthNotFound with default headers values
func NewFindInstanceBandwidthNotFound() *FindInstanceBandwidthNotFound {
	return &FindInstanceBandwidthNotFound{}
}

/*FindInstanceBandwidthNotFound handles this case with default header values.

not found
*/
type FindInstanceBandwidthNotFound struct {
}

func (o *FindInstanceBandwidthNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/{id}/bandwidth][%d] findInstanceBandwidthNotFound ", 404)
}

func (o *FindInstanceBandwidthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
