// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationNoContent creates a DeleteOrganizationNoContent with default headers values
func NewDeleteOrganizationNoContent() *DeleteOrganizationNoContent {
	return &DeleteOrganizationNoContent{}
}

/*DeleteOrganizationNoContent handles this case with default header values.

no content
*/
type DeleteOrganizationNoContent struct {
}

func (o *DeleteOrganizationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{id}][%d] deleteOrganizationNoContent ", 204)
}

func (o *DeleteOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationUnauthorized creates a DeleteOrganizationUnauthorized with default headers values
func NewDeleteOrganizationUnauthorized() *DeleteOrganizationUnauthorized {
	return &DeleteOrganizationUnauthorized{}
}

/*DeleteOrganizationUnauthorized handles this case with default header values.

unauthorized
*/
type DeleteOrganizationUnauthorized struct {
}

func (o *DeleteOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{id}][%d] deleteOrganizationUnauthorized ", 401)
}

func (o *DeleteOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationNotFound creates a DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {
	return &DeleteOrganizationNotFound{}
}

/*DeleteOrganizationNotFound handles this case with default header values.

not found
*/
type DeleteOrganizationNotFound struct {
}

func (o *DeleteOrganizationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{id}][%d] deleteOrganizationNotFound ", 404)
}

func (o *DeleteOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
