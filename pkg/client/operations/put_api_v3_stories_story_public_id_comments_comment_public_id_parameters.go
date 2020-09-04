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

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams creates a new PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams object
// with the default values initialized.
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams() *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithTimeout creates a new PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams{

		timeout: timeout,
	}
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithContext creates a new PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams{

		Context: ctx,
	}
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithHTTPClient creates a new PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParamsWithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	var ()
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams{
		HTTPClient: client,
	}
}

/*PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams contains all the parameters to send to the API endpoint
for the put API v3 stories story public ID comments comment public ID operation typically these are written to a http.Request
*/
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams struct {

	/*UpdateComment*/
	UpdateComment *models.UpdateComment
	/*CommentPublicID
	  The ID of the Comment

	*/
	CommentPublicID int64
	/*StoryPublicID
	  The ID of the Story that the Comment is in.

	*/
	StoryPublicID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithTimeout(timeout time.Duration) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithContext(ctx context.Context) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithHTTPClient(client *http.Client) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateComment adds the updateComment to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithUpdateComment(updateComment *models.UpdateComment) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetUpdateComment(updateComment)
	return o
}

// SetUpdateComment adds the updateComment to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetUpdateComment(updateComment *models.UpdateComment) {
	o.UpdateComment = updateComment
}

// WithCommentPublicID adds the commentPublicID to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithCommentPublicID(commentPublicID int64) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetCommentPublicID(commentPublicID)
	return o
}

// SetCommentPublicID adds the commentPublicId to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetCommentPublicID(commentPublicID int64) {
	o.CommentPublicID = commentPublicID
}

// WithStoryPublicID adds the storyPublicID to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WithStoryPublicID(storyPublicID int64) *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams {
	o.SetStoryPublicID(storyPublicID)
	return o
}

// SetStoryPublicID adds the storyPublicId to the put API v3 stories story public ID comments comment public ID params
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) SetStoryPublicID(storyPublicID int64) {
	o.StoryPublicID = storyPublicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param story-public-id
	if err := r.SetPathParam("story-public-id", swag.FormatInt64(o.StoryPublicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}