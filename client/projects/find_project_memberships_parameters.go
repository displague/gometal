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
	"github.com/go-openapi/swag"
)

// NewFindProjectMembershipsParams creates a new FindProjectMembershipsParams object
// with the default values initialized.
func NewFindProjectMembershipsParams() *FindProjectMembershipsParams {
	var ()
	return &FindProjectMembershipsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectMembershipsParamsWithTimeout creates a new FindProjectMembershipsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindProjectMembershipsParamsWithTimeout(timeout time.Duration) *FindProjectMembershipsParams {
	var ()
	return &FindProjectMembershipsParams{

		timeout: timeout,
	}
}

// NewFindProjectMembershipsParamsWithContext creates a new FindProjectMembershipsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindProjectMembershipsParamsWithContext(ctx context.Context) *FindProjectMembershipsParams {
	var ()
	return &FindProjectMembershipsParams{

		Context: ctx,
	}
}

// NewFindProjectMembershipsParamsWithHTTPClient creates a new FindProjectMembershipsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindProjectMembershipsParamsWithHTTPClient(client *http.Client) *FindProjectMembershipsParams {
	var ()
	return &FindProjectMembershipsParams{
		HTTPClient: client,
	}
}

/*FindProjectMembershipsParams contains all the parameters to send to the API endpoint
for the find project memberships operation typically these are written to a http.Request
*/
type FindProjectMembershipsParams struct {

	/*Include
	  related attributes to include

	*/
	Include *string
	/*Page
	  page to display, default to 1, max 100_000

	*/
	Page *int64
	/*PerPage
	  items per page, default to 10, max 1_000

	*/
	PerPage *int64
	/*ProjectID
	  Project UUID

	*/
	ProjectID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find project memberships params
func (o *FindProjectMembershipsParams) WithTimeout(timeout time.Duration) *FindProjectMembershipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find project memberships params
func (o *FindProjectMembershipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find project memberships params
func (o *FindProjectMembershipsParams) WithContext(ctx context.Context) *FindProjectMembershipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find project memberships params
func (o *FindProjectMembershipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find project memberships params
func (o *FindProjectMembershipsParams) WithHTTPClient(client *http.Client) *FindProjectMembershipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find project memberships params
func (o *FindProjectMembershipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the find project memberships params
func (o *FindProjectMembershipsParams) WithInclude(include *string) *FindProjectMembershipsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find project memberships params
func (o *FindProjectMembershipsParams) SetInclude(include *string) {
	o.Include = include
}

// WithPage adds the page to the find project memberships params
func (o *FindProjectMembershipsParams) WithPage(page *int64) *FindProjectMembershipsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find project memberships params
func (o *FindProjectMembershipsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the find project memberships params
func (o *FindProjectMembershipsParams) WithPerPage(perPage *int64) *FindProjectMembershipsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the find project memberships params
func (o *FindProjectMembershipsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithProjectID adds the projectID to the find project memberships params
func (o *FindProjectMembershipsParams) WithProjectID(projectID strfmt.UUID) *FindProjectMembershipsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the find project memberships params
func (o *FindProjectMembershipsParams) SetProjectID(projectID strfmt.UUID) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectMembershipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
