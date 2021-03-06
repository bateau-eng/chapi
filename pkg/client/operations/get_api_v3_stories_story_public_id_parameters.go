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

// NewGetAPIV3StoriesStoryPublicIDParams creates a new GetAPIV3StoriesStoryPublicIDParams object
// with the default values initialized.
func NewGetAPIV3StoriesStoryPublicIDParams() *GetAPIV3StoriesStoryPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3StoriesStoryPublicIDParamsWithTimeout creates a new GetAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3StoriesStoryPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3StoriesStoryPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3StoriesStoryPublicIDParamsWithContext creates a new GetAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3StoriesStoryPublicIDParamsWithContext(ctx context.Context) *GetAPIV3StoriesStoryPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3StoriesStoryPublicIDParamsWithHTTPClient creates a new GetAPIV3StoriesStoryPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3StoriesStoryPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3StoriesStoryPublicIDParams {
	var ()
	return &GetAPIV3StoriesStoryPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3StoriesStoryPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 stories story public ID operation typically these are written to a http.Request
*/
type GetAPIV3StoriesStoryPublicIDParams struct {

	/*StoryPublicID
	  The ID of the Story.

	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3StoriesStoryPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) WithContext(ctx context.Context) *GetAPIV3StoriesStoryPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3StoriesStoryPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoryPublicID adds the storyPublicID to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) WithStoryPublicID(storyPublicID int64) *GetAPIV3StoriesStoryPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the get API v3 stories story public ID params
func (o *GetAPIV3StoriesStoryPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3StoriesStoryPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
