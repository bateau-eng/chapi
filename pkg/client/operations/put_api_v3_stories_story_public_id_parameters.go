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

// NewPutAPIV3StoriesStoryPublicIDParams creates a new PutAPIV3StoriesStoryPublicIDParams object
// with the default values initialized.
func NewPutAPIV3StoriesStoryPublicIDParams() *PutAPIV3StoriesStoryPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDParamsWithTimeout creates a new PutAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3StoriesStoryPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDParamsWithContext creates a new PutAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3StoriesStoryPublicIDParamsWithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3StoriesStoryPublicIDParamsWithHTTPClient creates a new PutAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3StoriesStoryPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3StoriesStoryPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 stories story public ID operation typically these are written to a http.Request
*/
type PutAPIV3StoriesStoryPublicIDParams struct {

	/*UpdateStory*/
	UpdateStory *models.UpdateStory
	/*StoryPublicID
	  The unique identifier of this story.

	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) WithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateStory adds the updateStory to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) WithUpdateStory(updateStory *models.UpdateStory) *PutAPIV3StoriesStoryPublicIDParams {
	o.SetUpdateStory(updateStory)
	return o
}

// SetUpdateStory adds the updateStory to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) SetUpdateStory(updateStory *models.UpdateStory) {
	o.UpdateStory = updateStory
}

// WithStoryPublicID adds the storyPublicID to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) WithStoryPublicID(storyPublicID int64) *PutAPIV3StoriesStoryPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the put API v3 stories story public ID params
func (o *PutAPIV3StoriesStoryPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3StoriesStoryPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateStory != nil {
		if err := r.SetBodyParam(o.UpdateStory); err != nil {
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
