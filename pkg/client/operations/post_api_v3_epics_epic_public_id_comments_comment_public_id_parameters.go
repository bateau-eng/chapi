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

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams creates a new PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized.
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams() *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout creates a new PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithTimeout(timeout time.Duration) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		timeout: timeout,
	}
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext creates a new PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithContext(ctx context.Context) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{

		Context: ctx,
	}
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient creates a new PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParamsWithHTTPClient(client *http.Client) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams{
		HTTPClient: client,
	}
}

/*PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams contains all the parameters to send to the API endpoint
for the post API v3 epics epic public ID comments comment public ID operation typically these are written to a http.Request
*/
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams struct {

	/*CreateCommentComment*/
	CreateCommentComment *models.CreateCommentComment
	/*CommentPublicID
	  The ID of the parent Epic Comment.

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

// WithTimeout adds the timeout to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithTimeout(timeout time.Duration) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithContext(ctx context.Context) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithHTTPClient(client *http.Client) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateCommentComment adds the createCommentComment to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithCreateCommentComment(createCommentComment *models.CreateCommentComment) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetCreateCommentComment(createCommentComment)
	return o
}

// SetCreateCommentComment adds the createCommentComment to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetCreateCommentComment(createCommentComment *models.CreateCommentComment) {
	o.CreateCommentComment = createCommentComment
}

// WithCommentPublicID adds the commentPublicID to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithCommentPublicID(commentPublicID int64) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithEpicPublicID adds the epicPublicID to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WithEpicPublicID(epicPublicID int64) *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams {
	o.SetEpicPublicID(epicPublicID)
	return o
}

// SetEpicPublicID adds the epicPublicId to the post API v3 epics epic public ID comments comment public ID params
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) SetEpicPublicID(epicPublicID int64) {
	o.EpicPublicID = epicPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateCommentComment != nil {
		if err := r.SetBodyParam(o.CreateCommentComment); err != nil {
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
