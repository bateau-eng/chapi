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

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// NewPostAPIV3StoriesStoryPublicIDTasksParams creates a new PostAPIV3StoriesStoryPublicIDTasksParams object
// with the default values initialized.
func NewPostAPIV3StoriesStoryPublicIDTasksParams() *PostAPIV3StoriesStoryPublicIDTasksParams {
	var ()
	return &PostAPIV3StoriesStoryPublicIDTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3StoriesStoryPublicIDTasksParamsWithTimeout creates a new PostAPIV3StoriesStoryPublicIDTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3StoriesStoryPublicIDTasksParamsWithTimeout(timeout time.Duration) *PostAPIV3StoriesStoryPublicIDTasksParams {
	var ()
	return &PostAPIV3StoriesStoryPublicIDTasksParams{

		timeout: timeout,
	}
}

// NewPostAPIV3StoriesStoryPublicIDTasksParamsWithContext creates a new PostAPIV3StoriesStoryPublicIDTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3StoriesStoryPublicIDTasksParamsWithContext(ctx context.Context) *PostAPIV3StoriesStoryPublicIDTasksParams {
	var ()
	return &PostAPIV3StoriesStoryPublicIDTasksParams{

		Context: ctx,
	}
}

// NewPostAPIV3StoriesStoryPublicIDTasksParamsWithHTTPClient creates a new PostAPIV3StoriesStoryPublicIDTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3StoriesStoryPublicIDTasksParamsWithHTTPClient(client *http.Client) *PostAPIV3StoriesStoryPublicIDTasksParams {
	var ()
	return &PostAPIV3StoriesStoryPublicIDTasksParams{
		HTTPClient: client,
	}
}

/*PostAPIV3StoriesStoryPublicIDTasksParams contains all the parameters to send to the API endpoint
for the post API v3 stories story public ID tasks operation typically these are written to a http.Request
*/
type PostAPIV3StoriesStoryPublicIDTasksParams struct {

	/*CreateTask*/
	CreateTask *models.CreateTask
	/*StoryPublicID
	  The ID of the Story that the Task will be in.

	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WithTimeout(timeout time.Duration) *PostAPIV3StoriesStoryPublicIDTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WithContext(ctx context.Context) *PostAPIV3StoriesStoryPublicIDTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WithHTTPClient(client *http.Client) *PostAPIV3StoriesStoryPublicIDTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateTask adds the createTask to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WithCreateTask(createTask *models.CreateTask) *PostAPIV3StoriesStoryPublicIDTasksParams {
	o.SetCreateTask(createTask)
	return o
}

// SetCreateTask adds the createTask to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) SetCreateTask(createTask *models.CreateTask) {
	o.CreateTask = createTask
}

// WithStoryPublicID adds the storyPublicID to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WithStoryPublicID(storyPublicID int64) *PostAPIV3StoriesStoryPublicIDTasksParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the post API v3 stories story public ID tasks params
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3StoriesStoryPublicIDTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateTask != nil {
		if err := r.SetBodyParam(o.CreateTask); err != nil {
			return err
		}
	}

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
