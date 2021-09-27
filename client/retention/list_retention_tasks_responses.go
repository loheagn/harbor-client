// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loheagn/harbor-client/models"
)

// ListRetentionTasksReader is a Reader for the ListRetentionTasks structure.
type ListRetentionTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRetentionTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRetentionTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRetentionTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRetentionTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRetentionTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRetentionTasksOK creates a ListRetentionTasksOK with default headers values
func NewListRetentionTasksOK() *ListRetentionTasksOK {
	return &ListRetentionTasksOK{}
}

/* ListRetentionTasksOK describes a response with status code 200, with default header values.

Get Retention job tasks successfully.
*/
type ListRetentionTasksOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.RetentionExecutionTask
}

func (o *ListRetentionTasksOK) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks][%d] listRetentionTasksOK  %+v", 200, o.Payload)
}
func (o *ListRetentionTasksOK) GetPayload() []*models.RetentionExecutionTask {
	return o.Payload
}

func (o *ListRetentionTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRetentionTasksUnauthorized creates a ListRetentionTasksUnauthorized with default headers values
func NewListRetentionTasksUnauthorized() *ListRetentionTasksUnauthorized {
	return &ListRetentionTasksUnauthorized{}
}

/* ListRetentionTasksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListRetentionTasksUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRetentionTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks][%d] listRetentionTasksUnauthorized  %+v", 401, o.Payload)
}
func (o *ListRetentionTasksUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRetentionTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRetentionTasksForbidden creates a ListRetentionTasksForbidden with default headers values
func NewListRetentionTasksForbidden() *ListRetentionTasksForbidden {
	return &ListRetentionTasksForbidden{}
}

/* ListRetentionTasksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListRetentionTasksForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRetentionTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks][%d] listRetentionTasksForbidden  %+v", 403, o.Payload)
}
func (o *ListRetentionTasksForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRetentionTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRetentionTasksInternalServerError creates a ListRetentionTasksInternalServerError with default headers values
func NewListRetentionTasksInternalServerError() *ListRetentionTasksInternalServerError {
	return &ListRetentionTasksInternalServerError{}
}

/* ListRetentionTasksInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRetentionTasksInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRetentionTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks][%d] listRetentionTasksInternalServerError  %+v", 500, o.Payload)
}
func (o *ListRetentionTasksInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRetentionTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
