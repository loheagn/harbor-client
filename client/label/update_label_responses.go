// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loheagn/harbor-client/models"
)

// UpdateLabelReader is a Reader for the UpdateLabel structure.
type UpdateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLabelOK creates a UpdateLabelOK with default headers values
func NewUpdateLabelOK() *UpdateLabelOK {
	return &UpdateLabelOK{}
}

/* UpdateLabelOK describes a response with status code 200, with default header values.

Success
*/
type UpdateLabelOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateLabelOK) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelOK ", 200)
}

func (o *UpdateLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateLabelBadRequest creates a UpdateLabelBadRequest with default headers values
func NewUpdateLabelBadRequest() *UpdateLabelBadRequest {
	return &UpdateLabelBadRequest{}
}

/* UpdateLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateLabelBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateLabelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateLabelBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelUnauthorized creates a UpdateLabelUnauthorized with default headers values
func NewUpdateLabelUnauthorized() *UpdateLabelUnauthorized {
	return &UpdateLabelUnauthorized{}
}

/* UpdateLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateLabelUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelNotFound creates a UpdateLabelNotFound with default headers values
func NewUpdateLabelNotFound() *UpdateLabelNotFound {
	return &UpdateLabelNotFound{}
}

/* UpdateLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateLabelNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateLabelNotFound) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelNotFound  %+v", 404, o.Payload)
}
func (o *UpdateLabelNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelConflict creates a UpdateLabelConflict with default headers values
func NewUpdateLabelConflict() *UpdateLabelConflict {
	return &UpdateLabelConflict{}
}

/* UpdateLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateLabelConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateLabelConflict) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelConflict  %+v", 409, o.Payload)
}
func (o *UpdateLabelConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelInternalServerError creates a UpdateLabelInternalServerError with default headers values
func NewUpdateLabelInternalServerError() *UpdateLabelInternalServerError {
	return &UpdateLabelInternalServerError{}
}

/* UpdateLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateLabelInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /labels/{label_id}][%d] updateLabelInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
