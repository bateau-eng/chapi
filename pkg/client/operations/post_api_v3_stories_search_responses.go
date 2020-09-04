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

// PostAPIV3StoriesSearchReader is a Reader for the PostAPIV3StoriesSearch structure.
type PostAPIV3StoriesSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3StoriesSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3StoriesSearchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3StoriesSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3StoriesSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3StoriesSearchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3StoriesSearchCreated creates a PostAPIV3StoriesSearchCreated with default headers values
func NewPostAPIV3StoriesSearchCreated() *PostAPIV3StoriesSearchCreated {
	return &PostAPIV3StoriesSearchCreated{}
}

/*PostAPIV3StoriesSearchCreated handles this case with default header values.

Resource
*/
type PostAPIV3StoriesSearchCreated struct {
	Payload []*models.StorySlim
}

func (o *PostAPIV3StoriesSearchCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/search][%d] postApiV3StoriesSearchCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3StoriesSearchCreated) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *PostAPIV3StoriesSearchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3StoriesSearchBadRequest creates a PostAPIV3StoriesSearchBadRequest with default headers values
func NewPostAPIV3StoriesSearchBadRequest() *PostAPIV3StoriesSearchBadRequest {
	return &PostAPIV3StoriesSearchBadRequest{}
}

/*PostAPIV3StoriesSearchBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3StoriesSearchBadRequest struct {
}

func (o *PostAPIV3StoriesSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/search][%d] postApiV3StoriesSearchBadRequest ", 400)
}

func (o *PostAPIV3StoriesSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesSearchNotFound creates a PostAPIV3StoriesSearchNotFound with default headers values
func NewPostAPIV3StoriesSearchNotFound() *PostAPIV3StoriesSearchNotFound {
	return &PostAPIV3StoriesSearchNotFound{}
}

/*PostAPIV3StoriesSearchNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3StoriesSearchNotFound struct {
}

func (o *PostAPIV3StoriesSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/search][%d] postApiV3StoriesSearchNotFound ", 404)
}

func (o *PostAPIV3StoriesSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesSearchUnprocessableEntity creates a PostAPIV3StoriesSearchUnprocessableEntity with default headers values
func NewPostAPIV3StoriesSearchUnprocessableEntity() *PostAPIV3StoriesSearchUnprocessableEntity {
	return &PostAPIV3StoriesSearchUnprocessableEntity{}
}

/*PostAPIV3StoriesSearchUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3StoriesSearchUnprocessableEntity struct {
}

func (o *PostAPIV3StoriesSearchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/search][%d] postApiV3StoriesSearchUnprocessableEntity ", 422)
}

func (o *PostAPIV3StoriesSearchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
