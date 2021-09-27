// Code generated by go-swagger; DO NOT EDIT.

package user

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
)

// NewGetCurrentUserInfoParams creates a new GetCurrentUserInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCurrentUserInfoParams() *GetCurrentUserInfoParams {
	return &GetCurrentUserInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentUserInfoParamsWithTimeout creates a new GetCurrentUserInfoParams object
// with the ability to set a timeout on a request.
func NewGetCurrentUserInfoParamsWithTimeout(timeout time.Duration) *GetCurrentUserInfoParams {
	return &GetCurrentUserInfoParams{
		timeout: timeout,
	}
}

// NewGetCurrentUserInfoParamsWithContext creates a new GetCurrentUserInfoParams object
// with the ability to set a context for a request.
func NewGetCurrentUserInfoParamsWithContext(ctx context.Context) *GetCurrentUserInfoParams {
	return &GetCurrentUserInfoParams{
		Context: ctx,
	}
}

// NewGetCurrentUserInfoParamsWithHTTPClient creates a new GetCurrentUserInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCurrentUserInfoParamsWithHTTPClient(client *http.Client) *GetCurrentUserInfoParams {
	return &GetCurrentUserInfoParams{
		HTTPClient: client,
	}
}

/* GetCurrentUserInfoParams contains all the parameters to send to the API endpoint
   for the get current user info operation.

   Typically these are written to a http.Request.
*/
type GetCurrentUserInfoParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get current user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserInfoParams) WithDefaults() *GetCurrentUserInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get current user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get current user info params
func (o *GetCurrentUserInfoParams) WithTimeout(timeout time.Duration) *GetCurrentUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current user info params
func (o *GetCurrentUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current user info params
func (o *GetCurrentUserInfoParams) WithContext(ctx context.Context) *GetCurrentUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current user info params
func (o *GetCurrentUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current user info params
func (o *GetCurrentUserInfoParams) WithHTTPClient(client *http.Client) *GetCurrentUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current user info params
func (o *GetCurrentUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get current user info params
func (o *GetCurrentUserInfoParams) WithXRequestID(xRequestID *string) *GetCurrentUserInfoParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get current user info params
func (o *GetCurrentUserInfoParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
