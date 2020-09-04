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

// PostAPIV3IterationsReader is a Reader for the PostAPIV3Iterations structure.
type PostAPIV3IterationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3IterationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3IterationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3IterationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3IterationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3IterationsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3IterationsCreated creates a PostAPIV3IterationsCreated with default headers values
func NewPostAPIV3IterationsCreated() *PostAPIV3IterationsCreated {
	return &PostAPIV3IterationsCreated{}
}

/*PostAPIV3IterationsCreated handles this case with default header values.

Resource
*/
type PostAPIV3IterationsCreated struct {
	Payload *models.Iteration
}

func (o *PostAPIV3IterationsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/iterations][%d] postApiV3IterationsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3IterationsCreated) GetPayload() *models.Iteration {
	return o.Payload
}

func (o *PostAPIV3IterationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Iteration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3IterationsBadRequest creates a PostAPIV3IterationsBadRequest with default headers values
func NewPostAPIV3IterationsBadRequest() *PostAPIV3IterationsBadRequest {
	return &PostAPIV3IterationsBadRequest{}
}

/*PostAPIV3IterationsBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3IterationsBadRequest struct {
}

func (o *PostAPIV3IterationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/iterations][%d] postApiV3IterationsBadRequest ", 400)
}

func (o *PostAPIV3IterationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3IterationsNotFound creates a PostAPIV3IterationsNotFound with default headers values
func NewPostAPIV3IterationsNotFound() *PostAPIV3IterationsNotFound {
	return &PostAPIV3IterationsNotFound{}
}

/*PostAPIV3IterationsNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3IterationsNotFound struct {
}

func (o *PostAPIV3IterationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/iterations][%d] postApiV3IterationsNotFound ", 404)
}

func (o *PostAPIV3IterationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3IterationsUnprocessableEntity creates a PostAPIV3IterationsUnprocessableEntity with default headers values
func NewPostAPIV3IterationsUnprocessableEntity() *PostAPIV3IterationsUnprocessableEntity {
	return &PostAPIV3IterationsUnprocessableEntity{}
}

/*PostAPIV3IterationsUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3IterationsUnprocessableEntity struct {
}

func (o *PostAPIV3IterationsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/iterations][%d] postApiV3IterationsUnprocessableEntity ", 422)
}

func (o *PostAPIV3IterationsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
