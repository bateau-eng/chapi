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

// NewPostAPIV3IterationsParams creates a new PostAPIV3IterationsParams object
// with the default values initialized.
func NewPostAPIV3IterationsParams() *PostAPIV3IterationsParams {
	var ()
	return &PostAPIV3IterationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3IterationsParamsWithTimeout creates a new PostAPIV3IterationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3IterationsParamsWithTimeout(timeout time.Duration) *PostAPIV3IterationsParams {
	var ()
	return &PostAPIV3IterationsParams{

		timeout: timeout,
	}
}

// NewPostAPIV3IterationsParamsWithContext creates a new PostAPIV3IterationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3IterationsParamsWithContext(ctx context.Context) *PostAPIV3IterationsParams {
	var ()
	return &PostAPIV3IterationsParams{

		Context: ctx,
	}
}

// NewPostAPIV3IterationsParamsWithHTTPClient creates a new PostAPIV3IterationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3IterationsParamsWithHTTPClient(client *http.Client) *PostAPIV3IterationsParams {
	var ()
	return &PostAPIV3IterationsParams{
		HTTPClient: client,
	}
}

/*PostAPIV3IterationsParams contains all the parameters to send to the API endpoint
for the post API v3 iterations operation typically these are written to a http.Request
*/
type PostAPIV3IterationsParams struct {

	/*CreateIteration*/
	CreateIteration *models.CreateIteration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) WithTimeout(timeout time.Duration) *PostAPIV3IterationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) WithContext(ctx context.Context) *PostAPIV3IterationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) WithHTTPClient(client *http.Client) *PostAPIV3IterationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateIteration adds the createIteration to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) WithCreateIteration(createIteration *models.CreateIteration) *PostAPIV3IterationsParams {
	o.SetCreateIteration(createIteration)
	return o
}

// SetCreateIteration adds the createIteration to the post API v3 iterations params
func (o *PostAPIV3IterationsParams) SetCreateIteration(createIteration *models.CreateIteration) {
	o.CreateIteration = createIteration
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3IterationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateIteration != nil {
		if err := r.SetBodyParam(o.CreateIteration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
