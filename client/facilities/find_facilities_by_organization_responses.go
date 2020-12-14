// Code generated by go-swagger; DO NOT EDIT.

package facilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/models"
)

// FindFacilitiesByOrganizationReader is a Reader for the FindFacilitiesByOrganization structure.
type FindFacilitiesByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindFacilitiesByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindFacilitiesByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindFacilitiesByOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindFacilitiesByOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindFacilitiesByOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindFacilitiesByOrganizationOK creates a FindFacilitiesByOrganizationOK with default headers values
func NewFindFacilitiesByOrganizationOK() *FindFacilitiesByOrganizationOK {
	return &FindFacilitiesByOrganizationOK{}
}

/*FindFacilitiesByOrganizationOK handles this case with default header values.

ok
*/
type FindFacilitiesByOrganizationOK struct {
	Payload *models.FacilityList
}

func (o *FindFacilitiesByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/facilities][%d] findFacilitiesByOrganizationOK  %+v", 200, o.Payload)
}

func (o *FindFacilitiesByOrganizationOK) GetPayload() *models.FacilityList {
	return o.Payload
}

func (o *FindFacilitiesByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FacilityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindFacilitiesByOrganizationUnauthorized creates a FindFacilitiesByOrganizationUnauthorized with default headers values
func NewFindFacilitiesByOrganizationUnauthorized() *FindFacilitiesByOrganizationUnauthorized {
	return &FindFacilitiesByOrganizationUnauthorized{}
}

/*FindFacilitiesByOrganizationUnauthorized handles this case with default header values.

unauthorized
*/
type FindFacilitiesByOrganizationUnauthorized struct {
}

func (o *FindFacilitiesByOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/facilities][%d] findFacilitiesByOrganizationUnauthorized ", 401)
}

func (o *FindFacilitiesByOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindFacilitiesByOrganizationForbidden creates a FindFacilitiesByOrganizationForbidden with default headers values
func NewFindFacilitiesByOrganizationForbidden() *FindFacilitiesByOrganizationForbidden {
	return &FindFacilitiesByOrganizationForbidden{}
}

/*FindFacilitiesByOrganizationForbidden handles this case with default header values.

forbidden
*/
type FindFacilitiesByOrganizationForbidden struct {
}

func (o *FindFacilitiesByOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/facilities][%d] findFacilitiesByOrganizationForbidden ", 403)
}

func (o *FindFacilitiesByOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindFacilitiesByOrganizationNotFound creates a FindFacilitiesByOrganizationNotFound with default headers values
func NewFindFacilitiesByOrganizationNotFound() *FindFacilitiesByOrganizationNotFound {
	return &FindFacilitiesByOrganizationNotFound{}
}

/*FindFacilitiesByOrganizationNotFound handles this case with default header values.

not found
*/
type FindFacilitiesByOrganizationNotFound struct {
}

func (o *FindFacilitiesByOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/facilities][%d] findFacilitiesByOrganizationNotFound ", 404)
}

func (o *FindFacilitiesByOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}