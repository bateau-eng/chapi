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

// NewPutAPIV3GroupsGroupPublicIDParams creates a new PutAPIV3GroupsGroupPublicIDParams object
// with the default values initialized.
func NewPutAPIV3GroupsGroupPublicIDParams() *PutAPIV3GroupsGroupPublicIDParams {
	var ()
	return &PutAPIV3GroupsGroupPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3GroupsGroupPublicIDParamsWithTimeout creates a new PutAPIV3GroupsGroupPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3GroupsGroupPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3GroupsGroupPublicIDParams {
	var ()
	return &PutAPIV3GroupsGroupPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3GroupsGroupPublicIDParamsWithContext creates a new PutAPIV3GroupsGroupPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3GroupsGroupPublicIDParamsWithContext(ctx context.Context) *PutAPIV3GroupsGroupPublicIDParams {
	var ()
	return &PutAPIV3GroupsGroupPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3GroupsGroupPublicIDParamsWithHTTPClient creates a new PutAPIV3GroupsGroupPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3GroupsGroupPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3GroupsGroupPublicIDParams {
	var ()
	return &PutAPIV3GroupsGroupPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3GroupsGroupPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 groups group public ID operation typically these are written to a http.Request
*/
type PutAPIV3GroupsGroupPublicIDParams struct {

	/*UpdateGroup*/
	UpdateGroup *models.UpdateGroup
	/*GroupPublicID
	  The unique ID of the Group.

	*/
	GroupPublicID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3GroupsGroupPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) WithContext(ctx context.Context) *PutAPIV3GroupsGroupPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3GroupsGroupPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateGroup adds the updateGroup to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) WithUpdateGroup(updateGroup *models.UpdateGroup) *PutAPIV3GroupsGroupPublicIDParams {
	o.SetUpdateGroup(updateGroup)
	return o
}

// SetUpdateGroup adds the updateGroup to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) SetUpdateGroup(updateGroup *models.UpdateGroup) {
	o.UpdateGroup = updateGroup
}

// WithGroupPublicID adds the groupPublicID to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) WithGroupPublicID(groupPublicID strfmt.UUID) *PutAPIV3GroupsGroupPublicIDParams {
	o.SetGroupPublicID(groupPublicID)
	return o
}

// SetGroupPublicID adds the groupPublicId to the put API v3 groups group public ID params
func (o *PutAPIV3GroupsGroupPublicIDParams) SetGroupPublicID(groupPublicID strfmt.UUID) {
	o.GroupPublicID = groupPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3GroupsGroupPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateGroup != nil {
		if err := r.SetBodyParam(o.UpdateGroup); err != nil {
			return err
		}
	}

	// path param group-public-id
	if err := r.SetPathParam("group-public-id", o.GroupPublicID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
