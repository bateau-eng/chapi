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

// PutAPIV3StoriesStoryPublicIDReader is a Reader for the PutAPIV3StoriesStoryPublicID structure.
type PutAPIV3StoriesStoryPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3StoriesStoryPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV3StoriesStoryPublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3StoriesStoryPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3StoriesStoryPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3StoriesStoryPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3StoriesStoryPublicIDOK creates a PutAPIV3StoriesStoryPublicIDOK with default headers values
func NewPutAPIV3StoriesStoryPublicIDOK() *PutAPIV3StoriesStoryPublicIDOK {
	return &PutAPIV3StoriesStoryPublicIDOK{}
}

/*PutAPIV3StoriesStoryPublicIDOK handles this case with default header values.

Resource
*/
type PutAPIV3StoriesStoryPublicIDOK struct {
	Payload *models.Story
}

func (o *PutAPIV3StoriesStoryPublicIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] putApiV3StoriesStoryPublicIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3StoriesStoryPublicIDOK) GetPayload() *models.Story {
	return o.Payload
}

func (o *PutAPIV3StoriesStoryPublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Story)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDBadRequest creates a PutAPIV3StoriesStoryPublicIDBadRequest with default headers values
func NewPutAPIV3StoriesStoryPublicIDBadRequest() *PutAPIV3StoriesStoryPublicIDBadRequest {
	return &PutAPIV3StoriesStoryPublicIDBadRequest{}
}

/*PutAPIV3StoriesStoryPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3StoriesStoryPublicIDBadRequest struct {
}

func (o *PutAPIV3StoriesStoryPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] putApiV3StoriesStoryPublicIdBadRequest ", 400)
}

func (o *PutAPIV3StoriesStoryPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDNotFound creates a PutAPIV3StoriesStoryPublicIDNotFound with default headers values
func NewPutAPIV3StoriesStoryPublicIDNotFound() *PutAPIV3StoriesStoryPublicIDNotFound {
	return &PutAPIV3StoriesStoryPublicIDNotFound{}
}

/*PutAPIV3StoriesStoryPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3StoriesStoryPublicIDNotFound struct {
}

func (o *PutAPIV3StoriesStoryPublicIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] putApiV3StoriesStoryPublicIdNotFound ", 404)
}

func (o *PutAPIV3StoriesStoryPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesStoryPublicIDUnprocessableEntity creates a PutAPIV3StoriesStoryPublicIDUnprocessableEntity with default headers values
func NewPutAPIV3StoriesStoryPublicIDUnprocessableEntity() *PutAPIV3StoriesStoryPublicIDUnprocessableEntity {
	return &PutAPIV3StoriesStoryPublicIDUnprocessableEntity{}
}

/*PutAPIV3StoriesStoryPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3StoriesStoryPublicIDUnprocessableEntity struct {
}

func (o *PutAPIV3StoriesStoryPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/{story-public-id}][%d] putApiV3StoriesStoryPublicIdUnprocessableEntity ", 422)
}

func (o *PutAPIV3StoriesStoryPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
