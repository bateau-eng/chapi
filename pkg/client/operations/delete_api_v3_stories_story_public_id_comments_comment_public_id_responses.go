// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader is a Reader for the DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicID structure.
type DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent creates a DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent with default headers values
func NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent() *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent {
	return &DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent{}
}

/*DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent handles this case with default header values.

No Content
*/
type DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent struct {
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] deleteApiV3StoriesStoryPublicIdCommentsCommentPublicIdNoContent ", 204)
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest creates a DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest with default headers values
func NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest() *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest {
	return &DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest{}
}

/*DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest struct {
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] deleteApiV3StoriesStoryPublicIdCommentsCommentPublicIdBadRequest ", 400)
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound creates a DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound with default headers values
func NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound() *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound {
	return &DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound{}
}

/*DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound struct {
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] deleteApiV3StoriesStoryPublicIdCommentsCommentPublicIdNotFound ", 404)
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity creates a DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity with default headers values
func NewDeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity() *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity {
	return &DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity{}
}

/*DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity struct {
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] deleteApiV3StoriesStoryPublicIdCommentsCommentPublicIdUnprocessableEntity ", 422)
}

func (o *DeleteAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}