// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindOrganizationCustomdataReader is a Reader for the FindOrganizationCustomdata structure.
type FindOrganizationCustomdataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOrganizationCustomdataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOrganizationCustomdataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindOrganizationCustomdataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindOrganizationCustomdataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindOrganizationCustomdataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindOrganizationCustomdataOK creates a FindOrganizationCustomdataOK with default headers values
func NewFindOrganizationCustomdataOK() *FindOrganizationCustomdataOK {
	return &FindOrganizationCustomdataOK{}
}

/*FindOrganizationCustomdataOK handles this case with default header values.

ok
*/
type FindOrganizationCustomdataOK struct {
}

func (o *FindOrganizationCustomdataOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/customdata][%d] findOrganizationCustomdataOK ", 200)
}

func (o *FindOrganizationCustomdataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindOrganizationCustomdataUnauthorized creates a FindOrganizationCustomdataUnauthorized with default headers values
func NewFindOrganizationCustomdataUnauthorized() *FindOrganizationCustomdataUnauthorized {
	return &FindOrganizationCustomdataUnauthorized{}
}

/*FindOrganizationCustomdataUnauthorized handles this case with default header values.

unauthorized
*/
type FindOrganizationCustomdataUnauthorized struct {
}

func (o *FindOrganizationCustomdataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/customdata][%d] findOrganizationCustomdataUnauthorized ", 401)
}

func (o *FindOrganizationCustomdataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindOrganizationCustomdataForbidden creates a FindOrganizationCustomdataForbidden with default headers values
func NewFindOrganizationCustomdataForbidden() *FindOrganizationCustomdataForbidden {
	return &FindOrganizationCustomdataForbidden{}
}

/*FindOrganizationCustomdataForbidden handles this case with default header values.

forbidden
*/
type FindOrganizationCustomdataForbidden struct {
}

func (o *FindOrganizationCustomdataForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/customdata][%d] findOrganizationCustomdataForbidden ", 403)
}

func (o *FindOrganizationCustomdataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindOrganizationCustomdataNotFound creates a FindOrganizationCustomdataNotFound with default headers values
func NewFindOrganizationCustomdataNotFound() *FindOrganizationCustomdataNotFound {
	return &FindOrganizationCustomdataNotFound{}
}

/*FindOrganizationCustomdataNotFound handles this case with default header values.

not found
*/
type FindOrganizationCustomdataNotFound struct {
}

func (o *FindOrganizationCustomdataNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/customdata][%d] findOrganizationCustomdataNotFound ", 404)
}

func (o *FindOrganizationCustomdataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
