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

// GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader is a Reader for the GetAPIV3EpicsEpicPublicIDCommentsCommentPublicID structure.
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK creates a GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK with default headers values
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK() *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK {
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK{}
}

/*GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK handles this case with default header values.

Resource
*/
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK struct {
	Payload *models.ThreadedComment
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] getApiV3EpicsEpicPublicIdCommentsCommentPublicIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK) GetPayload() *models.ThreadedComment {
	return o.Payload
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreadedComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest creates a GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest with default headers values
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest() *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest {
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest{}
}

/*GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest struct {
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] getApiV3EpicsEpicPublicIdCommentsCommentPublicIdBadRequest ", 400)
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound creates a GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound with default headers values
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound() *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound {
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound{}
}

/*GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound struct {
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] getApiV3EpicsEpicPublicIdCommentsCommentPublicIdNotFound ", 404)
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity creates a GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity with default headers values
func NewGetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity() *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity {
	return &GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity{}
}

/*GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity struct {
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics/{epic-public-id}/comments/{comment-public-id}][%d] getApiV3EpicsEpicPublicIdCommentsCommentPublicIdUnprocessableEntity ", 422)
}

func (o *GetAPIV3EpicsEpicPublicIDCommentsCommentPublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
