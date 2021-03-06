// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/bateau-eng/chapi/pkg/models"
)

// NewPutAPIV3ProjectsProjectPublicIDParams creates a new PutAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized.
func NewPutAPIV3ProjectsProjectPublicIDParams() *PutAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &PutAPIV3ProjectsProjectPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3ProjectsProjectPublicIDParamsWithTimeout creates a new PutAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3ProjectsProjectPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &PutAPIV3ProjectsProjectPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3ProjectsProjectPublicIDParamsWithContext creates a new PutAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3ProjectsProjectPublicIDParamsWithContext(ctx context.Context) *PutAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &PutAPIV3ProjectsProjectPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3ProjectsProjectPublicIDParamsWithHTTPClient creates a new PutAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3ProjectsProjectPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &PutAPIV3ProjectsProjectPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3ProjectsProjectPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 projects project public ID operation typically these are written to a http.Request
*/
type PutAPIV3ProjectsProjectPublicIDParams struct {

	/*UpdateProject*/
	UpdateProject *models.UpdateProject
	/*ProjectPublicID
	  The unique ID of the Project.

	*/
	ProjectPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3ProjectsProjectPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) WithContext(ctx context.Context) *PutAPIV3ProjectsProjectPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3ProjectsProjectPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateProject adds the updateProject to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) WithUpdateProject(updateProject *models.UpdateProject) *PutAPIV3ProjectsProjectPublicIDParams {
	o.SetUpdateProject(updateProject)
	return o
}

// SetUpdateProject adds the updateProject to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) SetUpdateProject(updateProject *models.UpdateProject) {
	o.UpdateProject = updateProject
}

// WithProjectPublicID adds the projectPublicID to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) WithProjectPublicID(projectPublicID int64) *PutAPIV3ProjectsProjectPublicIDParams {
	o.SetProjectPublicID(projectPublicID)
	return o
}

// SetProjectPublicID adds the projectPublicId to the put API v3 projects project public ID params
func (o *PutAPIV3ProjectsProjectPublicIDParams) SetProjectPublicID(projectPublicID int64) {
	o.ProjectPublicID = projectPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3ProjectsProjectPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateProject != nil {
		if err := r.SetBodyParam(o.UpdateProject); err != nil {
			return err
		}
	}

	// path param project-public-id
	if err := r.SetPathParam("project-public-id", swag.FormatInt64(o.ProjectPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
