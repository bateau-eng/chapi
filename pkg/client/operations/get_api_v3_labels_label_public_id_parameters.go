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

// NewGetAPIV3LabelsLabelPublicIDParams creates a new GetAPIV3LabelsLabelPublicIDParams object
// with the default values initialized.
func NewGetAPIV3LabelsLabelPublicIDParams() *GetAPIV3LabelsLabelPublicIDParams {
	var ()
	return &GetAPIV3LabelsLabelPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3LabelsLabelPublicIDParamsWithTimeout creates a new GetAPIV3LabelsLabelPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3LabelsLabelPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3LabelsLabelPublicIDParams {
	var ()
	return &GetAPIV3LabelsLabelPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3LabelsLabelPublicIDParamsWithContext creates a new GetAPIV3LabelsLabelPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3LabelsLabelPublicIDParamsWithContext(ctx context.Context) *GetAPIV3LabelsLabelPublicIDParams {
	var ()
	return &GetAPIV3LabelsLabelPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3LabelsLabelPublicIDParamsWithHTTPClient creates a new GetAPIV3LabelsLabelPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3LabelsLabelPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3LabelsLabelPublicIDParams {
	var ()
	return &GetAPIV3LabelsLabelPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3LabelsLabelPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 labels label public ID operation typically these are written to a http.Request
*/
type GetAPIV3LabelsLabelPublicIDParams struct {

	/*LabelPublicID
	  The unique ID of the Label.

	*/
	LabelPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3LabelsLabelPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) WithContext(ctx context.Context) *GetAPIV3LabelsLabelPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3LabelsLabelPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabelPublicID adds the labelPublicID to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) WithLabelPublicID(labelPublicID int64) *GetAPIV3LabelsLabelPublicIDParams {
	o.SetLabelPublicID(labelPublicID)
	return o
}

// SetLabelPublicID adds the labelPublicId to the get API v3 labels label public ID params
func (o *GetAPIV3LabelsLabelPublicIDParams) SetLabelPublicID(labelPublicID int64) {
	o.LabelPublicID = labelPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3LabelsLabelPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param label-public-id
	if err := r.SetPathParam("label-public-id", swag.FormatInt64(o.LabelPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}