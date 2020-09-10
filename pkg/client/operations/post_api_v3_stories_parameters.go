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

	"github.com/bateau-eng/chapi/pkg/models"
)

// NewPostAPIV3StoriesParams creates a new PostAPIV3StoriesParams object
// with the default values initialized.
func NewPostAPIV3StoriesParams() *PostAPIV3StoriesParams {
	var ()
	return &PostAPIV3StoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3StoriesParamsWithTimeout creates a new PostAPIV3StoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3StoriesParamsWithTimeout(timeout time.Duration) *PostAPIV3StoriesParams {
	var ()
	return &PostAPIV3StoriesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3StoriesParamsWithContext creates a new PostAPIV3StoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3StoriesParamsWithContext(ctx context.Context) *PostAPIV3StoriesParams {
	var ()
	return &PostAPIV3StoriesParams{

		Context: ctx,
	}
}

// NewPostAPIV3StoriesParamsWithHTTPClient creates a new PostAPIV3StoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3StoriesParamsWithHTTPClient(client *http.Client) *PostAPIV3StoriesParams {
	var ()
	return &PostAPIV3StoriesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3StoriesParams contains all the parameters to send to the API endpoint
for the post API v3 stories operation typically these are written to a http.Request
*/
type PostAPIV3StoriesParams struct {

	/*CreateStoryParams
	  Used to create multiple stories in a single request.

	*/
	CreateStoryParams *models.CreateStoryParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 stories params
func (o *PostAPIV3StoriesParams) WithTimeout(timeout time.Duration) *PostAPIV3StoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 stories params
func (o *PostAPIV3StoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 stories params
func (o *PostAPIV3StoriesParams) WithContext(ctx context.Context) *PostAPIV3StoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 stories params
func (o *PostAPIV3StoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 stories params
func (o *PostAPIV3StoriesParams) WithHTTPClient(client *http.Client) *PostAPIV3StoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 stories params
func (o *PostAPIV3StoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateStoryParams adds the createStoryParams to the post API v3 stories params
func (o *PostAPIV3StoriesParams) WithCreateStoryParams(createStoryParams *models.CreateStoryParams) *PostAPIV3StoriesParams {
	o.SetCreateStoryParams(createStoryParams)
	return o
}

// SetCreateStoryParams adds the createStoryParams to the post API v3 stories params
func (o *PostAPIV3StoriesParams) SetCreateStoryParams(createStoryParams *models.CreateStoryParams) {
	o.CreateStoryParams = createStoryParams
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3StoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateStoryParams != nil {
		if err := r.SetBodyParam(o.CreateStoryParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
