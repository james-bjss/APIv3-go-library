// Code generated by go-swagger; DO NOT EDIT.

package transactional_s_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// GetTransacSmsReportReader is a Reader for the GetTransacSmsReport structure.
type GetTransacSmsReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransacSmsReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransacSmsReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransacSmsReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransacSmsReportOK creates a GetTransacSmsReportOK with default headers values
func NewGetTransacSmsReportOK() *GetTransacSmsReportOK {
	return &GetTransacSmsReportOK{}
}

/*GetTransacSmsReportOK handles this case with default header values.

Aggregated SMS report informations
*/
type GetTransacSmsReportOK struct {
	Payload *models.GetTransacSmsReport
}

func (o *GetTransacSmsReportOK) Error() string {
	return fmt.Sprintf("[GET /transactionalSMS/statistics/reports][%d] getTransacSmsReportOK  %+v", 200, o.Payload)
}

func (o *GetTransacSmsReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTransacSmsReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransacSmsReportBadRequest creates a GetTransacSmsReportBadRequest with default headers values
func NewGetTransacSmsReportBadRequest() *GetTransacSmsReportBadRequest {
	return &GetTransacSmsReportBadRequest{}
}

/*GetTransacSmsReportBadRequest handles this case with default header values.

bad request
*/
type GetTransacSmsReportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetTransacSmsReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /transactionalSMS/statistics/reports][%d] getTransacSmsReportBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransacSmsReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
