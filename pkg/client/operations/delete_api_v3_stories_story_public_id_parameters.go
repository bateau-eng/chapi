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

// NewDeleteAPIV3StoriesStoryPublicIDParams creates a new DeleteAPIV3StoriesStoryPublicIDParams object
// with the default values initialized.
func NewDeleteAPIV3StoriesStoryPublicIDParams() *DeleteAPIV3StoriesStoryPublicIDParams {
	var ()
	return &DeleteAPIV3StoriesStoryPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3StoriesStoryPublicIDParamsWithTimeout creates a new DeleteAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3StoriesStoryPublicIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV3StoriesStoryPublicIDParams {
	var ()
	return &DeleteAPIV3StoriesStoryPublicIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3StoriesStoryPublicIDParamsWithContext creates a new DeleteAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3StoriesStoryPublicIDParamsWithContext(ctx context.Context) *DeleteAPIV3StoriesStoryPublicIDParams {
	var ()
	return &DeleteAPIV3StoriesStoryPublicIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3StoriesStoryPublicIDParamsWithHTTPClient creates a new DeleteAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3StoriesStoryPublicIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV3StoriesStoryPublicIDParams {
	var ()
	return &DeleteAPIV3StoriesStoryPublicIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3StoriesStoryPublicIDParams contains all the parameters to send to the API endpoint
for the delete API v3 stories story public ID operation typically these are written to a http.Request
*/
type DeleteAPIV3StoriesStoryPublicIDParams struct {

	/*StoryPublicID
	  The ID of the Story.

	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV3StoriesStoryPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) WithContext(ctx context.Context) *DeleteAPIV3StoriesStoryPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV3StoriesStoryPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) WithStoryPublicID(storyPublicID int64) *DeleteAPIV3StoriesStoryPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the delete API v3 stories story public ID params
func (o *DeleteAPIV3StoriesStoryPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3StoriesStoryPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
