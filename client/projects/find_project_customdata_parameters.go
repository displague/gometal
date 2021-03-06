// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewFindProjectCustomdataParams creates a new FindProjectCustomdataParams object
// with the default values initialized.
func NewFindProjectCustomdataParams() *FindProjectCustomdataParams {
	var ()
	return &FindProjectCustomdataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectCustomdataParamsWithTimeout creates a new FindProjectCustomdataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindProjectCustomdataParamsWithTimeout(timeout time.Duration) *FindProjectCustomdataParams {
	var ()
	return &FindProjectCustomdataParams{

		timeout: timeout,
	}
}

// NewFindProjectCustomdataParamsWithContext creates a new FindProjectCustomdataParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindProjectCustomdataParamsWithContext(ctx context.Context) *FindProjectCustomdataParams {
	var ()
	return &FindProjectCustomdataParams{

		Context: ctx,
	}
}

// NewFindProjectCustomdataParamsWithHTTPClient creates a new FindProjectCustomdataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindProjectCustomdataParamsWithHTTPClient(client *http.Client) *FindProjectCustomdataParams {
	var ()
	return &FindProjectCustomdataParams{
		HTTPClient: client,
	}
}

/*FindProjectCustomdataParams contains all the parameters to send to the API endpoint
for the find project customdata operation typically these are written to a http.Request
*/
type FindProjectCustomdataParams struct {

	/*ID
	  Project UUID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find project customdata params
func (o *FindProjectCustomdataParams) WithTimeout(timeout time.Duration) *FindProjectCustomdataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find project customdata params
func (o *FindProjectCustomdataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find project customdata params
func (o *FindProjectCustomdataParams) WithContext(ctx context.Context) *FindProjectCustomdataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find project customdata params
func (o *FindProjectCustomdataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find project customdata params
func (o *FindProjectCustomdataParams) WithHTTPClient(client *http.Client) *FindProjectCustomdataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find project customdata params
func (o *FindProjectCustomdataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find project customdata params
func (o *FindProjectCustomdataParams) WithID(id strfmt.UUID) *FindProjectCustomdataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find project customdata params
func (o *FindProjectCustomdataParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectCustomdataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
