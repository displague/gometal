// Code generated by go-swagger; DO NOT EDIT.

package userdata

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

// NewValidateUserdataParams creates a new ValidateUserdataParams object
// with the default values initialized.
func NewValidateUserdataParams() *ValidateUserdataParams {
	var ()
	return &ValidateUserdataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUserdataParamsWithTimeout creates a new ValidateUserdataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUserdataParamsWithTimeout(timeout time.Duration) *ValidateUserdataParams {
	var ()
	return &ValidateUserdataParams{

		timeout: timeout,
	}
}

// NewValidateUserdataParamsWithContext creates a new ValidateUserdataParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUserdataParamsWithContext(ctx context.Context) *ValidateUserdataParams {
	var ()
	return &ValidateUserdataParams{

		Context: ctx,
	}
}

// NewValidateUserdataParamsWithHTTPClient creates a new ValidateUserdataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUserdataParamsWithHTTPClient(client *http.Client) *ValidateUserdataParams {
	var ()
	return &ValidateUserdataParams{
		HTTPClient: client,
	}
}

/*ValidateUserdataParams contains all the parameters to send to the API endpoint
for the validate userdata operation typically these are written to a http.Request
*/
type ValidateUserdataParams struct {

	/*Userdata
	  Userdata to validate

	*/
	Userdata *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate userdata params
func (o *ValidateUserdataParams) WithTimeout(timeout time.Duration) *ValidateUserdataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate userdata params
func (o *ValidateUserdataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate userdata params
func (o *ValidateUserdataParams) WithContext(ctx context.Context) *ValidateUserdataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate userdata params
func (o *ValidateUserdataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate userdata params
func (o *ValidateUserdataParams) WithHTTPClient(client *http.Client) *ValidateUserdataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate userdata params
func (o *ValidateUserdataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserdata adds the userdata to the validate userdata params
func (o *ValidateUserdataParams) WithUserdata(userdata *string) *ValidateUserdataParams {
	o.SetUserdata(userdata)
	return o
}

// SetUserdata adds the userdata to the validate userdata params
func (o *ValidateUserdataParams) SetUserdata(userdata *string) {
	o.Userdata = userdata
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUserdataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Userdata != nil {

		// query param userdata
		var qrUserdata string
		if o.Userdata != nil {
			qrUserdata = *o.Userdata
		}
		qUserdata := qrUserdata
		if qUserdata != "" {
			if err := r.SetQueryParam("userdata", qUserdata); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
