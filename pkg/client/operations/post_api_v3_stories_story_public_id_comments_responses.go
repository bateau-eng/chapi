// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bateau-eng/chapi/pkg/models"
)

// PostAPIV3StoriesStoryPublicIDCommentsReader is a Reader for the PostAPIV3StoriesStoryPublicIDComments structure.
type PostAPIV3StoriesStoryPublicIDCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3StoriesStoryPublicIDCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3StoriesStoryPublicIDCommentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3StoriesStoryPublicIDCommentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3StoriesStoryPublicIDCommentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3StoriesStoryPublicIDCommentsCreated creates a PostAPIV3StoriesStoryPublicIDCommentsCreated with default headers values
func NewPostAPIV3StoriesStoryPublicIDCommentsCreated() *PostAPIV3StoriesStoryPublicIDCommentsCreated {
	return &PostAPIV3StoriesStoryPublicIDCommentsCreated{}
}

/*PostAPIV3StoriesStoryPublicIDCommentsCreated handles this case with default header values.

Resource
*/
type PostAPIV3StoriesStoryPublicIDCommentsCreated struct {
	Payload *models.Comment
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments][%d] postApiV3StoriesStoryPublicIdCommentsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsCreated) GetPayload() *models.Comment {
	return o.Payload
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3StoriesStoryPublicIDCommentsBadRequest creates a PostAPIV3StoriesStoryPublicIDCommentsBadRequest with default headers values
func NewPostAPIV3StoriesStoryPublicIDCommentsBadRequest() *PostAPIV3StoriesStoryPublicIDCommentsBadRequest {
	return &PostAPIV3StoriesStoryPublicIDCommentsBadRequest{}
}

/*PostAPIV3StoriesStoryPublicIDCommentsBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3StoriesStoryPublicIDCommentsBadRequest struct {
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments][%d] postApiV3StoriesStoryPublicIdCommentsBadRequest ", 400)
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesStoryPublicIDCommentsNotFound creates a PostAPIV3StoriesStoryPublicIDCommentsNotFound with default headers values
func NewPostAPIV3StoriesStoryPublicIDCommentsNotFound() *PostAPIV3StoriesStoryPublicIDCommentsNotFound {
	return &PostAPIV3StoriesStoryPublicIDCommentsNotFound{}
}

/*PostAPIV3StoriesStoryPublicIDCommentsNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3StoriesStoryPublicIDCommentsNotFound struct {
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments][%d] postApiV3StoriesStoryPublicIdCommentsNotFound ", 404)
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity creates a PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity with default headers values
func NewPostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity() *PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity {
	return &PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity{}
}

/*PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity struct {
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/{story-public-id}/comments][%d] postApiV3StoriesStoryPublicIdCommentsUnprocessableEntity ", 422)
}

func (o *PostAPIV3StoriesStoryPublicIDCommentsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
