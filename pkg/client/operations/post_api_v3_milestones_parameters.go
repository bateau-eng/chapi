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

// NewPostAPIV3MilestonesParams creates a new PostAPIV3MilestonesParams object
// with the default values initialized.
func NewPostAPIV3MilestonesParams() *PostAPIV3MilestonesParams {
	var ()
	return &PostAPIV3MilestonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3MilestonesParamsWithTimeout creates a new PostAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3MilestonesParamsWithTimeout(timeout time.Duration) *PostAPIV3MilestonesParams {
	var ()
	return &PostAPIV3MilestonesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3MilestonesParamsWithContext creates a new PostAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3MilestonesParamsWithContext(ctx context.Context) *PostAPIV3MilestonesParams {
	var ()
	return &PostAPIV3MilestonesParams{

		Context: ctx,
	}
}

// NewPostAPIV3MilestonesParamsWithHTTPClient creates a new PostAPIV3MilestonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3MilestonesParamsWithHTTPClient(client *http.Client) *PostAPIV3MilestonesParams {
	var ()
	return &PostAPIV3MilestonesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3MilestonesParams contains all the parameters to send to the API endpoint
for the post API v3 milestones operation typically these are written to a http.Request
*/
type PostAPIV3MilestonesParams struct {

	/*CreateMilestone*/
	CreateMilestone *models.CreateMilestone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) WithTimeout(timeout time.Duration) *PostAPIV3MilestonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) WithContext(ctx context.Context) *PostAPIV3MilestonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) WithHTTPClient(client *http.Client) *PostAPIV3MilestonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateMilestone adds the createMilestone to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) WithCreateMilestone(createMilestone *models.CreateMilestone) *PostAPIV3MilestonesParams {
	o.SetCreateMilestone(createMilestone)
	return o
}

// SetCreateMilestone adds the createMilestone to the post API v3 milestones params
func (o *PostAPIV3MilestonesParams) SetCreateMilestone(createMilestone *models.CreateMilestone) {
	o.CreateMilestone = createMilestone
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3MilestonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateMilestone != nil {
		if err := r.SetBodyParam(o.CreateMilestone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
