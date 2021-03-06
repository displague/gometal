// Code generated by go-swagger; DO NOT EDIT.

package invitations

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

// NewCreateProjectInvitationParams creates a new CreateProjectInvitationParams object
// with the default values initialized.
func NewCreateProjectInvitationParams() *CreateProjectInvitationParams {
	var ()
	return &CreateProjectInvitationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectInvitationParamsWithTimeout creates a new CreateProjectInvitationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectInvitationParamsWithTimeout(timeout time.Duration) *CreateProjectInvitationParams {
	var ()
	return &CreateProjectInvitationParams{

		timeout: timeout,
	}
}

// NewCreateProjectInvitationParamsWithContext creates a new CreateProjectInvitationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectInvitationParamsWithContext(ctx context.Context) *CreateProjectInvitationParams {
	var ()
	return &CreateProjectInvitationParams{

		Context: ctx,
	}
}

// NewCreateProjectInvitationParamsWithHTTPClient creates a new CreateProjectInvitationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectInvitationParamsWithHTTPClient(client *http.Client) *CreateProjectInvitationParams {
	var ()
	return &CreateProjectInvitationParams{
		HTTPClient: client,
	}
}

/*CreateProjectInvitationParams contains all the parameters to send to the API endpoint
for the create project invitation operation typically these are written to a http.Request
*/
type CreateProjectInvitationParams struct {

	/*Invitation
	  Invitation to create

	*/
	Invitation *types.InvitationInput
	/*ProjectID
	  Project UUID

	*/
	ProjectID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create project invitation params
func (o *CreateProjectInvitationParams) WithTimeout(timeout time.Duration) *CreateProjectInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project invitation params
func (o *CreateProjectInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project invitation params
func (o *CreateProjectInvitationParams) WithContext(ctx context.Context) *CreateProjectInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project invitation params
func (o *CreateProjectInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project invitation params
func (o *CreateProjectInvitationParams) WithHTTPClient(client *http.Client) *CreateProjectInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project invitation params
func (o *CreateProjectInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitation adds the invitation to the create project invitation params
func (o *CreateProjectInvitationParams) WithInvitation(invitation *types.InvitationInput) *CreateProjectInvitationParams {
	o.SetInvitation(invitation)
	return o
}

// SetInvitation adds the invitation to the create project invitation params
func (o *CreateProjectInvitationParams) SetInvitation(invitation *types.InvitationInput) {
	o.Invitation = invitation
}

// WithProjectID adds the projectID to the create project invitation params
func (o *CreateProjectInvitationParams) WithProjectID(projectID strfmt.UUID) *CreateProjectInvitationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create project invitation params
func (o *CreateProjectInvitationParams) SetProjectID(projectID strfmt.UUID) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Invitation != nil {
		if err := r.SetBodyParam(o.Invitation); err != nil {
			return err
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
