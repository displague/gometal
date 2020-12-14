// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/models"
)

// FindConnectionEventsReader is a Reader for the FindConnectionEvents structure.
type FindConnectionEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConnectionEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConnectionEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindConnectionEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindConnectionEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindConnectionEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindConnectionEventsOK creates a FindConnectionEventsOK with default headers values
func NewFindConnectionEventsOK() *FindConnectionEventsOK {
	return &FindConnectionEventsOK{}
}

/*FindConnectionEventsOK handles this case with default header values.

ok
*/
type FindConnectionEventsOK struct {
	Payload *models.Event
}

func (o *FindConnectionEventsOK) Error() string {
	return fmt.Sprintf("[GET /connections/{id}/events][%d] findConnectionEventsOK  %+v", 200, o.Payload)
}

func (o *FindConnectionEventsOK) GetPayload() *models.Event {
	return o.Payload
}

func (o *FindConnectionEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConnectionEventsUnauthorized creates a FindConnectionEventsUnauthorized with default headers values
func NewFindConnectionEventsUnauthorized() *FindConnectionEventsUnauthorized {
	return &FindConnectionEventsUnauthorized{}
}

/*FindConnectionEventsUnauthorized handles this case with default header values.

unauthorized
*/
type FindConnectionEventsUnauthorized struct {
}

func (o *FindConnectionEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /connections/{id}/events][%d] findConnectionEventsUnauthorized ", 401)
}

func (o *FindConnectionEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindConnectionEventsForbidden creates a FindConnectionEventsForbidden with default headers values
func NewFindConnectionEventsForbidden() *FindConnectionEventsForbidden {
	return &FindConnectionEventsForbidden{}
}

/*FindConnectionEventsForbidden handles this case with default header values.

forbidden
*/
type FindConnectionEventsForbidden struct {
}

func (o *FindConnectionEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /connections/{id}/events][%d] findConnectionEventsForbidden ", 403)
}

func (o *FindConnectionEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindConnectionEventsNotFound creates a FindConnectionEventsNotFound with default headers values
func NewFindConnectionEventsNotFound() *FindConnectionEventsNotFound {
	return &FindConnectionEventsNotFound{}
}

/*FindConnectionEventsNotFound handles this case with default header values.

not found
*/
type FindConnectionEventsNotFound struct {
}

func (o *FindConnectionEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /connections/{id}/events][%d] findConnectionEventsNotFound ", 404)
}

func (o *FindConnectionEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}