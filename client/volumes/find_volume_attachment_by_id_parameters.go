// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewFindVolumeAttachmentByIDParams creates a new FindVolumeAttachmentByIDParams object
// with the default values initialized.
func NewFindVolumeAttachmentByIDParams() *FindVolumeAttachmentByIDParams {
	var ()
	return &FindVolumeAttachmentByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindVolumeAttachmentByIDParamsWithTimeout creates a new FindVolumeAttachmentByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindVolumeAttachmentByIDParamsWithTimeout(timeout time.Duration) *FindVolumeAttachmentByIDParams {
	var ()
	return &FindVolumeAttachmentByIDParams{

		timeout: timeout,
	}
}

// NewFindVolumeAttachmentByIDParamsWithContext creates a new FindVolumeAttachmentByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindVolumeAttachmentByIDParamsWithContext(ctx context.Context) *FindVolumeAttachmentByIDParams {
	var ()
	return &FindVolumeAttachmentByIDParams{

		Context: ctx,
	}
}

// NewFindVolumeAttachmentByIDParamsWithHTTPClient creates a new FindVolumeAttachmentByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindVolumeAttachmentByIDParamsWithHTTPClient(client *http.Client) *FindVolumeAttachmentByIDParams {
	var ()
	return &FindVolumeAttachmentByIDParams{
		HTTPClient: client,
	}
}

/*FindVolumeAttachmentByIDParams contains all the parameters to send to the API endpoint
for the find volume attachment by Id operation typically these are written to a http.Request
*/
type FindVolumeAttachmentByIDParams struct {

	/*ID
	  Attachment UUID

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

// WithTimeout adds the timeout to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) WithTimeout(timeout time.Duration) *FindVolumeAttachmentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) WithContext(ctx context.Context) *FindVolumeAttachmentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) WithHTTPClient(client *http.Client) *FindVolumeAttachmentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) WithID(id strfmt.UUID) *FindVolumeAttachmentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) WithInclude(include *string) *FindVolumeAttachmentByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find volume attachment by Id params
func (o *FindVolumeAttachmentByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindVolumeAttachmentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
