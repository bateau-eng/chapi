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

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// NewPostAPIV3StoriesSearchParams creates a new PostAPIV3StoriesSearchParams object
// with the default values initialized.
func NewPostAPIV3StoriesSearchParams() *PostAPIV3StoriesSearchParams {
	var ()
	return &PostAPIV3StoriesSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3StoriesSearchParamsWithTimeout creates a new PostAPIV3StoriesSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3StoriesSearchParamsWithTimeout(timeout time.Duration) *PostAPIV3StoriesSearchParams {
	var ()
	return &PostAPIV3StoriesSearchParams{

		timeout: timeout,
	}
}

// NewPostAPIV3StoriesSearchParamsWithContext creates a new PostAPIV3StoriesSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3StoriesSearchParamsWithContext(ctx context.Context) *PostAPIV3StoriesSearchParams {
	var ()
	return &PostAPIV3StoriesSearchParams{

		Context: ctx,
	}
}

// NewPostAPIV3StoriesSearchParamsWithHTTPClient creates a new PostAPIV3StoriesSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3StoriesSearchParamsWithHTTPClient(client *http.Client) *PostAPIV3StoriesSearchParams {
	var ()
	return &PostAPIV3StoriesSearchParams{
		HTTPClient: client,
	}
}

/*PostAPIV3StoriesSearchParams contains all the parameters to send to the API endpoint
for the post API v3 stories search operation typically these are written to a http.Request
*/
type PostAPIV3StoriesSearchParams struct {

	/*SearchStories*/
	SearchStories *models.SearchStories

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) WithTimeout(timeout time.Duration) *PostAPIV3StoriesSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) WithContext(ctx context.Context) *PostAPIV3StoriesSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) WithHTTPClient(client *http.Client) *PostAPIV3StoriesSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchStories adds the searchStories to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) WithSearchStories(searchStories *models.SearchStories) *PostAPIV3StoriesSearchParams {
	o.SetSearchStories(searchStories)
	return o
}

// SetSearchStories adds the searchStories to the post API v3 stories search params
func (o *PostAPIV3StoriesSearchParams) SetSearchStories(searchStories *models.SearchStories) {
	o.SearchStories = searchStories
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3StoriesSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchStories != nil {
		if err := r.SetBodyParam(o.SearchStories); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
