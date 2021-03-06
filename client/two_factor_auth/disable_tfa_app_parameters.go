// Code generated by go-swagger; DO NOT EDIT.

package two_factor_auth

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

// NewDisableTfaAppParams creates a new DisableTfaAppParams object
// with the default values initialized.
func NewDisableTfaAppParams() *DisableTfaAppParams {

	return &DisableTfaAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableTfaAppParamsWithTimeout creates a new DisableTfaAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableTfaAppParamsWithTimeout(timeout time.Duration) *DisableTfaAppParams {

	return &DisableTfaAppParams{

		timeout: timeout,
	}
}

// NewDisableTfaAppParamsWithContext creates a new DisableTfaAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableTfaAppParamsWithContext(ctx context.Context) *DisableTfaAppParams {

	return &DisableTfaAppParams{

		Context: ctx,
	}
}

// NewDisableTfaAppParamsWithHTTPClient creates a new DisableTfaAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableTfaAppParamsWithHTTPClient(client *http.Client) *DisableTfaAppParams {

	return &DisableTfaAppParams{
		HTTPClient: client,
	}
}

/*DisableTfaAppParams contains all the parameters to send to the API endpoint
for the disable tfa app operation typically these are written to a http.Request
*/
type DisableTfaAppParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable tfa app params
func (o *DisableTfaAppParams) WithTimeout(timeout time.Duration) *DisableTfaAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable tfa app params
func (o *DisableTfaAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable tfa app params
func (o *DisableTfaAppParams) WithContext(ctx context.Context) *DisableTfaAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable tfa app params
func (o *DisableTfaAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable tfa app params
func (o *DisableTfaAppParams) WithHTTPClient(client *http.Client) *DisableTfaAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable tfa app params
func (o *DisableTfaAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DisableTfaAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
