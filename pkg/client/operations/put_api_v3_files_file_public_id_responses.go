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

// PutAPIV3FilesFilePublicIDReader is a Reader for the PutAPIV3FilesFilePublicID structure.
type PutAPIV3FilesFilePublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3FilesFilePublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV3FilesFilePublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3FilesFilePublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3FilesFilePublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3FilesFilePublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3FilesFilePublicIDOK creates a PutAPIV3FilesFilePublicIDOK with default headers values
func NewPutAPIV3FilesFilePublicIDOK() *PutAPIV3FilesFilePublicIDOK {
	return &PutAPIV3FilesFilePublicIDOK{}
}

/*PutAPIV3FilesFilePublicIDOK handles this case with default header values.

Resource
*/
type PutAPIV3FilesFilePublicIDOK struct {
	Payload *models.File
}

func (o *PutAPIV3FilesFilePublicIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] putApiV3FilesFilePublicIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3FilesFilePublicIDOK) GetPayload() *models.File {
	return o.Payload
}

func (o *PutAPIV3FilesFilePublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV3FilesFilePublicIDBadRequest creates a PutAPIV3FilesFilePublicIDBadRequest with default headers values
func NewPutAPIV3FilesFilePublicIDBadRequest() *PutAPIV3FilesFilePublicIDBadRequest {
	return &PutAPIV3FilesFilePublicIDBadRequest{}
}

/*PutAPIV3FilesFilePublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3FilesFilePublicIDBadRequest struct {
}

func (o *PutAPIV3FilesFilePublicIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] putApiV3FilesFilePublicIdBadRequest ", 400)
}

func (o *PutAPIV3FilesFilePublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3FilesFilePublicIDNotFound creates a PutAPIV3FilesFilePublicIDNotFound with default headers values
func NewPutAPIV3FilesFilePublicIDNotFound() *PutAPIV3FilesFilePublicIDNotFound {
	return &PutAPIV3FilesFilePublicIDNotFound{}
}

/*PutAPIV3FilesFilePublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3FilesFilePublicIDNotFound struct {
}

func (o *PutAPIV3FilesFilePublicIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] putApiV3FilesFilePublicIdNotFound ", 404)
}

func (o *PutAPIV3FilesFilePublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3FilesFilePublicIDUnprocessableEntity creates a PutAPIV3FilesFilePublicIDUnprocessableEntity with default headers values
func NewPutAPIV3FilesFilePublicIDUnprocessableEntity() *PutAPIV3FilesFilePublicIDUnprocessableEntity {
	return &PutAPIV3FilesFilePublicIDUnprocessableEntity{}
}

/*PutAPIV3FilesFilePublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3FilesFilePublicIDUnprocessableEntity struct {
}

func (o *PutAPIV3FilesFilePublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/files/{file-public-id}][%d] putApiV3FilesFilePublicIdUnprocessableEntity ", 422)
}

func (o *PutAPIV3FilesFilePublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}