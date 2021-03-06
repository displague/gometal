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

// NewFindBGPConfigByProjectParams creates a new FindBGPConfigByProjectParams object
// with the default values initialized.
func NewFindBGPConfigByProjectParams() *FindBGPConfigByProjectParams {
	var ()
	return &FindBGPConfigByProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindBGPConfigByProjectParamsWithTimeout creates a new FindBGPConfigByProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindBGPConfigByProjectParamsWithTimeout(timeout time.Duration) *FindBGPConfigByProjectParams {
	var ()
	return &FindBGPConfigByProjectParams{

		timeout: timeout,
	}
}

// NewFindBGPConfigByProjectParamsWithContext creates a new FindBGPConfigByProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindBGPConfigByProjectParamsWithContext(ctx context.Context) *FindBGPConfigByProjectParams {
	var ()
	return &FindBGPConfigByProjectParams{

		Context: ctx,
	}
}

// NewFindBGPConfigByProjectParamsWithHTTPClient creates a new FindBGPConfigByProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindBGPConfigByProjectParamsWithHTTPClient(client *http.Client) *FindBGPConfigByProjectParams {
	var ()
	return &FindBGPConfigByProjectParams{
		HTTPClient: client,
	}
}

/*FindBGPConfigByProjectParams contains all the parameters to send to the API endpoint
for the find Bgp config by project operation typically these are written to a http.Request
*/
type FindBGPConfigByProjectParams struct {

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

// WithTimeout adds the timeout to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) WithTimeout(timeout time.Duration) *FindBGPConfigByProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) WithContext(ctx context.Context) *FindBGPConfigByProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) WithHTTPClient(client *http.Client) *FindBGPConfigByProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) WithID(id strfmt.UUID) *FindBGPConfigByProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) WithInclude(include *string) *FindBGPConfigByProjectParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find Bgp config by project params
func (o *FindBGPConfigByProjectParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindBGPConfigByProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
