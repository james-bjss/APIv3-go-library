// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// DeleteListReader is a Reader for the DeleteList structure.
type DeleteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteListNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteListNoContent creates a DeleteListNoContent with default headers values
func NewDeleteListNoContent() *DeleteListNoContent {
	return &DeleteListNoContent{}
}

/*DeleteListNoContent handles this case with default header values.

List deleted
*/
type DeleteListNoContent struct {
}

func (o *DeleteListNoContent) Error() string {
	return fmt.Sprintf("[DELETE /contacts/lists/{listId}][%d] deleteListNoContent ", 204)
}

func (o *DeleteListNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListBadRequest creates a DeleteListBadRequest with default headers values
func NewDeleteListBadRequest() *DeleteListBadRequest {
	return &DeleteListBadRequest{}
}

/*DeleteListBadRequest handles this case with default header values.

bad request
*/
type DeleteListBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteListBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /contacts/lists/{listId}][%d] deleteListBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteListNotFound creates a DeleteListNotFound with default headers values
func NewDeleteListNotFound() *DeleteListNotFound {
	return &DeleteListNotFound{}
}

/*DeleteListNotFound handles this case with default header values.

List ID not found
*/
type DeleteListNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteListNotFound) Error() string {
	return fmt.Sprintf("[DELETE /contacts/lists/{listId}][%d] deleteListNotFound  %+v", 404, o.Payload)
}

func (o *DeleteListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
