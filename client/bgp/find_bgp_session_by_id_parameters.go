// Code generated by go-swagger; DO NOT EDIT.

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindBGPSessionByIDParams creates a new FindBGPSessionByIDParams object
// with the default values initialized.
func NewFindBGPSessionByIDParams() *FindBGPSessionByIDParams {
	var ()
	return &FindBGPSessionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindBGPSessionByIDParamsWithTimeout creates a new FindBGPSessionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindBGPSessionByIDParamsWithTimeout(timeout time.Duration) *FindBGPSessionByIDParams {
	var ()
	return &FindBGPSessionByIDParams{

		timeout: timeout,
	}
}

// NewFindBGPSessionByIDParamsWithContext creates a new FindBGPSessionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindBGPSessionByIDParamsWithContext(ctx context.Context) *FindBGPSessionByIDParams {
	var ()
	return &FindBGPSessionByIDParams{

		Context: ctx,
	}
}

// NewFindBGPSessionByIDParamsWithHTTPClient creates a new FindBGPSessionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindBGPSessionByIDParamsWithHTTPClient(client *http.Client) *FindBGPSessionByIDParams {
	var ()
	return &FindBGPSessionByIDParams{
		HTTPClient: client,
	}
}

/*FindBGPSessionByIDParams contains all the parameters to send to the API endpoint
for the find Bgp session by Id operation typically these are written to a http.Request
*/
type FindBGPSessionByIDParams struct {

	/*ID
	  BGP session UUID

	*/
	ID strfmt.UUID
	/*Include
	  related attributes to include

	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) WithTimeout(timeout time.Duration) *FindBGPSessionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) WithContext(ctx context.Context) *FindBGPSessionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) WithHTTPClient(client *http.Client) *FindBGPSessionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) WithID(id strfmt.UUID) *FindBGPSessionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) WithInclude(include *string) *FindBGPSessionByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find Bgp session by Id params
func (o *FindBGPSessionByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindBGPSessionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// query param include
		var qrInclude string
		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {
			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
