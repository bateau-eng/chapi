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

// NewPostAPIV3LinkedFilesParams creates a new PostAPIV3LinkedFilesParams object
// with the default values initialized.
func NewPostAPIV3LinkedFilesParams() *PostAPIV3LinkedFilesParams {
	var ()
	return &PostAPIV3LinkedFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3LinkedFilesParamsWithTimeout creates a new PostAPIV3LinkedFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3LinkedFilesParamsWithTimeout(timeout time.Duration) *PostAPIV3LinkedFilesParams {
	var ()
	return &PostAPIV3LinkedFilesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3LinkedFilesParamsWithContext creates a new PostAPIV3LinkedFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3LinkedFilesParamsWithContext(ctx context.Context) *PostAPIV3LinkedFilesParams {
	var ()
	return &PostAPIV3LinkedFilesParams{

		Context: ctx,
	}
}

// NewPostAPIV3LinkedFilesParamsWithHTTPClient creates a new PostAPIV3LinkedFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3LinkedFilesParamsWithHTTPClient(client *http.Client) *PostAPIV3LinkedFilesParams {
	var ()
	return &PostAPIV3LinkedFilesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3LinkedFilesParams contains all the parameters to send to the API endpoint
for the post API v3 linked files operation typically these are written to a http.Request
*/
type PostAPIV3LinkedFilesParams struct {

	/*CreateLinkedFile*/
	CreateLinkedFile *models.CreateLinkedFile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) WithTimeout(timeout time.Duration) *PostAPIV3LinkedFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) WithContext(ctx context.Context) *PostAPIV3LinkedFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) WithHTTPClient(client *http.Client) *PostAPIV3LinkedFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateLinkedFile adds the createLinkedFile to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) WithCreateLinkedFile(createLinkedFile *models.CreateLinkedFile) *PostAPIV3LinkedFilesParams {
	o.SetCreateLinkedFile(createLinkedFile)
	return o
}

// SetCreateLinkedFile adds the createLinkedFile to the post API v3 linked files params
func (o *PostAPIV3LinkedFilesParams) SetCreateLinkedFile(createLinkedFile *models.CreateLinkedFile) {
	o.CreateLinkedFile = createLinkedFile
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3LinkedFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateLinkedFile != nil {
		if err := r.SetBodyParam(o.CreateLinkedFile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
