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

// NewEnableTfaAppParams creates a new EnableTfaAppParams object
// with the default values initialized.
func NewEnableTfaAppParams() *EnableTfaAppParams {

	return &EnableTfaAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableTfaAppParamsWithTimeout creates a new EnableTfaAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableTfaAppParamsWithTimeout(timeout time.Duration) *EnableTfaAppParams {

	return &EnableTfaAppParams{

		timeout: timeout,
	}
}

// NewEnableTfaAppParamsWithContext creates a new EnableTfaAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableTfaAppParamsWithContext(ctx context.Context) *EnableTfaAppParams {

	return &EnableTfaAppParams{

		Context: ctx,
	}
}

// NewEnableTfaAppParamsWithHTTPClient creates a new EnableTfaAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableTfaAppParamsWithHTTPClient(client *http.Client) *EnableTfaAppParams {

	return &EnableTfaAppParams{
		HTTPClient: client,
	}
}

/*EnableTfaAppParams contains all the parameters to send to the API endpoint
for the enable tfa app operation typically these are written to a http.Request
*/
type EnableTfaAppParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable tfa app params
func (o *EnableTfaAppParams) WithTimeout(timeout time.Duration) *EnableTfaAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable tfa app params
func (o *EnableTfaAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable tfa app params
func (o *EnableTfaAppParams) WithContext(ctx context.Context) *EnableTfaAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable tfa app params
func (o *EnableTfaAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable tfa app params
func (o *EnableTfaAppParams) WithHTTPClient(client *http.Client) *EnableTfaAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable tfa app params
func (o *EnableTfaAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EnableTfaAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
