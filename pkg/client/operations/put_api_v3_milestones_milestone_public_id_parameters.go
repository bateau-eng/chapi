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

	"github.com/bateau-eng/chapi/pkg/models"
)

// NewPutAPIV3MilestonesMilestonePublicIDParams creates a new PutAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized.
func NewPutAPIV3MilestonesMilestonePublicIDParams() *PutAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &PutAPIV3MilestonesMilestonePublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3MilestonesMilestonePublicIDParamsWithTimeout creates a new PutAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3MilestonesMilestonePublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &PutAPIV3MilestonesMilestonePublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3MilestonesMilestonePublicIDParamsWithContext creates a new PutAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3MilestonesMilestonePublicIDParamsWithContext(ctx context.Context) *PutAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &PutAPIV3MilestonesMilestonePublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3MilestonesMilestonePublicIDParamsWithHTTPClient creates a new PutAPIV3MilestonesMilestonePublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3MilestonesMilestonePublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3MilestonesMilestonePublicIDParams {
	var ()
	return &PutAPIV3MilestonesMilestonePublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3MilestonesMilestonePublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 milestones milestone public ID operation typically these are written to a http.Request
*/
type PutAPIV3MilestonesMilestonePublicIDParams struct {

	/*UpdateMilestone*/
	UpdateMilestone *models.UpdateMilestone
	/*MilestonePublicID
	  The ID of the Milestone.

	*/
	MilestonePublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3MilestonesMilestonePublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WithContext(ctx context.Context) *PutAPIV3MilestonesMilestonePublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3MilestonesMilestonePublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateMilestone adds the updateMilestone to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WithUpdateMilestone(updateMilestone *models.UpdateMilestone) *PutAPIV3MilestonesMilestonePublicIDParams {
	o.SetUpdateMilestone(updateMilestone)
	return o
}

// SetUpdateMilestone adds the updateMilestone to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) SetUpdateMilestone(updateMilestone *models.UpdateMilestone) {
	o.UpdateMilestone = updateMilestone
}

// WithMilestonePublicID adds the milestonePublicID to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WithMilestonePublicID(milestonePublicID int64) *PutAPIV3MilestonesMilestonePublicIDParams {
	o.SetMilestonePublicID(milestonePublicID)
	return o
}

// SetMilestonePublicID adds the milestonePublicId to the put API v3 milestones milestone public ID params
func (o *PutAPIV3MilestonesMilestonePublicIDParams) SetMilestonePublicID(milestonePublicID int64) {
	o.MilestonePublicID = milestonePublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3MilestonesMilestonePublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateMilestone != nil {
		if err := r.SetBodyParam(o.UpdateMilestone); err != nil {
			return err
		}
	}

	// path param milestone-public-id
	if err := r.SetPathParam("milestone-public-id", swag.FormatInt64(o.MilestonePublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
