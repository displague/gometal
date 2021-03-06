// Code generated by go-swagger; DO NOT EDIT.

package emails

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

// NewDeleteEmailParams creates a new DeleteEmailParams object
// with the default values initialized.
func NewDeleteEmailParams() *DeleteEmailParams {
	var ()
	return &DeleteEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEmailParamsWithTimeout creates a new DeleteEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEmailParamsWithTimeout(timeout time.Duration) *DeleteEmailParams {
	var ()
	return &DeleteEmailParams{

		timeout: timeout,
	}
}

// NewDeleteEmailParamsWithContext creates a new DeleteEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEmailParamsWithContext(ctx context.Context) *DeleteEmailParams {
	var ()
	return &DeleteEmailParams{

		Context: ctx,
	}
}

// NewDeleteEmailParamsWithHTTPClient creates a new DeleteEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEmailParamsWithHTTPClient(client *http.Client) *DeleteEmailParams {
	var ()
	return &DeleteEmailParams{
		HTTPClient: client,
	}
}

/*DeleteEmailParams contains all the parameters to send to the API endpoint
for the delete email operation typically these are written to a http.Request
*/
type DeleteEmailParams struct {

	/*ID
	  Email UUID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete email params
func (o *DeleteEmailParams) WithTimeout(timeout time.Duration) *DeleteEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete email params
func (o *DeleteEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete email params
func (o *DeleteEmailParams) WithContext(ctx context.Context) *DeleteEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete email params
func (o *DeleteEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete email params
func (o *DeleteEmailParams) WithHTTPClient(client *http.Client) *DeleteEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete email params
func (o *DeleteEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete email params
func (o *DeleteEmailParams) WithID(id strfmt.UUID) *DeleteEmailParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete email params
func (o *DeleteEmailParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
