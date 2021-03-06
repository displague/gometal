// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// DeleteInterconnectionReader is a Reader for the DeleteInterconnection structure.
type DeleteInterconnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInterconnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteInterconnectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInterconnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInterconnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInterconnectionAccepted creates a DeleteInterconnectionAccepted with default headers values
func NewDeleteInterconnectionAccepted() *DeleteInterconnectionAccepted {
	return &DeleteInterconnectionAccepted{}
}

/*DeleteInterconnectionAccepted handles this case with default header values.

accepted
*/
type DeleteInterconnectionAccepted struct {
	Payload *types.Interconnection
}

func (o *DeleteInterconnectionAccepted) Error() string {
	return fmt.Sprintf("[DELETE /connections/{connection_id}][%d] deleteInterconnectionAccepted  %+v", 202, o.Payload)
}

func (o *DeleteInterconnectionAccepted) GetPayload() *types.Interconnection {
	return o.Payload
}

func (o *DeleteInterconnectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Interconnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInterconnectionForbidden creates a DeleteInterconnectionForbidden with default headers values
func NewDeleteInterconnectionForbidden() *DeleteInterconnectionForbidden {
	return &DeleteInterconnectionForbidden{}
}

/*DeleteInterconnectionForbidden handles this case with default header values.

forbidden
*/
type DeleteInterconnectionForbidden struct {
}

func (o *DeleteInterconnectionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /connections/{connection_id}][%d] deleteInterconnectionForbidden ", 403)
}

func (o *DeleteInterconnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInterconnectionNotFound creates a DeleteInterconnectionNotFound with default headers values
func NewDeleteInterconnectionNotFound() *DeleteInterconnectionNotFound {
	return &DeleteInterconnectionNotFound{}
}

/*DeleteInterconnectionNotFound handles this case with default header values.

not found
*/
type DeleteInterconnectionNotFound struct {
}

func (o *DeleteInterconnectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /connections/{connection_id}][%d] deleteInterconnectionNotFound ", 404)
}

func (o *DeleteInterconnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
