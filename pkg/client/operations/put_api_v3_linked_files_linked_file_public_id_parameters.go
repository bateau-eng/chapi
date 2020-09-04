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

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// NewPutAPIV3LinkedFilesLinkedFilePublicIDParams creates a new PutAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized.
func NewPutAPIV3LinkedFilesLinkedFilePublicIDParams() *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &PutAPIV3LinkedFilesLinkedFilePublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithTimeout creates a new PutAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &PutAPIV3LinkedFilesLinkedFilePublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithContext creates a new PutAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithContext(ctx context.Context) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &PutAPIV3LinkedFilesLinkedFilePublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithHTTPClient creates a new PutAPIV3LinkedFilesLinkedFilePublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3LinkedFilesLinkedFilePublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	var ()
	return &PutAPIV3LinkedFilesLinkedFilePublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3LinkedFilesLinkedFilePublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 linked files linked file public ID operation typically these are written to a http.Request
*/
type PutAPIV3LinkedFilesLinkedFilePublicIDParams struct {

	/*UpdateLinkedFile*/
	UpdateLinkedFile *models.UpdateLinkedFile
	/*LinkedFilePublicID
	  The unique identifier of the linked file.

	*/
	LinkedFilePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WithContext(ctx context.Context) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateLinkedFile adds the updateLinkedFile to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WithUpdateLinkedFile(updateLinkedFile *models.UpdateLinkedFile) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetUpdateLinkedFile(updateLinkedFile)
	return o
}

// SetUpdateLinkedFile adds the updateLinkedFile to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) SetUpdateLinkedFile(updateLinkedFile *models.UpdateLinkedFile) {
	o.UpdateLinkedFile = updateLinkedFile
}

// WithLinkedFilePublicID adds the linkedFilePublicID to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WithLinkedFilePublicID(linkedFilePublicID int64) *PutAPIV3LinkedFilesLinkedFilePublicIDParams {
	o.SetLinkedFilePublicID(linkedFilePublicID)
	return o
}

// SetLinkedFilePublicID adds the linkedFilePublicId to the put API v3 linked files linked file public ID params
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) SetLinkedFilePublicID(linkedFilePublicID int64) {
	o.LinkedFilePublicID = linkedFilePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3LinkedFilesLinkedFilePublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateLinkedFile != nil {
		if err := r.SetBodyParam(o.UpdateLinkedFile); err != nil {
			return err
		}
	}

	// path param linked-file-public-id
	if err := r.SetPathParam("linked-file-public-id", swag.FormatInt64(o.LinkedFilePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
