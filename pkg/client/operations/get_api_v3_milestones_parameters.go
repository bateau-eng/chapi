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

// NewGetAPIV3MilestonesParams creates a new GetAPIV3MilestonesParams object
// with the default values initialized.
func NewGetAPIV3MilestonesParams() *GetAPIV3MilestonesParams {

	return &GetAPIV3MilestonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MilestonesParamsWithTimeout creates a new GetAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MilestonesParamsWithTimeout(timeout time.Duration) *GetAPIV3MilestonesParams {

	return &GetAPIV3MilestonesParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MilestonesParamsWithContext creates a new GetAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MilestonesParamsWithContext(ctx context.Context) *GetAPIV3MilestonesParams {

	return &GetAPIV3MilestonesParams{

		Context: ctx,
	}
}

// NewGetAPIV3MilestonesParamsWithHTTPClient creates a new GetAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MilestonesParamsWithHTTPClient(client *http.Client) *GetAPIV3MilestonesParams {

	return &GetAPIV3MilestonesParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MilestonesParams contains all the parameters to send to the API endpoint
for the get API v3 milestones operation typically these are written to a http.Request
*/
type GetAPIV3MilestonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) WithTimeout(timeout time.Duration) *GetAPIV3MilestonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) WithContext(ctx context.Context) *GetAPIV3MilestonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) WithHTTPClient(client *http.Client) *GetAPIV3MilestonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 milestones params
func (o *GetAPIV3MilestonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MilestonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
