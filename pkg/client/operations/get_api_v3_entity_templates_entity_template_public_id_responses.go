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

// GetAPIV3EntityTemplatesEntityTemplatePublicIDReader is a Reader for the GetAPIV3EntityTemplatesEntityTemplatePublicID structure.
type GetAPIV3EntityTemplatesEntityTemplatePublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3EntityTemplatesEntityTemplatePublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3EntityTemplatesEntityTemplatePublicIDOK creates a GetAPIV3EntityTemplatesEntityTemplatePublicIDOK with default headers values
func NewGetAPIV3EntityTemplatesEntityTemplatePublicIDOK() *GetAPIV3EntityTemplatesEntityTemplatePublicIDOK {
	return &GetAPIV3EntityTemplatesEntityTemplatePublicIDOK{}
}

/*GetAPIV3EntityTemplatesEntityTemplatePublicIDOK handles this case with default header values.

Resource
*/
type GetAPIV3EntityTemplatesEntityTemplatePublicIDOK struct {
	Payload *models.EntityTemplate
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/entity-templates/{entity-template-public-id}][%d] getApiV3EntityTemplatesEntityTemplatePublicIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDOK) GetPayload() *models.EntityTemplate {
	return o.Payload
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest creates a GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest with default headers values
func NewGetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest() *GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest {
	return &GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest{}
}

/*GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest struct {
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/entity-templates/{entity-template-public-id}][%d] getApiV3EntityTemplatesEntityTemplatePublicIdBadRequest ", 400)
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound creates a GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound with default headers values
func NewGetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound() *GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound {
	return &GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound{}
}

/*GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound struct {
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/entity-templates/{entity-template-public-id}][%d] getApiV3EntityTemplatesEntityTemplatePublicIdNotFound ", 404)
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity creates a GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity with default headers values
func NewGetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity() *GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity {
	return &GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity{}
}

/*GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity struct {
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/entity-templates/{entity-template-public-id}][%d] getApiV3EntityTemplatesEntityTemplatePublicIdUnprocessableEntity ", 422)
}

func (o *GetAPIV3EntityTemplatesEntityTemplatePublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
