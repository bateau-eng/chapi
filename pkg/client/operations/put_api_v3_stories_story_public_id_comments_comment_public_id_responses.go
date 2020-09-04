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

// PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader is a Reader for the PutAPIV3StoriesStoryPublicIDCommentsCommentPublicID structure.
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK creates a PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK with default headers values
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK() *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK {
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK{}
}

/*PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK handles this case with default header values.

Resource
*/
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK struct {
	Payload *models.Comment
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] putApiV3StoriesStoryPublicIdCommentsCommentPublicIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK) GetPayload() *models.Comment {
	return o.Payload
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest creates a PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest with default headers values
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest() *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest {
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest{}
}

/*PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest struct {
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] putApiV3StoriesStoryPublicIdCommentsCommentPublicIdBadRequest ", 400)
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound creates a PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound with default headers values
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound() *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound {
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound{}
}

/*PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound struct {
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] putApiV3StoriesStoryPublicIdCommentsCommentPublicIdNotFound ", 404)
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity creates a PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity with default headers values
func NewPutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity() *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity {
	return &PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity{}
}

/*PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity struct {
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}/comments/{comment-public-id}][%d] putApiV3StoriesStoryPublicIdCommentsCommentPublicIdUnprocessableEntity ", 422)
}

func (o *PutAPIV3StoriesStoryPublicIDCommentsCommentPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}