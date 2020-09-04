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

// NewGetAPIV3LinkedFilesLinkedFilePublicIDParams creates a new GetAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized.
func NewGetAPIV3LinkedFilesLinkedFilePublicIDParams() *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &GetAPIV3LinkedFilesLinkedFilePublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithTimeout creates a new GetAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &GetAPIV3LinkedFilesLinkedFilePublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithContext creates a new GetAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithContext(ctx context.Context) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &GetAPIV3LinkedFilesLinkedFilePublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithHTTPClient creates a new GetAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3LinkedFilesLinkedFilePublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &GetAPIV3LinkedFilesLinkedFilePublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3LinkedFilesLinkedFilePublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 linked files linked file public ID operation typically these are written to a http.Request
*/
type GetAPIV3LinkedFilesLinkedFilePublicIDParams struct {

	/*LinkedFilePublicID
	  The unique identifier of the linked file.

	*/
	LinkedFilePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) WithContext(ctx context.Context) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkedFilePublicID adds the linkedFilePublicID to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) WithLinkedFilePublicID(linkedFilePublicID int64) *GetAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetLinkedFilePublicID(linkedFilePublicID)
	return o
}

// SetLinkedFilePublicID adds the linkedFilePublicId to the get API v3 linked files linked file public ID params
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) SetLinkedFilePublicID(linkedFilePublicID int64) {
	o.LinkedFilePublicID = linkedFilePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3LinkedFilesLinkedFilePublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linked-file-public-id
	if err := r.SetPathParam("linked-file-public-id", swag.FormatInt64(o.LinkedFilePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}