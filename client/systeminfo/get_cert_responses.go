// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loheagn/harbor-client/models"
)

// GetCertReader is a Reader for the GetCert structure.
type GetCertReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetCertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCertOK creates a GetCertOK with default headers values
func NewGetCertOK(writer io.Writer) *GetCertOK {
	return &GetCertOK{

		Payload: writer,
	}
}

/* GetCertOK describes a response with status code 200, with default header values.

Get default root certificate successfully.
*/
type GetCertOK struct {

	/* To set the filename of the downloaded file.
	 */
	ContentDisposition string

	Payload io.Writer
}

func (o *GetCertOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getCertOK  %+v", 200, o.Payload)
}
func (o *GetCertOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetCertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Disposition
	hdrContentDisposition := response.GetHeader("Content-Disposition")

	if hdrContentDisposition != "" {
		o.ContentDisposition = hdrContentDisposition
	}

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertNotFound creates a GetCertNotFound with default headers values
func NewGetCertNotFound() *GetCertNotFound {
	return &GetCertNotFound{}
}

/* GetCertNotFound describes a response with status code 404, with default header values.

Not found the default root certificate.
*/
type GetCertNotFound struct {
}

func (o *GetCertNotFound) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getCertNotFound ", 404)
}

func (o *GetCertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCertInternalServerError creates a GetCertInternalServerError with default headers values
func NewGetCertInternalServerError() *GetCertInternalServerError {
	return &GetCertInternalServerError{}
}

/* GetCertInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCertInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetCertInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo/getcert][%d] getCertInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCertInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetCertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
