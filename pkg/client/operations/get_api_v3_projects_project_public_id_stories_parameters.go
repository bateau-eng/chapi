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

// NewGetAPIV3ProjectsProjectPublicIDStoriesParams creates a new GetAPIV3ProjectsProjectPublicIDStoriesParams object
// with the default values initialized.
func NewGetAPIV3ProjectsProjectPublicIDStoriesParams() *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	var ()
	return &GetAPIV3ProjectsProjectPublicIDStoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithTimeout creates a new GetAPIV3ProjectsProjectPublicIDStoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithTimeout(timeout time.Duration) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	var ()
	return &GetAPIV3ProjectsProjectPublicIDStoriesParams{

		timeout: timeout,
	}
}

// NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithContext creates a new GetAPIV3ProjectsProjectPublicIDStoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithContext(ctx context.Context) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	var ()
	return &GetAPIV3ProjectsProjectPublicIDStoriesParams{

		Context: ctx,
	}
}

// NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithHTTPClient creates a new GetAPIV3ProjectsProjectPublicIDStoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3ProjectsProjectPublicIDStoriesParamsWithHTTPClient(client *http.Client) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	var ()
	return &GetAPIV3ProjectsProjectPublicIDStoriesParams{
		HTTPClient: client,
	}
}

/*GetAPIV3ProjectsProjectPublicIDStoriesParams contains all the parameters to send to the API endpoint
for the get API v3 projects project public ID stories operation typically these are written to a http.Request
*/
type GetAPIV3ProjectsProjectPublicIDStoriesParams struct {

	/*GetProjectStories*/
	GetProjectStories *models.GetProjectStories
	/*ProjectPublicID
	  The unique ID of the Project.

	*/
	ProjectPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WithTimeout(timeout time.Duration) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WithContext(ctx context.Context) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WithHTTPClient(client *http.Client) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetProjectStories adds the getProjectStories to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WithGetProjectStories(getProjectStories *models.GetProjectStories) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	o.SetGetProjectStories(getProjectStories)
	return o
}

// SetGetProjectStories adds the getProjectStories to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) SetGetProjectStories(getProjectStories *models.GetProjectStories) {
	o.GetProjectStories = getProjectStories
}

// WithProjectPublicID adds the projectPublicID to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WithProjectPublicID(projectPublicID int64) *GetAPIV3ProjectsProjectPublicIDStoriesParams {
	o.SetProjectPublicID(projectPublicID)
	return o
}

// SetProjectPublicID adds the projectPublicId to the get API v3 projects project public ID stories params
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) SetProjectPublicID(projectPublicID int64) {
	o.ProjectPublicID = projectPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3ProjectsProjectPublicIDStoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GetProjectStories != nil {
		if err := r.SetBodyParam(o.GetProjectStories); err != nil {
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
