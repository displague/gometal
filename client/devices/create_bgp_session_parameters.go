// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewCreateBgpSessionParams creates a new CreateBgpSessionParams object
// with the default values initialized.
func NewCreateBgpSessionParams() *CreateBgpSessionParams {
	var ()
	return &CreateBgpSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBgpSessionParamsWithTimeout creates a new CreateBgpSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBgpSessionParamsWithTimeout(timeout time.Duration) *CreateBgpSessionParams {
	var ()
	return &CreateBgpSessionParams{

		timeout: timeout,
	}
}

// NewCreateBgpSessionParamsWithContext creates a new CreateBgpSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBgpSessionParamsWithContext(ctx context.Context) *CreateBgpSessionParams {
	var ()
	return &CreateBgpSessionParams{

		Context: ctx,
	}
}

// NewCreateBgpSessionParamsWithHTTPClient creates a new CreateBgpSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBgpSessionParamsWithHTTPClient(client *http.Client) *CreateBgpSessionParams {
	var ()
	return &CreateBgpSessionParams{
		HTTPClient: client,
	}
}

/*CreateBgpSessionParams contains all the parameters to send to the API endpoint
for the create bgp session operation typically these are written to a http.Request
*/
type CreateBgpSessionParams struct {

	/*BgpSession
	  BGP session to create

	*/
	BgpSession *models.BGPSessionInput
	/*ID
	  Device UUID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create bgp session params
func (o *CreateBgpSessionParams) WithTimeout(timeout time.Duration) *CreateBgpSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bgp session params
func (o *CreateBgpSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bgp session params
func (o *CreateBgpSessionParams) WithContext(ctx context.Context) *CreateBgpSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bgp session params
func (o *CreateBgpSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bgp session params
func (o *CreateBgpSessionParams) WithHTTPClient(client *http.Client) *CreateBgpSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bgp session params
func (o *CreateBgpSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBgpSession adds the bgpSession to the create bgp session params
func (o *CreateBgpSessionParams) WithBgpSession(bgpSession *models.BGPSessionInput) *CreateBgpSessionParams {
	o.SetBgpSession(bgpSession)
	return o
}

// SetBgpSession adds the bgpSession to the create bgp session params
func (o *CreateBgpSessionParams) SetBgpSession(bgpSession *models.BGPSessionInput) {
	o.BgpSession = bgpSession
}

// WithID adds the id to the create bgp session params
func (o *CreateBgpSessionParams) WithID(id strfmt.UUID) *CreateBgpSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create bgp session params
func (o *CreateBgpSessionParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBgpSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BgpSession != nil {
		if err := r.SetBodyParam(o.BgpSession); err != nil {
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