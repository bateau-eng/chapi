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

// NewPostAPIV3GroupsParams creates a new PostAPIV3GroupsParams object
// with the default values initialized.
func NewPostAPIV3GroupsParams() *PostAPIV3GroupsParams {
	var ()
	return &PostAPIV3GroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3GroupsParamsWithTimeout creates a new PostAPIV3GroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3GroupsParamsWithTimeout(timeout time.Duration) *PostAPIV3GroupsParams {
	var ()
	return &PostAPIV3GroupsParams{

		timeout: timeout,
	}
}

// NewPostAPIV3GroupsParamsWithContext creates a new PostAPIV3GroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3GroupsParamsWithContext(ctx context.Context) *PostAPIV3GroupsParams {
	var ()
	return &PostAPIV3GroupsParams{

		Context: ctx,
	}
}

// NewPostAPIV3GroupsParamsWithHTTPClient creates a new PostAPIV3GroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3GroupsParamsWithHTTPClient(client *http.Client) *PostAPIV3GroupsParams {
	var ()
	return &PostAPIV3GroupsParams{
		HTTPClient: client,
	}
}

/*PostAPIV3GroupsParams contains all the parameters to send to the API endpoint
for the post API v3 groups operation typically these are written to a http.Request
*/
type PostAPIV3GroupsParams struct {

	/*CreateGroup*/
	CreateGroup *models.CreateGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API v3 groups params
func (o *PostAPIV3GroupsParams) WithTimeout(timeout time.Duration) *PostAPIV3GroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 groups params
func (o *PostAPIV3GroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 groups params
func (o *PostAPIV3GroupsParams) WithContext(ctx context.Context) *PostAPIV3GroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 groups params
func (o *PostAPIV3GroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 groups params
func (o *PostAPIV3GroupsParams) WithHTTPClient(client *http.Client) *PostAPIV3GroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 groups params
func (o *PostAPIV3GroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateGroup adds the createGroup to the post API v3 groups params
func (o *PostAPIV3GroupsParams) WithCreateGroup(createGroup *models.CreateGroup) *PostAPIV3GroupsParams {
	o.SetCreateGroup(createGroup)
	return o
}

// SetCreateGroup adds the createGroup to the post API v3 groups params
func (o *PostAPIV3GroupsParams) SetCreateGroup(createGroup *models.CreateGroup) {
	o.CreateGroup = createGroup
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3GroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateGroup != nil {
		if err := r.SetBodyParam(o.CreateGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}