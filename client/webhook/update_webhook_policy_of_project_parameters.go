// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/loheagn/harbor-client/models"
)

// NewUpdateWebhookPolicyOfProjectParams creates a new UpdateWebhookPolicyOfProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWebhookPolicyOfProjectParams() *UpdateWebhookPolicyOfProjectParams {
	return &UpdateWebhookPolicyOfProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWebhookPolicyOfProjectParamsWithTimeout creates a new UpdateWebhookPolicyOfProjectParams object
// with the ability to set a timeout on a request.
func NewUpdateWebhookPolicyOfProjectParamsWithTimeout(timeout time.Duration) *UpdateWebhookPolicyOfProjectParams {
	return &UpdateWebhookPolicyOfProjectParams{
		timeout: timeout,
	}
}

// NewUpdateWebhookPolicyOfProjectParamsWithContext creates a new UpdateWebhookPolicyOfProjectParams object
// with the ability to set a context for a request.
func NewUpdateWebhookPolicyOfProjectParamsWithContext(ctx context.Context) *UpdateWebhookPolicyOfProjectParams {
	return &UpdateWebhookPolicyOfProjectParams{
		Context: ctx,
	}
}

// NewUpdateWebhookPolicyOfProjectParamsWithHTTPClient creates a new UpdateWebhookPolicyOfProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWebhookPolicyOfProjectParamsWithHTTPClient(client *http.Client) *UpdateWebhookPolicyOfProjectParams {
	return &UpdateWebhookPolicyOfProjectParams{
		HTTPClient: client,
	}
}

/* UpdateWebhookPolicyOfProjectParams contains all the parameters to send to the API endpoint
   for the update webhook policy of project operation.

   Typically these are written to a http.Request.
*/
type UpdateWebhookPolicyOfProjectParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Policy.

	   All properties needed except "id", "project_id", "creation_time", "update_time".
	*/
	Policy *models.WebhookPolicy

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	/* WebhookPolicyID.

	   The ID of the webhook policy

	   Format: int64
	*/
	WebhookPolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update webhook policy of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebhookPolicyOfProjectParams) WithDefaults() *UpdateWebhookPolicyOfProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update webhook policy of project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebhookPolicyOfProjectParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := UpdateWebhookPolicyOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithTimeout(timeout time.Duration) *UpdateWebhookPolicyOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithContext(ctx context.Context) *UpdateWebhookPolicyOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithHTTPClient(client *http.Client) *UpdateWebhookPolicyOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithXIsResourceName(xIsResourceName *bool) *UpdateWebhookPolicyOfProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithXRequestID(xRequestID *string) *UpdateWebhookPolicyOfProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithPolicy(policy *models.WebhookPolicy) *UpdateWebhookPolicyOfProjectParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetPolicy(policy *models.WebhookPolicy) {
	o.Policy = policy
}

// WithProjectNameOrID adds the projectNameOrID to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithProjectNameOrID(projectNameOrID string) *UpdateWebhookPolicyOfProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithWebhookPolicyID adds the webhookPolicyID to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) WithWebhookPolicyID(webhookPolicyID int64) *UpdateWebhookPolicyOfProjectParams {
	o.SetWebhookPolicyID(webhookPolicyID)
	return o
}

// SetWebhookPolicyID adds the webhookPolicyId to the update webhook policy of project params
func (o *UpdateWebhookPolicyOfProjectParams) SetWebhookPolicyID(webhookPolicyID int64) {
	o.WebhookPolicyID = webhookPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWebhookPolicyOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	// path param webhook_policy_id
	if err := r.SetPathParam("webhook_policy_id", swag.FormatInt64(o.WebhookPolicyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
