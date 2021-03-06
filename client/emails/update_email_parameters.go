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

	"github.com/t0mk/gometal/types"
)

// NewUpdateEmailParams creates a new UpdateEmailParams object
// with the default values initialized.
func NewUpdateEmailParams() *UpdateEmailParams {
	var ()
	return &UpdateEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEmailParamsWithTimeout creates a new UpdateEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEmailParamsWithTimeout(timeout time.Duration) *UpdateEmailParams {
	var ()
	return &UpdateEmailParams{

		timeout: timeout,
	}
}

// NewUpdateEmailParamsWithContext creates a new UpdateEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEmailParamsWithContext(ctx context.Context) *UpdateEmailParams {
	var ()
	return &UpdateEmailParams{

		Context: ctx,
	}
}

// NewUpdateEmailParamsWithHTTPClient creates a new UpdateEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEmailParamsWithHTTPClient(client *http.Client) *UpdateEmailParams {
	var ()
	return &UpdateEmailParams{
		HTTPClient: client,
	}
}

/*UpdateEmailParams contains all the parameters to send to the API endpoint
for the update email operation typically these are written to a http.Request
*/
type UpdateEmailParams struct {

	/*Email
	  email to update

	*/
	Email *types.UpdateEmailInput
	/*ID
	  Email UUID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update email params
func (o *UpdateEmailParams) WithTimeout(timeout time.Duration) *UpdateEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update email params
func (o *UpdateEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update email params
func (o *UpdateEmailParams) WithContext(ctx context.Context) *UpdateEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update email params
func (o *UpdateEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update email params
func (o *UpdateEmailParams) WithHTTPClient(client *http.Client) *UpdateEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update email params
func (o *UpdateEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the update email params
func (o *UpdateEmailParams) WithEmail(email *types.UpdateEmailInput) *UpdateEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the update email params
func (o *UpdateEmailParams) SetEmail(email *types.UpdateEmailInput) {
	o.Email = email
}

// WithID adds the id to the update email params
func (o *UpdateEmailParams) WithID(id strfmt.UUID) *UpdateEmailParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update email params
func (o *UpdateEmailParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {
		if err := r.SetBodyParam(o.Email); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
