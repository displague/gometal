// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindPlansByOrganizationReader is a Reader for the FindPlansByOrganization structure.
type FindPlansByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPlansByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPlansByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindPlansByOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindPlansByOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindPlansByOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindPlansByOrganizationOK creates a FindPlansByOrganizationOK with default headers values
func NewFindPlansByOrganizationOK() *FindPlansByOrganizationOK {
	return &FindPlansByOrganizationOK{}
}

/*FindPlansByOrganizationOK handles this case with default header values.

ok
*/
type FindPlansByOrganizationOK struct {
	Payload *types.PlanList
}

func (o *FindPlansByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/plans][%d] findPlansByOrganizationOK  %+v", 200, o.Payload)
}

func (o *FindPlansByOrganizationOK) GetPayload() *types.PlanList {
	return o.Payload
}

func (o *FindPlansByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.PlanList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPlansByOrganizationUnauthorized creates a FindPlansByOrganizationUnauthorized with default headers values
func NewFindPlansByOrganizationUnauthorized() *FindPlansByOrganizationUnauthorized {
	return &FindPlansByOrganizationUnauthorized{}
}

/*FindPlansByOrganizationUnauthorized handles this case with default header values.

unauthorized
*/
type FindPlansByOrganizationUnauthorized struct {
}

func (o *FindPlansByOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/plans][%d] findPlansByOrganizationUnauthorized ", 401)
}

func (o *FindPlansByOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindPlansByOrganizationForbidden creates a FindPlansByOrganizationForbidden with default headers values
func NewFindPlansByOrganizationForbidden() *FindPlansByOrganizationForbidden {
	return &FindPlansByOrganizationForbidden{}
}

/*FindPlansByOrganizationForbidden handles this case with default header values.

forbidden
*/
type FindPlansByOrganizationForbidden struct {
}

func (o *FindPlansByOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/plans][%d] findPlansByOrganizationForbidden ", 403)
}

func (o *FindPlansByOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindPlansByOrganizationNotFound creates a FindPlansByOrganizationNotFound with default headers values
func NewFindPlansByOrganizationNotFound() *FindPlansByOrganizationNotFound {
	return &FindPlansByOrganizationNotFound{}
}

/*FindPlansByOrganizationNotFound handles this case with default header values.

not found
*/
type FindPlansByOrganizationNotFound struct {
}

func (o *FindPlansByOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/plans][%d] findPlansByOrganizationNotFound ", 404)
}

func (o *FindPlansByOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
