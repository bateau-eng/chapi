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

// PostAPIV3StoriesBulkReader is a Reader for the PostAPIV3StoriesBulk structure.
type PostAPIV3StoriesBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3StoriesBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3StoriesBulkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3StoriesBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3StoriesBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3StoriesBulkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3StoriesBulkCreated creates a PostAPIV3StoriesBulkCreated with default headers values
func NewPostAPIV3StoriesBulkCreated() *PostAPIV3StoriesBulkCreated {
	return &PostAPIV3StoriesBulkCreated{}
}

/*PostAPIV3StoriesBulkCreated handles this case with default header values.

Resource
*/
type PostAPIV3StoriesBulkCreated struct {
	Payload []*models.StorySlim
}

func (o *PostAPIV3StoriesBulkCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/bulk][%d] postApiV3StoriesBulkCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3StoriesBulkCreated) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *PostAPIV3StoriesBulkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3StoriesBulkBadRequest creates a PostAPIV3StoriesBulkBadRequest with default headers values
func NewPostAPIV3StoriesBulkBadRequest() *PostAPIV3StoriesBulkBadRequest {
	return &PostAPIV3StoriesBulkBadRequest{}
}

/*PostAPIV3StoriesBulkBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3StoriesBulkBadRequest struct {
}

func (o *PostAPIV3StoriesBulkBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/bulk][%d] postApiV3StoriesBulkBadRequest ", 400)
}

func (o *PostAPIV3StoriesBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesBulkNotFound creates a PostAPIV3StoriesBulkNotFound with default headers values
func NewPostAPIV3StoriesBulkNotFound() *PostAPIV3StoriesBulkNotFound {
	return &PostAPIV3StoriesBulkNotFound{}
}

/*PostAPIV3StoriesBulkNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3StoriesBulkNotFound struct {
}

func (o *PostAPIV3StoriesBulkNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/bulk][%d] postApiV3StoriesBulkNotFound ", 404)
}

func (o *PostAPIV3StoriesBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3StoriesBulkUnprocessableEntity creates a PostAPIV3StoriesBulkUnprocessableEntity with default headers values
func NewPostAPIV3StoriesBulkUnprocessableEntity() *PostAPIV3StoriesBulkUnprocessableEntity {
	return &PostAPIV3StoriesBulkUnprocessableEntity{}
}

/*PostAPIV3StoriesBulkUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3StoriesBulkUnprocessableEntity struct {
}

func (o *PostAPIV3StoriesBulkUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/stories/bulk][%d] postApiV3StoriesBulkUnprocessableEntity ", 422)
}

func (o *PostAPIV3StoriesBulkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
