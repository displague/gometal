// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*DeleteProjectNoContent handles this case with default header values.

no content
*/
type DeleteProjectNoContent struct {
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectNoContent ", 204)
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/*DeleteProjectUnauthorized handles this case with default header values.

unauthorized
*/
type DeleteProjectUnauthorized struct {
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectUnauthorized ", 401)
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*DeleteProjectForbidden handles this case with default header values.

forbidden
*/
type DeleteProjectForbidden struct {
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*DeleteProjectNotFound handles this case with default header values.

not found
*/
type DeleteProjectNotFound struct {
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id}][%d] deleteProjectNotFound ", 404)
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
