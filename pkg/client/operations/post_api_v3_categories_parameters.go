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

// NewPostAPIV3CategoriesParams creates a new PostAPIV3CategoriesParams object
// with the default values initialized.
func NewPostAPIV3CategoriesParams() *PostAPIV3CategoriesParams {
	var ()
	return &PostAPIV3CategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3CategoriesParamsWithTimeout creates a new PostAPIV3CategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3CategoriesParamsWithTimeout(timeout time.Duration) *PostAPIV3CategoriesParams {
	var ()
	return &PostAPIV3CategoriesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3CategoriesParamsWithContext creates a new PostAPIV3CategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3CategoriesParamsWithContext(ctx context.Context) *PostAPIV3CategoriesParams {
	var ()
	return &PostAPIV3CategoriesParams{

		Context: ctx,
	}
}

// NewPostAPIV3CategoriesParamsWithHTTPClient creates a new PostAPIV3CategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3CategoriesParamsWithHTTPClient(client *http.Client) *PostAPIV3CategoriesParams {
	var ()
	return &PostAPIV3CategoriesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3CategoriesParams contains all the parameters to send to the API endpoint
for the post API v3 categories operation typically these are written to a http.Request
*/
type PostAPIV3CategoriesParams struct {

	/*CreateCategory*/
	CreateCategory *models.CreateCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) WithTimeout(timeout time.Duration) *PostAPIV3CategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) WithContext(ctx context.Context) *PostAPIV3CategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) WithHTTPClient(client *http.Client) *PostAPIV3CategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateCategory adds the createCategory to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) WithCreateCategory(createCategory *models.CreateCategory) *PostAPIV3CategoriesParams {
	o.SetCreateCategory(createCategory)
	return o
}

// SetCreateCategory adds the createCategory to the post API v3 categories params
func (o *PostAPIV3CategoriesParams) SetCreateCategory(createCategory *models.CreateCategory) {
	o.CreateCategory = createCategory
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3CategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateCategory != nil {
		if err := r.SetBodyParam(o.CreateCategory); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
