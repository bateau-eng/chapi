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

// NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams creates a new PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized.
func NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams() *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithTimeout creates a new PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithContext creates a new PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithHTTPClient creates a new PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 stories story public ID tasks task public ID operation typically these are written to a http.Request
*/
type PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams struct {

	/*UpdateTask*/
	UpdateTask *models.UpdateTask
	/*StoryPublicID
	  The unique identifier of the parent Story.

	*/
	StoryPublicID int64
	/*TaskPublicID
	  The unique identifier of the Task you wish to update.

	*/
	TaskPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateTask adds the updateTask to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithUpdateTask(updateTask *models.UpdateTask) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetUpdateTask(updateTask)
	return o
}

// SetUpdateTask adds the updateTask to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetUpdateTask(updateTask *models.UpdateTask) {
	o.UpdateTask = updateTask
}

// WithStoryPublicID adds the storyPublicID to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithStoryPublicID(storyPublicID int64) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WithTaskPublicID adds the taskPublicID to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WithTaskPublicID(taskPublicID int64) *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams {
	o.SetTaskPublicID(taskPublicID)
	return o
}

// SetTaskPublicID adds the taskPublicId to the put API v3 stories story public ID tasks task public ID params
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) SetTaskPublicID(taskPublicID int64) {
	o.TaskPublicID = taskPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3StoriesStoryPublicIDTasksTaskPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateTask != nil {
		if err := r.SetBodyParam(o.UpdateTask); err != nil {
			return err
		}
	}

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