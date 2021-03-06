// Code generated by go-swagger; DO NOT EDIT.

package otps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FindEnsureOtpReader is a Reader for the FindEnsureOtp structure.
type FindEnsureOtpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindEnsureOtpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewFindEnsureOtpNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFindEnsureOtpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFindEnsureOtpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindEnsureOtpNoContent creates a FindEnsureOtpNoContent with default headers values
func NewFindEnsureOtpNoContent() *FindEnsureOtpNoContent {
	return &FindEnsureOtpNoContent{}
}

/*FindEnsureOtpNoContent handles this case with default header values.

no content
*/
type FindEnsureOtpNoContent struct {
}

func (o *FindEnsureOtpNoContent) Error() string {
	return fmt.Sprintf("[POST /user/otp/verify/{otp}][%d] findEnsureOtpNoContent ", 204)
}

func (o *FindEnsureOtpNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindEnsureOtpBadRequest creates a FindEnsureOtpBadRequest with default headers values
func NewFindEnsureOtpBadRequest() *FindEnsureOtpBadRequest {
	return &FindEnsureOtpBadRequest{}
}

/*FindEnsureOtpBadRequest handles this case with default header values.

bad request
*/
type FindEnsureOtpBadRequest struct {
}

func (o *FindEnsureOtpBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/otp/verify/{otp}][%d] findEnsureOtpBadRequest ", 400)
}

func (o *FindEnsureOtpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindEnsureOtpUnauthorized creates a FindEnsureOtpUnauthorized with default headers values
func NewFindEnsureOtpUnauthorized() *FindEnsureOtpUnauthorized {
	return &FindEnsureOtpUnauthorized{}
}

/*FindEnsureOtpUnauthorized handles this case with default header values.

unauthorized
*/
type FindEnsureOtpUnauthorized struct {
}

func (o *FindEnsureOtpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/otp/verify/{otp}][%d] findEnsureOtpUnauthorized ", 401)
}

func (o *FindEnsureOtpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
