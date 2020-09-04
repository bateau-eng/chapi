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

// NewGetAPIV3EpicsParams creates a new GetAPIV3EpicsParams object
// with the default values initialized.
func NewGetAPIV3EpicsParams() *GetAPIV3EpicsParams {
	var ()
	return &GetAPIV3EpicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3EpicsParamsWithTimeout creates a new GetAPIV3EpicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3EpicsParamsWithTimeout(timeout time.Duration) *GetAPIV3EpicsParams {
	var ()
	return &GetAPIV3EpicsParams{

		timeout: timeout,
	}
}

// NewGetAPIV3EpicsParamsWithContext creates a new GetAPIV3EpicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3EpicsParamsWithContext(ctx context.Context) *GetAPIV3EpicsParams {
	var ()
	return &GetAPIV3EpicsParams{

		Context: ctx,
	}
}

// NewGetAPIV3EpicsParamsWithHTTPClient creates a new GetAPIV3EpicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3EpicsParamsWithHTTPClient(client *http.Client) *GetAPIV3EpicsParams {
	var ()
	return &GetAPIV3EpicsParams{
		HTTPClient: client,
	}
}

/*GetAPIV3EpicsParams contains all the parameters to send to the API endpoint
for the get API v3 epics operation typically these are written to a http.Request
*/
type GetAPIV3EpicsParams struct {

	/*ListEpics*/
	ListEpics *models.ListEpics

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 epics params
func (o *GetAPIV3EpicsParams) WithTimeout(timeout time.Duration) *GetAPIV3EpicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 epics params
func (o *GetAPIV3EpicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 epics params
func (o *GetAPIV3EpicsParams) WithContext(ctx context.Context) *GetAPIV3EpicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 epics params
func (o *GetAPIV3EpicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 epics params
func (o *GetAPIV3EpicsParams) WithHTTPClient(client *http.Client) *GetAPIV3EpicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 epics params
func (o *GetAPIV3EpicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListEpics adds the listEpics to the get API v3 epics params
func (o *GetAPIV3EpicsParams) WithListEpics(listEpics *models.ListEpics) *GetAPIV3EpicsParams {
	o.SetListEpics(listEpics)
	return o
}

// SetListEpics adds the listEpics to the get API v3 epics params
func (o *GetAPIV3EpicsParams) SetListEpics(listEpics *models.ListEpics) {
	o.ListEpics = listEpics
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3EpicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ListEpics != nil {
		if err := r.SetBodyParam(o.ListEpics); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
