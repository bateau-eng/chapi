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

// NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams creates a new GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized.
func NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams() *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithTimeout creates a new GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithContext creates a new GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithContext(ctx context.Context) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithHTTPClient creates a new GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 stories story public ID tasks task public ID operation typically these are written to a http.Request
*/
type GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams struct {

	/*StoryPublicID
	  The unique ID of the Story this Task is associated with.

	*/
	StoryPublicID int64
	/*TaskPublicID
	  The unique ID of the Task.

	*/
	TaskPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithContext(ctx context.Context) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithStoryPublicID(storyPublicID int64) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WithTaskPublicID adds the taskPublicID to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithTaskPublicID(taskPublicID int64) *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetTaskPublicID(taskPublicID)
	return o
}

// SetTaskPublicID adds the taskPublicId to the get API v3 stories story public ID tasks task public ID params
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetTaskPublicID(taskPublicID int64) {
	o.TaskPublicID = taskPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	// path param task-public-id
	if err := r.SetPathParam("task-public-id", swag.FormatInt64(o.TaskPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
