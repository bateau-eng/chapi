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

// PutAPIV3StoryLinksStoryLinkPublicIDReader is a Reader for the PutAPIV3StoryLinksStoryLinkPublicID structure.
type PutAPIV3StoryLinksStoryLinkPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3StoryLinksStoryLinkPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV3StoryLinksStoryLinkPublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3StoryLinksStoryLinkPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3StoryLinksStoryLinkPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3StoryLinksStoryLinkPublicIDOK creates a PutAPIV3StoryLinksStoryLinkPublicIDOK with default headers values
func NewPutAPIV3StoryLinksStoryLinkPublicIDOK() *PutAPIV3StoryLinksStoryLinkPublicIDOK {
	return &PutAPIV3StoryLinksStoryLinkPublicIDOK{}
}

/*PutAPIV3StoryLinksStoryLinkPublicIDOK handles this case with default header values.

Resource
*/
type PutAPIV3StoryLinksStoryLinkPublicIDOK struct {
	Payload *models.StoryLink
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] putApiV3StoryLinksStoryLinkPublicIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDOK) GetPayload() *models.StoryLink {
	return o.Payload
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoryLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV3StoryLinksStoryLinkPublicIDBadRequest creates a PutAPIV3StoryLinksStoryLinkPublicIDBadRequest with default headers values
func NewPutAPIV3StoryLinksStoryLinkPublicIDBadRequest() *PutAPIV3StoryLinksStoryLinkPublicIDBadRequest {
	return &PutAPIV3StoryLinksStoryLinkPublicIDBadRequest{}
}

/*PutAPIV3StoryLinksStoryLinkPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3StoryLinksStoryLinkPublicIDBadRequest struct {
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] putApiV3StoryLinksStoryLinkPublicIdBadRequest ", 400)
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoryLinksStoryLinkPublicIDNotFound creates a PutAPIV3StoryLinksStoryLinkPublicIDNotFound with default headers values
func NewPutAPIV3StoryLinksStoryLinkPublicIDNotFound() *PutAPIV3StoryLinksStoryLinkPublicIDNotFound {
	return &PutAPIV3StoryLinksStoryLinkPublicIDNotFound{}
}

/*PutAPIV3StoryLinksStoryLinkPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3StoryLinksStoryLinkPublicIDNotFound struct {
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] putApiV3StoryLinksStoryLinkPublicIdNotFound ", 404)
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity creates a PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity with default headers values
func NewPutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity() *PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity {
	return &PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity{}
}

/*PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity struct {
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/story-links/{story-link-public-id}][%d] putApiV3StoryLinksStoryLinkPublicIdUnprocessableEntity ", 422)
}

func (o *PutAPIV3StoryLinksStoryLinkPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
