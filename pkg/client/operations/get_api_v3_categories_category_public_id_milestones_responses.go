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

// GetAPIV3CategoriesCategoryPublicIDMilestonesReader is a Reader for the GetAPIV3CategoriesCategoryPublicIDMilestones structure.
type GetAPIV3CategoriesCategoryPublicIDMilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3CategoriesCategoryPublicIDMilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3CategoriesCategoryPublicIDMilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3CategoriesCategoryPublicIDMilestonesOK creates a GetAPIV3CategoriesCategoryPublicIDMilestonesOK with default headers values
func NewGetAPIV3CategoriesCategoryPublicIDMilestonesOK() *GetAPIV3CategoriesCategoryPublicIDMilestonesOK {
	return &GetAPIV3CategoriesCategoryPublicIDMilestonesOK{}
}

/*GetAPIV3CategoriesCategoryPublicIDMilestonesOK handles this case with default header values.

Resource
*/
type GetAPIV3CategoriesCategoryPublicIDMilestonesOK struct {
	Payload []*models.Milestone
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] getApiV3CategoriesCategoryPublicIdMilestonesOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesOK) GetPayload() []*models.Milestone {
	return o.Payload
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest creates a GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest with default headers values
func NewGetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest() *GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest {
	return &GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest{}
}

/*GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest struct {
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] getApiV3CategoriesCategoryPublicIdMilestonesBadRequest ", 400)
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3CategoriesCategoryPublicIDMilestonesNotFound creates a GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound with default headers values
func NewGetAPIV3CategoriesCategoryPublicIDMilestonesNotFound() *GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound {
	return &GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound{}
}

/*GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound struct {
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] getApiV3CategoriesCategoryPublicIdMilestonesNotFound ", 404)
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity creates a GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity with default headers values
func NewGetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity() *GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity {
	return &GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity{}
}

/*GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity struct {
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/categories/{category-public-id}/milestones][%d] getApiV3CategoriesCategoryPublicIdMilestonesUnprocessableEntity ", 422)
}

func (o *GetAPIV3CategoriesCategoryPublicIDMilestonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
