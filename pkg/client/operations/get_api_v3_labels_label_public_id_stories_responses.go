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

// GetAPIV3LabelsLabelPublicIDStoriesReader is a Reader for the GetAPIV3LabelsLabelPublicIDStories structure.
type GetAPIV3LabelsLabelPublicIDStoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3LabelsLabelPublicIDStoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3LabelsLabelPublicIDStoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3LabelsLabelPublicIDStoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3LabelsLabelPublicIDStoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3LabelsLabelPublicIDStoriesOK creates a GetAPIV3LabelsLabelPublicIDStoriesOK with default headers values
func NewGetAPIV3LabelsLabelPublicIDStoriesOK() *GetAPIV3LabelsLabelPublicIDStoriesOK {
	return &GetAPIV3LabelsLabelPublicIDStoriesOK{}
}

/*GetAPIV3LabelsLabelPublicIDStoriesOK handles this case with default header values.

Resource
*/
type GetAPIV3LabelsLabelPublicIDStoriesOK struct {
	Payload []*models.StorySlim
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] getApiV3LabelsLabelPublicIdStoriesOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesOK) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3LabelsLabelPublicIDStoriesBadRequest creates a GetAPIV3LabelsLabelPublicIDStoriesBadRequest with default headers values
func NewGetAPIV3LabelsLabelPublicIDStoriesBadRequest() *GetAPIV3LabelsLabelPublicIDStoriesBadRequest {
	return &GetAPIV3LabelsLabelPublicIDStoriesBadRequest{}
}

/*GetAPIV3LabelsLabelPublicIDStoriesBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3LabelsLabelPublicIDStoriesBadRequest struct {
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] getApiV3LabelsLabelPublicIdStoriesBadRequest ", 400)
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3LabelsLabelPublicIDStoriesNotFound creates a GetAPIV3LabelsLabelPublicIDStoriesNotFound with default headers values
func NewGetAPIV3LabelsLabelPublicIDStoriesNotFound() *GetAPIV3LabelsLabelPublicIDStoriesNotFound {
	return &GetAPIV3LabelsLabelPublicIDStoriesNotFound{}
}

/*GetAPIV3LabelsLabelPublicIDStoriesNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3LabelsLabelPublicIDStoriesNotFound struct {
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] getApiV3LabelsLabelPublicIdStoriesNotFound ", 404)
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity creates a GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity with default headers values
func NewGetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity() *GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity {
	return &GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity{}
}

/*GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity struct {
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/labels/{label-public-id}/stories][%d] getApiV3LabelsLabelPublicIdStoriesUnprocessableEntity ", 422)
}

func (o *GetAPIV3LabelsLabelPublicIDStoriesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
