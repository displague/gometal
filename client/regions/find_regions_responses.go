// Code generated by go-swagger; DO NOT EDIT.

package regions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindRegionsReader is a Reader for the FindRegions structure.
type FindRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindRegionsOK creates a FindRegionsOK with default headers values
func NewFindRegionsOK() *FindRegionsOK {
	return &FindRegionsOK{}
}

/*FindRegionsOK handles this case with default header values.

ok
*/
type FindRegionsOK struct {
	Payload *types.RegionsList
}

func (o *FindRegionsOK) Error() string {
	return fmt.Sprintf("[GET /regions][%d] findRegionsOK  %+v", 200, o.Payload)
}

func (o *FindRegionsOK) GetPayload() *types.RegionsList {
	return o.Payload
}

func (o *FindRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.RegionsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindRegionsUnauthorized creates a FindRegionsUnauthorized with default headers values
func NewFindRegionsUnauthorized() *FindRegionsUnauthorized {
	return &FindRegionsUnauthorized{}
}

/*FindRegionsUnauthorized handles this case with default header values.

unauthorized
*/
type FindRegionsUnauthorized struct {
}

func (o *FindRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /regions][%d] findRegionsUnauthorized ", 401)
}

func (o *FindRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
