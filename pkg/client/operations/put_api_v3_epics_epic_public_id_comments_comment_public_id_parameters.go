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

// NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams creates a new PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized.
func NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams() *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout creates a new PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext creates a new PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext(ctx context.Context) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient creates a new PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 epics epic public ID comments comment public ID operation typically these are written to a http.Request
*/
type PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams struct {

	/*UpdateComment*/
	UpdateComment *models.UpdateComment
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

// WithTimeout adds the timeout to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithContext(ctx context.Context) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateComment adds the updateComment to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithUpdateComment(updateComment *models.UpdateComment) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetUpdateComment(updateComment)
	return o
}

// SetUpdateComment adds the updateComment to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetUpdateComment(updateComment *models.UpdateComment) {
	o.UpdateComment = updateComment
}

// WithCommentPublicID adds the commentPublicID to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithCommentPublicID(commentPublicID int64) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithEpicPublicID adds the epicPublicID to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithEpicPublicID(epicPublicID int64) *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the put API v3 epics epic public ID comments comment public ID params
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateComment != nil {
		if err := r.SetBodyParam(o.UpdateComment); err != nil {
			return err
		}
	}

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
