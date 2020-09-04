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

// PostAPIV3ProjectsReader is a Reader for the PostAPIV3Projects structure.
type PostAPIV3ProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3ProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3ProjectsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3ProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3ProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3ProjectsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3ProjectsCreated creates a PostAPIV3ProjectsCreated with default headers values
func NewPostAPIV3ProjectsCreated() *PostAPIV3ProjectsCreated {
	return &PostAPIV3ProjectsCreated{}
}

/*PostAPIV3ProjectsCreated handles this case with default header values.

Resource
*/
type PostAPIV3ProjectsCreated struct {
	Payload *models.Project
}

func (o *PostAPIV3ProjectsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] postApiV3ProjectsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3ProjectsCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *PostAPIV3ProjectsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3ProjectsBadRequest creates a PostAPIV3ProjectsBadRequest with default headers values
func NewPostAPIV3ProjectsBadRequest() *PostAPIV3ProjectsBadRequest {
	return &PostAPIV3ProjectsBadRequest{}
}

/*PostAPIV3ProjectsBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3ProjectsBadRequest struct {
}

func (o *PostAPIV3ProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] postApiV3ProjectsBadRequest ", 400)
}

func (o *PostAPIV3ProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3ProjectsNotFound creates a PostAPIV3ProjectsNotFound with default headers values
func NewPostAPIV3ProjectsNotFound() *PostAPIV3ProjectsNotFound {
	return &PostAPIV3ProjectsNotFound{}
}

/*PostAPIV3ProjectsNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3ProjectsNotFound struct {
}

func (o *PostAPIV3ProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] postApiV3ProjectsNotFound ", 404)
}

func (o *PostAPIV3ProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3ProjectsUnprocessableEntity creates a PostAPIV3ProjectsUnprocessableEntity with default headers values
func NewPostAPIV3ProjectsUnprocessableEntity() *PostAPIV3ProjectsUnprocessableEntity {
	return &PostAPIV3ProjectsUnprocessableEntity{}
}

/*PostAPIV3ProjectsUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3ProjectsUnprocessableEntity struct {
}

func (o *PostAPIV3ProjectsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/projects][%d] postApiV3ProjectsUnprocessableEntity ", 422)
}

func (o *PostAPIV3ProjectsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
