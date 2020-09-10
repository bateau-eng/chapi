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

// NewGetAPIV3SearchStoriesParams creates a new GetAPIV3SearchStoriesParams object
// with the default values initialized.
func NewGetAPIV3SearchStoriesParams() *GetAPIV3SearchStoriesParams {
	var ()
	return &GetAPIV3SearchStoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3SearchStoriesParamsWithTimeout creates a new GetAPIV3SearchStoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3SearchStoriesParamsWithTimeout(timeout time.Duration) *GetAPIV3SearchStoriesParams {
	var ()
	return &GetAPIV3SearchStoriesParams{

		timeout: timeout,
	}
}

// NewGetAPIV3SearchStoriesParamsWithContext creates a new GetAPIV3SearchStoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3SearchStoriesParamsWithContext(ctx context.Context) *GetAPIV3SearchStoriesParams {
	var ()
	return &GetAPIV3SearchStoriesParams{

		Context: ctx,
	}
}

// NewGetAPIV3SearchStoriesParamsWithHTTPClient creates a new GetAPIV3SearchStoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3SearchStoriesParamsWithHTTPClient(client *http.Client) *GetAPIV3SearchStoriesParams {
	var ()
	return &GetAPIV3SearchStoriesParams{
		HTTPClient: client,
	}
}

/*GetAPIV3SearchStoriesParams contains all the parameters to send to the API endpoint
for the get API v3 search stories operation typically these are written to a http.Request
*/
type GetAPIV3SearchStoriesParams struct {

	/*Search*/
	Search *models.Search

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) WithTimeout(timeout time.Duration) *GetAPIV3SearchStoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) WithContext(ctx context.Context) *GetAPIV3SearchStoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) WithHTTPClient(client *http.Client) *GetAPIV3SearchStoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearch adds the search to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) WithSearch(search *models.Search) *GetAPIV3SearchStoriesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get API v3 search stories params
func (o *GetAPIV3SearchStoriesParams) SetSearch(search *models.Search) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3SearchStoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Search != nil {
		if err := r.SetBodyParam(o.Search); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
