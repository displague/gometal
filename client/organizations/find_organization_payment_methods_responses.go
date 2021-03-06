// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindOrganizationPaymentMethodsReader is a Reader for the FindOrganizationPaymentMethods structure.
type FindOrganizationPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOrganizationPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOrganizationPaymentMethodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindOrganizationPaymentMethodsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindOrganizationPaymentMethodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindOrganizationPaymentMethodsOK creates a FindOrganizationPaymentMethodsOK with default headers values
func NewFindOrganizationPaymentMethodsOK() *FindOrganizationPaymentMethodsOK {
	return &FindOrganizationPaymentMethodsOK{}
}

/*FindOrganizationPaymentMethodsOK handles this case with default header values.

ok
*/
type FindOrganizationPaymentMethodsOK struct {
	Payload *types.PaymentMethodList
}

func (o *FindOrganizationPaymentMethodsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/payment-methods][%d] findOrganizationPaymentMethodsOK  %+v", 200, o.Payload)
}

func (o *FindOrganizationPaymentMethodsOK) GetPayload() *types.PaymentMethodList {
	return o.Payload
}

func (o *FindOrganizationPaymentMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.PaymentMethodList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationPaymentMethodsUnauthorized creates a FindOrganizationPaymentMethodsUnauthorized with default headers values
func NewFindOrganizationPaymentMethodsUnauthorized() *FindOrganizationPaymentMethodsUnauthorized {
	return &FindOrganizationPaymentMethodsUnauthorized{}
}

/*FindOrganizationPaymentMethodsUnauthorized handles this case with default header values.

unauthorized
*/
type FindOrganizationPaymentMethodsUnauthorized struct {
}

func (o *FindOrganizationPaymentMethodsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/payment-methods][%d] findOrganizationPaymentMethodsUnauthorized ", 401)
}

func (o *FindOrganizationPaymentMethodsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindOrganizationPaymentMethodsNotFound creates a FindOrganizationPaymentMethodsNotFound with default headers values
func NewFindOrganizationPaymentMethodsNotFound() *FindOrganizationPaymentMethodsNotFound {
	return &FindOrganizationPaymentMethodsNotFound{}
}

/*FindOrganizationPaymentMethodsNotFound handles this case with default header values.

not found
*/
type FindOrganizationPaymentMethodsNotFound struct {
}

func (o *FindOrganizationPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{id}/payment-methods][%d] findOrganizationPaymentMethodsNotFound ", 404)
}

func (o *FindOrganizationPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
