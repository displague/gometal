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

// FindPlansByProjectReader is a Reader for the FindPlansByProject structure.
type FindPlansByProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindPlansByProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindPlansByProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindPlansByProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindPlansByProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindPlansByProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindPlansByProjectOK creates a FindPlansByProjectOK with default headers values
func NewFindPlansByProjectOK() *FindPlansByProjectOK {
	return &FindPlansByProjectOK{}
}

/*FindPlansByProjectOK handles this case with default header values.

ok
*/
type FindPlansByProjectOK struct {
	Payload *types.PlanList
}

func (o *FindPlansByProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/plans][%d] findPlansByProjectOK  %+v", 200, o.Payload)
}

func (o *FindPlansByProjectOK) GetPayload() *types.PlanList {
	return o.Payload
}

func (o *FindPlansByProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.PlanList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindPlansByProjectUnauthorized creates a FindPlansByProjectUnauthorized with default headers values
func NewFindPlansByProjectUnauthorized() *FindPlansByProjectUnauthorized {
	return &FindPlansByProjectUnauthorized{}
}

/*FindPlansByProjectUnauthorized handles this case with default header values.

unauthorized
*/
type FindPlansByProjectUnauthorized struct {
}

func (o *FindPlansByProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/plans][%d] findPlansByProjectUnauthorized ", 401)
}

func (o *FindPlansByProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindPlansByProjectForbidden creates a FindPlansByProjectForbidden with default headers values
func NewFindPlansByProjectForbidden() *FindPlansByProjectForbidden {
	return &FindPlansByProjectForbidden{}
}

/*FindPlansByProjectForbidden handles this case with default header values.

forbidden
*/
type FindPlansByProjectForbidden struct {
}

func (o *FindPlansByProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/plans][%d] findPlansByProjectForbidden ", 403)
}

func (o *FindPlansByProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindPlansByProjectNotFound creates a FindPlansByProjectNotFound with default headers values
func NewFindPlansByProjectNotFound() *FindPlansByProjectNotFound {
	return &FindPlansByProjectNotFound{}
}

/*FindPlansByProjectNotFound handles this case with default header values.

not found
*/
type FindPlansByProjectNotFound struct {
}

func (o *FindPlansByProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/plans][%d] findPlansByProjectNotFound ", 404)
}

func (o *FindPlansByProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
