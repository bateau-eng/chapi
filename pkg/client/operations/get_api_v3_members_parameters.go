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

// NewGetAPIV3MembersParams creates a new GetAPIV3MembersParams object
// with the default values initialized.
func NewGetAPIV3MembersParams() *GetAPIV3MembersParams {
	var ()
	return &GetAPIV3MembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MembersParamsWithTimeout creates a new GetAPIV3MembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MembersParamsWithTimeout(timeout time.Duration) *GetAPIV3MembersParams {
	var ()
	return &GetAPIV3MembersParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MembersParamsWithContext creates a new GetAPIV3MembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MembersParamsWithContext(ctx context.Context) *GetAPIV3MembersParams {
	var ()
	return &GetAPIV3MembersParams{

		Context: ctx,
	}
}

// NewGetAPIV3MembersParamsWithHTTPClient creates a new GetAPIV3MembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MembersParamsWithHTTPClient(client *http.Client) *GetAPIV3MembersParams {
	var ()
	return &GetAPIV3MembersParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MembersParams contains all the parameters to send to the API endpoint
for the get API v3 members operation typically these are written to a http.Request
*/
type GetAPIV3MembersParams struct {

	/*ListMembers*/
	ListMembers *models.ListMembers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 members params
func (o *GetAPIV3MembersParams) WithTimeout(timeout time.Duration) *GetAPIV3MembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 members params
func (o *GetAPIV3MembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 members params
func (o *GetAPIV3MembersParams) WithContext(ctx context.Context) *GetAPIV3MembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 members params
func (o *GetAPIV3MembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 members params
func (o *GetAPIV3MembersParams) WithHTTPClient(client *http.Client) *GetAPIV3MembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 members params
func (o *GetAPIV3MembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListMembers adds the listMembers to the get API v3 members params
func (o *GetAPIV3MembersParams) WithListMembers(listMembers *models.ListMembers) *GetAPIV3MembersParams {
	o.SetListMembers(listMembers)
	return o
}

// SetListMembers adds the listMembers to the get API v3 members params
func (o *GetAPIV3MembersParams) SetListMembers(listMembers *models.ListMembers) {
	o.ListMembers = listMembers
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ListMembers != nil {
		if err := r.SetBodyParam(o.ListMembers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}