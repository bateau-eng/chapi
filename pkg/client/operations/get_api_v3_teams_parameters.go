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
)

// NewGetAPIV3TeamsParams creates a new GetAPIV3TeamsParams object
// with the default values initialized.
func NewGetAPIV3TeamsParams() *GetAPIV3TeamsParams {

	return &GetAPIV3TeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3TeamsParamsWithTimeout creates a new GetAPIV3TeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3TeamsParamsWithTimeout(timeout time.Duration) *GetAPIV3TeamsParams {

	return &GetAPIV3TeamsParams{

		timeout: timeout,
	}
}

// NewGetAPIV3TeamsParamsWithContext creates a new GetAPIV3TeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3TeamsParamsWithContext(ctx context.Context) *GetAPIV3TeamsParams {

	return &GetAPIV3TeamsParams{

		Context: ctx,
	}
}

// NewGetAPIV3TeamsParamsWithHTTPClient creates a new GetAPIV3TeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3TeamsParamsWithHTTPClient(client *http.Client) *GetAPIV3TeamsParams {

	return &GetAPIV3TeamsParams{
		HTTPClient: client,
	}
}

/*GetAPIV3TeamsParams contains all the parameters to send to the API endpoint
for the get API v3 teams operation typically these are written to a http.Request
*/
type GetAPIV3TeamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 teams params
func (o *GetAPIV3TeamsParams) WithTimeout(timeout time.Duration) *GetAPIV3TeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 teams params
func (o *GetAPIV3TeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 teams params
func (o *GetAPIV3TeamsParams) WithContext(ctx context.Context) *GetAPIV3TeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 teams params
func (o *GetAPIV3TeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 teams params
func (o *GetAPIV3TeamsParams) WithHTTPClient(client *http.Client) *GetAPIV3TeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 teams params
func (o *GetAPIV3TeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3TeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
