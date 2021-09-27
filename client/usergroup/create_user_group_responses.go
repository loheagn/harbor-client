// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loheagn/harbor-client/models"
)

// CreateUserGroupReader is a Reader for the CreateUserGroup structure.
type CreateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserGroupCreated creates a CreateUserGroupCreated with default headers values
func NewCreateUserGroupCreated() *CreateUserGroupCreated {
	return &CreateUserGroupCreated{}
}

/* CreateUserGroupCreated describes a response with status code 201, with default header values.

User group created successfully.
*/
type CreateUserGroupCreated struct {

	/* The URL of the created resource
	 */
	Location string
}

func (o *CreateUserGroupCreated) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupCreated ", 201)
}

func (o *CreateUserGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewCreateUserGroupBadRequest creates a CreateUserGroupBadRequest with default headers values
func NewCreateUserGroupBadRequest() *CreateUserGroupBadRequest {
	return &CreateUserGroupBadRequest{}
}

/* CreateUserGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateUserGroupBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUserGroupBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupUnauthorized creates a CreateUserGroupUnauthorized with default headers values
func NewCreateUserGroupUnauthorized() *CreateUserGroupUnauthorized {
	return &CreateUserGroupUnauthorized{}
}

/* CreateUserGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateUserGroupUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateUserGroupUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupForbidden creates a CreateUserGroupForbidden with default headers values
func NewCreateUserGroupForbidden() *CreateUserGroupForbidden {
	return &CreateUserGroupForbidden{}
}

/* CreateUserGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserGroupForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupForbidden  %+v", 403, o.Payload)
}
func (o *CreateUserGroupForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupConflict creates a CreateUserGroupConflict with default headers values
func NewCreateUserGroupConflict() *CreateUserGroupConflict {
	return &CreateUserGroupConflict{}
}

/* CreateUserGroupConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateUserGroupConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserGroupConflict) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupConflict  %+v", 409, o.Payload)
}
func (o *CreateUserGroupConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserGroupInternalServerError creates a CreateUserGroupInternalServerError with default headers values
func NewCreateUserGroupInternalServerError() *CreateUserGroupInternalServerError {
	return &CreateUserGroupInternalServerError{}
}

/* CreateUserGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateUserGroupInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateUserGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /usergroups][%d] createUserGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateUserGroupInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateUserGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
