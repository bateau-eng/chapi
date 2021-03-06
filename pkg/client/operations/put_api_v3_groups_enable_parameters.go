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

// NewPutAPIV3GroupsEnableParams creates a new PutAPIV3GroupsEnableParams object
// with the default values initialized.
func NewPutAPIV3GroupsEnableParams() *PutAPIV3GroupsEnableParams {

	return &PutAPIV3GroupsEnableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3GroupsEnableParamsWithTimeout creates a new PutAPIV3GroupsEnableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3GroupsEnableParamsWithTimeout(timeout time.Duration) *PutAPIV3GroupsEnableParams {

	return &PutAPIV3GroupsEnableParams{

		timeout: timeout,
	}
}

// NewPutAPIV3GroupsEnableParamsWithContext creates a new PutAPIV3GroupsEnableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3GroupsEnableParamsWithContext(ctx context.Context) *PutAPIV3GroupsEnableParams {

	return &PutAPIV3GroupsEnableParams{

		Context: ctx,
	}
}

// NewPutAPIV3GroupsEnableParamsWithHTTPClient creates a new PutAPIV3GroupsEnableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3GroupsEnableParamsWithHTTPClient(client *http.Client) *PutAPIV3GroupsEnableParams {

	return &PutAPIV3GroupsEnableParams{
		HTTPClient: client,
	}
}

/*PutAPIV3GroupsEnableParams contains all the parameters to send to the API endpoint
for the put API v3 groups enable operation typically these are written to a http.Request
*/
type PutAPIV3GroupsEnableParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) WithTimeout(timeout time.Duration) *PutAPIV3GroupsEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) WithContext(ctx context.Context) *PutAPIV3GroupsEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) WithHTTPClient(client *http.Client) *PutAPIV3GroupsEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 groups enable params
func (o *PutAPIV3GroupsEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3GroupsEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
