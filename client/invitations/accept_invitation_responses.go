// Code generated by go-swagger; DO NOT EDIT.

package invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// AcceptInvitationReader is a Reader for the AcceptInvitation structure.
type AcceptInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAcceptInvitationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAcceptInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAcceptInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptInvitationOK creates a AcceptInvitationOK with default headers values
func NewAcceptInvitationOK() *AcceptInvitationOK {
	return &AcceptInvitationOK{}
}

/*AcceptInvitationOK handles this case with default header values.

ok
*/
type AcceptInvitationOK struct {
	Payload *types.Membership
}

func (o *AcceptInvitationOK) Error() string {
	return fmt.Sprintf("[PUT /invitations/{id}][%d] acceptInvitationOK  %+v", 200, o.Payload)
}

func (o *AcceptInvitationOK) GetPayload() *types.Membership {
	return o.Payload
}

func (o *AcceptInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Membership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAcceptInvitationUnauthorized creates a AcceptInvitationUnauthorized with default headers values
func NewAcceptInvitationUnauthorized() *AcceptInvitationUnauthorized {
	return &AcceptInvitationUnauthorized{}
}

/*AcceptInvitationUnauthorized handles this case with default header values.

unauthorized
*/
type AcceptInvitationUnauthorized struct {
}

func (o *AcceptInvitationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /invitations/{id}][%d] acceptInvitationUnauthorized ", 401)
}

func (o *AcceptInvitationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAcceptInvitationForbidden creates a AcceptInvitationForbidden with default headers values
func NewAcceptInvitationForbidden() *AcceptInvitationForbidden {
	return &AcceptInvitationForbidden{}
}

/*AcceptInvitationForbidden handles this case with default header values.

forbidden
*/
type AcceptInvitationForbidden struct {
}

func (o *AcceptInvitationForbidden) Error() string {
	return fmt.Sprintf("[PUT /invitations/{id}][%d] acceptInvitationForbidden ", 403)
}

func (o *AcceptInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAcceptInvitationNotFound creates a AcceptInvitationNotFound with default headers values
func NewAcceptInvitationNotFound() *AcceptInvitationNotFound {
	return &AcceptInvitationNotFound{}
}

/*AcceptInvitationNotFound handles this case with default header values.

not found
*/
type AcceptInvitationNotFound struct {
}

func (o *AcceptInvitationNotFound) Error() string {
	return fmt.Sprintf("[PUT /invitations/{id}][%d] acceptInvitationNotFound ", 404)
}

func (o *AcceptInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
