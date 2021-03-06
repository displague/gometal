// Code generated by go-swagger; DO NOT EDIT.

package facilities

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

// NewFindFacilitiesByProjectParams creates a new FindFacilitiesByProjectParams object
// with the default values initialized.
func NewFindFacilitiesByProjectParams() *FindFacilitiesByProjectParams {
	var ()
	return &FindFacilitiesByProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindFacilitiesByProjectParamsWithTimeout creates a new FindFacilitiesByProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindFacilitiesByProjectParamsWithTimeout(timeout time.Duration) *FindFacilitiesByProjectParams {
	var ()
	return &FindFacilitiesByProjectParams{

		timeout: timeout,
	}
}

// NewFindFacilitiesByProjectParamsWithContext creates a new FindFacilitiesByProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindFacilitiesByProjectParamsWithContext(ctx context.Context) *FindFacilitiesByProjectParams {
	var ()
	return &FindFacilitiesByProjectParams{

		Context: ctx,
	}
}

// NewFindFacilitiesByProjectParamsWithHTTPClient creates a new FindFacilitiesByProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindFacilitiesByProjectParamsWithHTTPClient(client *http.Client) *FindFacilitiesByProjectParams {
	var ()
	return &FindFacilitiesByProjectParams{
		HTTPClient: client,
	}
}

/*FindFacilitiesByProjectParams contains all the parameters to send to the API endpoint
for the find facilities by project operation typically these are written to a http.Request
*/
type FindFacilitiesByProjectParams struct {

	/*ID
	  Project UUID

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

// WithTimeout adds the timeout to the find facilities by project params
func (o *FindFacilitiesByProjectParams) WithTimeout(timeout time.Duration) *FindFacilitiesByProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find facilities by project params
func (o *FindFacilitiesByProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find facilities by project params
func (o *FindFacilitiesByProjectParams) WithContext(ctx context.Context) *FindFacilitiesByProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find facilities by project params
func (o *FindFacilitiesByProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find facilities by project params
func (o *FindFacilitiesByProjectParams) WithHTTPClient(client *http.Client) *FindFacilitiesByProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find facilities by project params
func (o *FindFacilitiesByProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find facilities by project params
func (o *FindFacilitiesByProjectParams) WithID(id strfmt.UUID) *FindFacilitiesByProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find facilities by project params
func (o *FindFacilitiesByProjectParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find facilities by project params
func (o *FindFacilitiesByProjectParams) WithInclude(include *string) *FindFacilitiesByProjectParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find facilities by project params
func (o *FindFacilitiesByProjectParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindFacilitiesByProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
