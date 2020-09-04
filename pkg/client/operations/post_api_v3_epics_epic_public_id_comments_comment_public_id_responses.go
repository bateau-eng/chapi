// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader is a Reader for the PostAPIV3EpicsEpicPublicIDCommentsCommentPublicID structure.
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated creates a PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated with default headers values
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated() *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated {
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated{}
}

/*PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated handles this case with default header values.

Resource
*/
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated struct {
	Payload *models.ThreadedComment
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] postApiV3EpicsEpicPublicIdCommentsCommentPublicIdCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated) GetPayload() *models.ThreadedComment {
	return o.Payload
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreadedComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest creates a PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest with default headers values
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest() *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest {
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest{}
}

/*PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest struct {
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] postApiV3EpicsEpicPublicIdCommentsCommentPublicIdBadRequest ", 400)
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound creates a PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound with default headers values
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound() *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound {
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound{}
}

/*PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound struct {
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] postApiV3EpicsEpicPublicIdCommentsCommentPublicIdNotFound ", 404)
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity creates a PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity with default headers values
func NewPostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity() *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity {
	return &PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity{}
}

/*PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity struct {
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] postApiV3EpicsEpicPublicIdCommentsCommentPublicIdUnprocessableEntity ", 422)
}

func (o *PostAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}