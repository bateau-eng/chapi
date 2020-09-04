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

// NewGetAPIV3MilestonesMilestonePublicIDParams creates a new GetAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized.
func NewGetAPIV3MilestonesMilestonePublicIDParams() *GetAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &GetAPIV3MilestonesMilestonePublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MilestonesMilestonePublicIDParamsWithTimeout creates a new GetAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MilestonesMilestonePublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &GetAPIV3MilestonesMilestonePublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MilestonesMilestonePublicIDParamsWithContext creates a new GetAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MilestonesMilestonePublicIDParamsWithContext(ctx context.Context) *GetAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &GetAPIV3MilestonesMilestonePublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3MilestonesMilestonePublicIDParamsWithHTTPClient creates a new GetAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MilestonesMilestonePublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &GetAPIV3MilestonesMilestonePublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MilestonesMilestonePublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 milestones milestone public ID operation typically these are written to a http.Request
*/
type GetAPIV3MilestonesMilestonePublicIDParams struct {

	/*MilestonePublicID
	  The ID of the Milestone.

	*/
	MilestonePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3MilestonesMilestonePublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) WithContext(ctx context.Context) *GetAPIV3MilestonesMilestonePublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3MilestonesMilestonePublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestonePublicID adds the milestonePublicID to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) WithMilestonePublicID(milestonePublicID int64) *GetAPIV3MilestonesMilestonePublicIDParams {
	o.SetMilestonePublicID(milestonePublicID)
	return o
}

// SetMilestonePublicID adds the milestonePublicId to the get API v3 milestones milestone public ID params
func (o *GetAPIV3MilestonesMilestonePublicIDParams) SetMilestonePublicID(milestonePublicID int64) {
	o.MilestonePublicID = milestonePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MilestonesMilestonePublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestone-public-id
	if err := r.SetPathParam("milestone-public-id", swag.FormatInt64(o.MilestonePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
