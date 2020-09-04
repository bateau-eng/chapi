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

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams creates a new GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized.
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams() *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout creates a new GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout(timeout time.Duration) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: timeout,
	}
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext creates a new GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext(ctx context.Context) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		Context: ctx,
	}
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient creates a new GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient(client *http.Client) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{
		HTTPClient: client,
	}
}

/*GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams contains all the parameters to send to the API endpoint
for the get API v3 epics epic public ID comments comment public ID operation typically these are written to a http.Request
*/
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams struct {

	/*CommentPublicID
	  The ID of the Comment.

	*/
	CommentPublicID int64
	/*EpicPublicID
	  The ID of the associated Epic.

	*/
	EpicPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithTimeout(timeout time.Duration) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithContext(ctx context.Context) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithHTTPClient(client *http.Client) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentPublicID adds the commentPublicID to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithCommentPublicID(commentPublicID int64) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithEpicPublicID adds the epicPublicID to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithEpicPublicID(epicPublicID int64) *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the get API v3 epics epic public ID comments comment public ID params
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment-public-id
	if err := r.SetPathParam("comment-public-id", swag.FormatInt64(o.CommentPublicID)); err != nil {
		return err
	}

	// path param epic-public-id
	if err := r.SetPathParam("epic-public-id", swag.FormatInt64(o.EpicPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
