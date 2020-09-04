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

// NewGetAPIV3TeamsTeamPublicIDParams creates a new GetAPIV3TeamsTeamPublicIDParams object
// with the default values initialized.
func NewGetAPIV3TeamsTeamPublicIDParams() *GetAPIV3TeamsTeamPublicIDParams {
	var ()
	return &GetAPIV3TeamsTeamPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3TeamsTeamPublicIDParamsWithTimeout creates a new GetAPIV3TeamsTeamPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3TeamsTeamPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3TeamsTeamPublicIDParams {
	var ()
	return &GetAPIV3TeamsTeamPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3TeamsTeamPublicIDParamsWithContext creates a new GetAPIV3TeamsTeamPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3TeamsTeamPublicIDParamsWithContext(ctx context.Context) *GetAPIV3TeamsTeamPublicIDParams {
	var ()
	return &GetAPIV3TeamsTeamPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3TeamsTeamPublicIDParamsWithHTTPClient creates a new GetAPIV3TeamsTeamPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3TeamsTeamPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3TeamsTeamPublicIDParams {
	var ()
	return &GetAPIV3TeamsTeamPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3TeamsTeamPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 teams team public ID operation typically these are written to a http.Request
*/
type GetAPIV3TeamsTeamPublicIDParams struct {

	/*TeamPublicID
	  The ID of the team.

	*/
	TeamPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3TeamsTeamPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) WithContext(ctx context.Context) *GetAPIV3TeamsTeamPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3TeamsTeamPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamPublicID adds the teamPublicID to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) WithTeamPublicID(teamPublicID int64) *GetAPIV3TeamsTeamPublicIDParams {
	o.SetTeamPublicID(teamPublicID)
	return o
}

// SetTeamPublicID adds the teamPublicId to the get API v3 teams team public ID params
func (o *GetAPIV3TeamsTeamPublicIDParams) SetTeamPublicID(teamPublicID int64) {
	o.TeamPublicID = teamPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3TeamsTeamPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team-public-id
	if err := r.SetPathParam("team-public-id", swag.FormatInt64(o.TeamPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
