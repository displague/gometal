// Code generated by go-swagger; DO NOT EDIT.

package otps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// RegenerateCodesReader is a Reader for the RegenerateCodes structure.
type RegenerateCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenerateCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegenerateCodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegenerateCodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRegenerateCodesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegenerateCodesOK creates a RegenerateCodesOK with default headers values
func NewRegenerateCodesOK() *RegenerateCodesOK {
	return &RegenerateCodesOK{}
}

/*RegenerateCodesOK handles this case with default header values.

ok
*/
type RegenerateCodesOK struct {
	Payload *types.RecoveryCodeList
}

func (o *RegenerateCodesOK) Error() string {
	return fmt.Sprintf("[POST /user/otp/recovery-codes][%d] regenerateCodesOK  %+v", 200, o.Payload)
}

func (o *RegenerateCodesOK) GetPayload() *types.RecoveryCodeList {
	return o.Payload
}

func (o *RegenerateCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.RecoveryCodeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateCodesUnauthorized creates a RegenerateCodesUnauthorized with default headers values
func NewRegenerateCodesUnauthorized() *RegenerateCodesUnauthorized {
	return &RegenerateCodesUnauthorized{}
}

/*RegenerateCodesUnauthorized handles this case with default header values.

unauthorized
*/
type RegenerateCodesUnauthorized struct {
}

func (o *RegenerateCodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/otp/recovery-codes][%d] regenerateCodesUnauthorized ", 401)
}

func (o *RegenerateCodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateCodesNotFound creates a RegenerateCodesNotFound with default headers values
func NewRegenerateCodesNotFound() *RegenerateCodesNotFound {
	return &RegenerateCodesNotFound{}
}

/*RegenerateCodesNotFound handles this case with default header values.

not found
*/
type RegenerateCodesNotFound struct {
}

func (o *RegenerateCodesNotFound) Error() string {
	return fmt.Sprintf("[POST /user/otp/recovery-codes][%d] regenerateCodesNotFound ", 404)
}

func (o *RegenerateCodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateCodesUnprocessableEntity creates a RegenerateCodesUnprocessableEntity with default headers values
func NewRegenerateCodesUnprocessableEntity() *RegenerateCodesUnprocessableEntity {
	return &RegenerateCodesUnprocessableEntity{}
}

/*RegenerateCodesUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type RegenerateCodesUnprocessableEntity struct {
}

func (o *RegenerateCodesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /user/otp/recovery-codes][%d] regenerateCodesUnprocessableEntity ", 422)
}

func (o *RegenerateCodesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
