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
)

// NewDeleteAPIV3ProjectsProjectPublicIDParams creates a new DeleteAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized.
func NewDeleteAPIV3ProjectsProjectPublicIDParams() *DeleteAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &DeleteAPIV3ProjectsProjectPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3ProjectsProjectPublicIDParamsWithTimeout creates a new DeleteAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3ProjectsProjectPublicIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &DeleteAPIV3ProjectsProjectPublicIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3ProjectsProjectPublicIDParamsWithContext creates a new DeleteAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3ProjectsProjectPublicIDParamsWithContext(ctx context.Context) *DeleteAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &DeleteAPIV3ProjectsProjectPublicIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3ProjectsProjectPublicIDParamsWithHTTPClient creates a new DeleteAPIV3ProjectsProjectPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3ProjectsProjectPublicIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV3ProjectsProjectPublicIDParams {
	var ()
	return &DeleteAPIV3ProjectsProjectPublicIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3ProjectsProjectPublicIDParams contains all the parameters to send to the API endpoint
for the delete API v3 projects project public ID operation typically these are written to a http.Request
*/
type DeleteAPIV3ProjectsProjectPublicIDParams struct {

	/*ProjectPublicID
	  The unique ID of the Project.

	*/
	ProjectPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV3ProjectsProjectPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) WithContext(ctx context.Context) *DeleteAPIV3ProjectsProjectPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV3ProjectsProjectPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectPublicID adds the projectPublicID to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) WithProjectPublicID(projectPublicID int64) *DeleteAPIV3ProjectsProjectPublicIDParams {
	o.SetProjectPublicID(projectPublicID)
	return o
}

// SetProjectPublicID adds the projectPublicId to the delete API v3 projects project public ID params
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) SetProjectPublicID(projectPublicID int64) {
	o.ProjectPublicID = projectPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3ProjectsProjectPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project-public-id
	if err := r.SetPathParam("project-public-id", swag.FormatInt64(o.ProjectPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
