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

	"github.com/t0mk/gometal/models"
)

// NewRequestBgpConfigParams creates a new RequestBgpConfigParams object
// with the default values initialized.
func NewRequestBgpConfigParams() *RequestBgpConfigParams {
	var ()
	return &RequestBgpConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestBgpConfigParamsWithTimeout creates a new RequestBgpConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestBgpConfigParamsWithTimeout(timeout time.Duration) *RequestBgpConfigParams {
	var ()
	return &RequestBgpConfigParams{

		timeout: timeout,
	}
}

// NewRequestBgpConfigParamsWithContext creates a new RequestBgpConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestBgpConfigParamsWithContext(ctx context.Context) *RequestBgpConfigParams {
	var ()
	return &RequestBgpConfigParams{

		Context: ctx,
	}
}

// NewRequestBgpConfigParamsWithHTTPClient creates a new RequestBgpConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestBgpConfigParamsWithHTTPClient(client *http.Client) *RequestBgpConfigParams {
	var ()
	return &RequestBgpConfigParams{
		HTTPClient: client,
	}
}

/*RequestBgpConfigParams contains all the parameters to send to the API endpoint
for the request bgp config operation typically these are written to a http.Request
*/
type RequestBgpConfigParams struct {

	/*BgpConfigRequest
	  BGP config Request to create

	*/
	BgpConfigRequest *models.BgpConfigRequestInput
	/*ID
	  Project UUID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request bgp config params
func (o *RequestBgpConfigParams) WithTimeout(timeout time.Duration) *RequestBgpConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request bgp config params
func (o *RequestBgpConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request bgp config params
func (o *RequestBgpConfigParams) WithContext(ctx context.Context) *RequestBgpConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request bgp config params
func (o *RequestBgpConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request bgp config params
func (o *RequestBgpConfigParams) WithHTTPClient(client *http.Client) *RequestBgpConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request bgp config params
func (o *RequestBgpConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBgpConfigRequest adds the bgpConfigRequest to the request bgp config params
func (o *RequestBgpConfigParams) WithBgpConfigRequest(bgpConfigRequest *models.BgpConfigRequestInput) *RequestBgpConfigParams {
	o.SetBgpConfigRequest(bgpConfigRequest)
	return o
}

// SetBgpConfigRequest adds the bgpConfigRequest to the request bgp config params
func (o *RequestBgpConfigParams) SetBgpConfigRequest(bgpConfigRequest *models.BgpConfigRequestInput) {
	o.BgpConfigRequest = bgpConfigRequest
}

// WithID adds the id to the request bgp config params
func (o *RequestBgpConfigParams) WithID(id strfmt.UUID) *RequestBgpConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the request bgp config params
func (o *RequestBgpConfigParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RequestBgpConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BgpConfigRequest != nil {
		if err := r.SetBodyParam(o.BgpConfigRequest); err != nil {
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