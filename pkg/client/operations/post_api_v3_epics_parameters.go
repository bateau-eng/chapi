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

// NewPostAPIV3EpicsParams creates a new PostAPIV3EpicsParams object
// with the default values initialized.
func NewPostAPIV3EpicsParams() *PostAPIV3EpicsParams {
	var ()
	return &PostAPIV3EpicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3EpicsParamsWithTimeout creates a new PostAPIV3EpicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3EpicsParamsWithTimeout(timeout time.Duration) *PostAPIV3EpicsParams {
	var ()
	return &PostAPIV3EpicsParams{

		timeout: timeout,
	}
}

// NewPostAPIV3EpicsParamsWithContext creates a new PostAPIV3EpicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3EpicsParamsWithContext(ctx context.Context) *PostAPIV3EpicsParams {
	var ()
	return &PostAPIV3EpicsParams{

		Context: ctx,
	}
}

// NewPostAPIV3EpicsParamsWithHTTPClient creates a new PostAPIV3EpicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3EpicsParamsWithHTTPClient(client *http.Client) *PostAPIV3EpicsParams {
	var ()
	return &PostAPIV3EpicsParams{
		HTTPClient: client,
	}
}

/*PostAPIV3EpicsParams contains all the parameters to send to the API endpoint
for the post API v3 epics operation typically these are written to a http.Request
*/
type PostAPIV3EpicsParams struct {

	/*CreateEpic*/
	CreateEpic *models.CreateEpic

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 epics params
func (o *PostAPIV3EpicsParams) WithTimeout(timeout time.Duration) *PostAPIV3EpicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 epics params
func (o *PostAPIV3EpicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 epics params
func (o *PostAPIV3EpicsParams) WithContext(ctx context.Context) *PostAPIV3EpicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 epics params
func (o *PostAPIV3EpicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 epics params
func (o *PostAPIV3EpicsParams) WithHTTPClient(client *http.Client) *PostAPIV3EpicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 epics params
func (o *PostAPIV3EpicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateEpic adds the createEpic to the post API v3 epics params
func (o *PostAPIV3EpicsParams) WithCreateEpic(createEpic *models.CreateEpic) *PostAPIV3EpicsParams {
	o.SetCreateEpic(createEpic)
	return o
}

// SetCreateEpic adds the createEpic to the post API v3 epics params
func (o *PostAPIV3EpicsParams) SetCreateEpic(createEpic *models.CreateEpic) {
	o.CreateEpic = createEpic
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3EpicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateEpic != nil {
		if err := r.SetBodyParam(o.CreateEpic); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
