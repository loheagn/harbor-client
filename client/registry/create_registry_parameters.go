// Code generated by go-swagger; DO NOT EDIT.

package registry

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

	"github.com/loheagn/harbor-client/models"
)

// NewCreateRegistryParams creates a new CreateRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRegistryParams() *CreateRegistryParams {
	return &CreateRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRegistryParamsWithTimeout creates a new CreateRegistryParams object
// with the ability to set a timeout on a request.
func NewCreateRegistryParamsWithTimeout(timeout time.Duration) *CreateRegistryParams {
	return &CreateRegistryParams{
		timeout: timeout,
	}
}

// NewCreateRegistryParamsWithContext creates a new CreateRegistryParams object
// with the ability to set a context for a request.
func NewCreateRegistryParamsWithContext(ctx context.Context) *CreateRegistryParams {
	return &CreateRegistryParams{
		Context: ctx,
	}
}

// NewCreateRegistryParamsWithHTTPClient creates a new CreateRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRegistryParamsWithHTTPClient(client *http.Client) *CreateRegistryParams {
	return &CreateRegistryParams{
		HTTPClient: client,
	}
}

/* CreateRegistryParams contains all the parameters to send to the API endpoint
   for the create registry operation.

   Typically these are written to a http.Request.
*/
type CreateRegistryParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Registry.

	   The registry
	*/
	Registry *models.Registry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegistryParams) WithDefaults() *CreateRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) WithTimeout(timeout time.Duration) *CreateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create registry params
func (o *CreateRegistryParams) WithContext(ctx context.Context) *CreateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create registry params
func (o *CreateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) WithHTTPClient(client *http.Client) *CreateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create registry params
func (o *CreateRegistryParams) WithXRequestID(xRequestID *string) *CreateRegistryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create registry params
func (o *CreateRegistryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistry adds the registry to the create registry params
func (o *CreateRegistryParams) WithRegistry(registry *models.Registry) *CreateRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the create registry params
func (o *CreateRegistryParams) SetRegistry(registry *models.Registry) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Registry != nil {
		if err := r.SetBodyParam(o.Registry); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
