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

// NewPostAPIV3LabelsParams creates a new PostAPIV3LabelsParams object
// with the default values initialized.
func NewPostAPIV3LabelsParams() *PostAPIV3LabelsParams {
	var ()
	return &PostAPIV3LabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3LabelsParamsWithTimeout creates a new PostAPIV3LabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3LabelsParamsWithTimeout(timeout time.Duration) *PostAPIV3LabelsParams {
	var ()
	return &PostAPIV3LabelsParams{

		timeout: timeout,
	}
}

// NewPostAPIV3LabelsParamsWithContext creates a new PostAPIV3LabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3LabelsParamsWithContext(ctx context.Context) *PostAPIV3LabelsParams {
	var ()
	return &PostAPIV3LabelsParams{

		Context: ctx,
	}
}

// NewPostAPIV3LabelsParamsWithHTTPClient creates a new PostAPIV3LabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3LabelsParamsWithHTTPClient(client *http.Client) *PostAPIV3LabelsParams {
	var ()
	return &PostAPIV3LabelsParams{
		HTTPClient: client,
	}
}

/*PostAPIV3LabelsParams contains all the parameters to send to the API endpoint
for the post API v3 labels operation typically these are written to a http.Request
*/
type PostAPIV3LabelsParams struct {

	/*CreateLabelParams
	  Request parameters for creating a Label on a Clubhouse story.

	*/
	CreateLabelParams *models.CreateLabelParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 labels params
func (o *PostAPIV3LabelsParams) WithTimeout(timeout time.Duration) *PostAPIV3LabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 labels params
func (o *PostAPIV3LabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 labels params
func (o *PostAPIV3LabelsParams) WithContext(ctx context.Context) *PostAPIV3LabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 labels params
func (o *PostAPIV3LabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 labels params
func (o *PostAPIV3LabelsParams) WithHTTPClient(client *http.Client) *PostAPIV3LabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 labels params
func (o *PostAPIV3LabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateLabelParams adds the createLabelParams to the post API v3 labels params
func (o *PostAPIV3LabelsParams) WithCreateLabelParams(createLabelParams *models.CreateLabelParams) *PostAPIV3LabelsParams {
	o.SetCreateLabelParams(createLabelParams)
	return o
}

// SetCreateLabelParams adds the createLabelParams to the post API v3 labels params
func (o *PostAPIV3LabelsParams) SetCreateLabelParams(createLabelParams *models.CreateLabelParams) {
	o.CreateLabelParams = createLabelParams
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3LabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateLabelParams != nil {
		if err := r.SetBodyParam(o.CreateLabelParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
