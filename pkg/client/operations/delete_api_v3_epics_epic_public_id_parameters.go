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

// NewDeleteAPIV3EpicsEpicPublicIDParams creates a new DeleteAPIV3EpicsEpicPublicIDParams object
// with the default values initialized.
func NewDeleteAPIV3EpicsEpicPublicIDParams() *DeleteAPIV3EpicsEpicPublicIDParams {
	var ()
	return &DeleteAPIV3EpicsEpicPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3EpicsEpicPublicIDParamsWithTimeout creates a new DeleteAPIV3EpicsEpicPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3EpicsEpicPublicIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV3EpicsEpicPublicIDParams {
	var ()
	return &DeleteAPIV3EpicsEpicPublicIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3EpicsEpicPublicIDParamsWithContext creates a new DeleteAPIV3EpicsEpicPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3EpicsEpicPublicIDParamsWithContext(ctx context.Context) *DeleteAPIV3EpicsEpicPublicIDParams {
	var ()
	return &DeleteAPIV3EpicsEpicPublicIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3EpicsEpicPublicIDParamsWithHTTPClient creates a new DeleteAPIV3EpicsEpicPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3EpicsEpicPublicIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV3EpicsEpicPublicIDParams {
	var ()
	return &DeleteAPIV3EpicsEpicPublicIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3EpicsEpicPublicIDParams contains all the parameters to send to the API endpoint
for the delete API v3 epics epic public ID operation typically these are written to a http.Request
*/
type DeleteAPIV3EpicsEpicPublicIDParams struct {

	/*EpicPublicID
	  The unique ID of the Epic.

	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV3EpicsEpicPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) WithContext(ctx context.Context) *DeleteAPIV3EpicsEpicPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV3EpicsEpicPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEpicPublicID adds the epicPublicID to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) WithEpicPublicID(epicPublicID int64) *DeleteAPIV3EpicsEpicPublicIDParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the delete API v3 epics epic public ID params
func (o *DeleteAPIV3EpicsEpicPublicIDParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3EpicsEpicPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param epic-public-id
	if err := r.SetPathParam("epic-public-id", swag.FormatInt64(o.EpicPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
