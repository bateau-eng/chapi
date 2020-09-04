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

// NewPostAPIV3FilesParams creates a new PostAPIV3FilesParams object
// with the default values initialized.
func NewPostAPIV3FilesParams() *PostAPIV3FilesParams {
	var ()
	return &PostAPIV3FilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3FilesParamsWithTimeout creates a new PostAPIV3FilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3FilesParamsWithTimeout(timeout time.Duration) *PostAPIV3FilesParams {
	var ()
	return &PostAPIV3FilesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3FilesParamsWithContext creates a new PostAPIV3FilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3FilesParamsWithContext(ctx context.Context) *PostAPIV3FilesParams {
	var ()
	return &PostAPIV3FilesParams{

		Context: ctx,
	}
}

// NewPostAPIV3FilesParamsWithHTTPClient creates a new PostAPIV3FilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3FilesParamsWithHTTPClient(client *http.Client) *PostAPIV3FilesParams {
	var ()
	return &PostAPIV3FilesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3FilesParams contains all the parameters to send to the API endpoint
for the post API v3 files operation typically these are written to a http.Request
*/
type PostAPIV3FilesParams struct {

	/*CreateFiles*/
	CreateFiles *models.CreateFiles

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 files params
func (o *PostAPIV3FilesParams) WithTimeout(timeout time.Duration) *PostAPIV3FilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 files params
func (o *PostAPIV3FilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 files params
func (o *PostAPIV3FilesParams) WithContext(ctx context.Context) *PostAPIV3FilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 files params
func (o *PostAPIV3FilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 files params
func (o *PostAPIV3FilesParams) WithHTTPClient(client *http.Client) *PostAPIV3FilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 files params
func (o *PostAPIV3FilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateFiles adds the createFiles to the post API v3 files params
func (o *PostAPIV3FilesParams) WithCreateFiles(createFiles *models.CreateFiles) *PostAPIV3FilesParams {
	o.SetCreateFiles(createFiles)
	return o
}

// SetCreateFiles adds the createFiles to the post API v3 files params
func (o *PostAPIV3FilesParams) SetCreateFiles(createFiles *models.CreateFiles) {
	o.CreateFiles = createFiles
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3FilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateFiles != nil {
		if err := r.SetBodyParam(o.CreateFiles); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}