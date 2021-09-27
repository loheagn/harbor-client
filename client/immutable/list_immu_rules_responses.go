// Code generated by go-swagger; DO NOT EDIT.

package immutable

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

// ListImmuRulesReader is a Reader for the ListImmuRules structure.
type ListImmuRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImmuRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImmuRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListImmuRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListImmuRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListImmuRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListImmuRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListImmuRulesOK creates a ListImmuRulesOK with default headers values
func NewListImmuRulesOK() *ListImmuRulesOK {
	return &ListImmuRulesOK{}
}

/* ListImmuRulesOK describes a response with status code 200, with default header values.

Success
*/
type ListImmuRulesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of immutable tag
	 */
	XTotalCount int64

	Payload []*models.ImmutableRule
}

func (o *ListImmuRulesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/immutabletagrules][%d] listImmuRulesOK  %+v", 200, o.Payload)
}
func (o *ListImmuRulesOK) GetPayload() []*models.ImmutableRule {
	return o.Payload
}

func (o *ListImmuRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListImmuRulesBadRequest creates a ListImmuRulesBadRequest with default headers values
func NewListImmuRulesBadRequest() *ListImmuRulesBadRequest {
	return &ListImmuRulesBadRequest{}
}

/* ListImmuRulesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListImmuRulesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListImmuRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/immutabletagrules][%d] listImmuRulesBadRequest  %+v", 400, o.Payload)
}
func (o *ListImmuRulesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListImmuRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListImmuRulesUnauthorized creates a ListImmuRulesUnauthorized with default headers values
func NewListImmuRulesUnauthorized() *ListImmuRulesUnauthorized {
	return &ListImmuRulesUnauthorized{}
}

/* ListImmuRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListImmuRulesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListImmuRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/immutabletagrules][%d] listImmuRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListImmuRulesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListImmuRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListImmuRulesForbidden creates a ListImmuRulesForbidden with default headers values
func NewListImmuRulesForbidden() *ListImmuRulesForbidden {
	return &ListImmuRulesForbidden{}
}

/* ListImmuRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListImmuRulesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListImmuRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/immutabletagrules][%d] listImmuRulesForbidden  %+v", 403, o.Payload)
}
func (o *ListImmuRulesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListImmuRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListImmuRulesInternalServerError creates a ListImmuRulesInternalServerError with default headers values
func NewListImmuRulesInternalServerError() *ListImmuRulesInternalServerError {
	return &ListImmuRulesInternalServerError{}
}

/* ListImmuRulesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListImmuRulesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListImmuRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/immutabletagrules][%d] listImmuRulesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListImmuRulesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListImmuRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
