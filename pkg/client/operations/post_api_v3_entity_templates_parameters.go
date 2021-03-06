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

// NewPostAPIV3EntityTemplatesParams creates a new PostAPIV3EntityTemplatesParams object
// with the default values initialized.
func NewPostAPIV3EntityTemplatesParams() *PostAPIV3EntityTemplatesParams {
	var ()
	return &PostAPIV3EntityTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3EntityTemplatesParamsWithTimeout creates a new PostAPIV3EntityTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3EntityTemplatesParamsWithTimeout(timeout time.Duration) *PostAPIV3EntityTemplatesParams {
	var ()
	return &PostAPIV3EntityTemplatesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3EntityTemplatesParamsWithContext creates a new PostAPIV3EntityTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3EntityTemplatesParamsWithContext(ctx context.Context) *PostAPIV3EntityTemplatesParams {
	var ()
	return &PostAPIV3EntityTemplatesParams{

		Context: ctx,
	}
}

// NewPostAPIV3EntityTemplatesParamsWithHTTPClient creates a new PostAPIV3EntityTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3EntityTemplatesParamsWithHTTPClient(client *http.Client) *PostAPIV3EntityTemplatesParams {
	var ()
	return &PostAPIV3EntityTemplatesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3EntityTemplatesParams contains all the parameters to send to the API endpoint
for the post API v3 entity templates operation typically these are written to a http.Request
*/
type PostAPIV3EntityTemplatesParams struct {

	/*CreateEntityTemplate
	  Request paramaters for creating an entirely new entity template.

	*/
	CreateEntityTemplate *models.CreateEntityTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) WithTimeout(timeout time.Duration) *PostAPIV3EntityTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) WithContext(ctx context.Context) *PostAPIV3EntityTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) WithHTTPClient(client *http.Client) *PostAPIV3EntityTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateEntityTemplate adds the createEntityTemplate to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) WithCreateEntityTemplate(createEntityTemplate *models.CreateEntityTemplate) *PostAPIV3EntityTemplatesParams {
	o.SetCreateEntityTemplate(createEntityTemplate)
	return o
}

// SetCreateEntityTemplate adds the createEntityTemplate to the post API v3 entity templates params
func (o *PostAPIV3EntityTemplatesParams) SetCreateEntityTemplate(createEntityTemplate *models.CreateEntityTemplate) {
	o.CreateEntityTemplate = createEntityTemplate
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3EntityTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateEntityTemplate != nil {
		if err := r.SetBodyParam(o.CreateEntityTemplate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
