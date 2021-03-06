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

// NewGetAPIV3CategoriesCategoryPublicIDParams creates a new GetAPIV3CategoriesCategoryPublicIDParams object
// with the default values initialized.
func NewGetAPIV3CategoriesCategoryPublicIDParams() *GetAPIV3CategoriesCategoryPublicIDParams {
	var ()
	return &GetAPIV3CategoriesCategoryPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3CategoriesCategoryPublicIDParamsWithTimeout creates a new GetAPIV3CategoriesCategoryPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3CategoriesCategoryPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3CategoriesCategoryPublicIDParams {
	var ()
	return &GetAPIV3CategoriesCategoryPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3CategoriesCategoryPublicIDParamsWithContext creates a new GetAPIV3CategoriesCategoryPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3CategoriesCategoryPublicIDParamsWithContext(ctx context.Context) *GetAPIV3CategoriesCategoryPublicIDParams {
	var ()
	return &GetAPIV3CategoriesCategoryPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3CategoriesCategoryPublicIDParamsWithHTTPClient creates a new GetAPIV3CategoriesCategoryPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3CategoriesCategoryPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3CategoriesCategoryPublicIDParams {
	var ()
	return &GetAPIV3CategoriesCategoryPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3CategoriesCategoryPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 categories category public ID operation typically these are written to a http.Request
*/
type GetAPIV3CategoriesCategoryPublicIDParams struct {

	/*CategoryPublicID
	  The unique ID of the Category.

	*/
	CategoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3CategoriesCategoryPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) WithContext(ctx context.Context) *GetAPIV3CategoriesCategoryPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3CategoriesCategoryPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryPublicID adds the categoryPublicID to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) WithCategoryPublicID(categoryPublicID int64) *GetAPIV3CategoriesCategoryPublicIDParams {
	o.SetCategoryPublicID(categoryPublicID)
	return o
}

// SetCategoryPublicID adds the categoryPublicId to the get API v3 categories category public ID params
func (o *GetAPIV3CategoriesCategoryPublicIDParams) SetCategoryPublicID(categoryPublicID int64) {
	o.CategoryPublicID = categoryPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3CategoriesCategoryPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param category-public-id
	if err := r.SetPathParam("category-public-id", swag.FormatInt64(o.CategoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
