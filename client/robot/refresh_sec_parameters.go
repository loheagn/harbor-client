// Code generated by go-swagger; DO NOT EDIT.

package robot

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

// NewRefreshSecParams creates a new RefreshSecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefreshSecParams() *RefreshSecParams {
	return &RefreshSecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshSecParamsWithTimeout creates a new RefreshSecParams object
// with the ability to set a timeout on a request.
func NewRefreshSecParamsWithTimeout(timeout time.Duration) *RefreshSecParams {
	return &RefreshSecParams{
		timeout: timeout,
	}
}

// NewRefreshSecParamsWithContext creates a new RefreshSecParams object
// with the ability to set a context for a request.
func NewRefreshSecParamsWithContext(ctx context.Context) *RefreshSecParams {
	return &RefreshSecParams{
		Context: ctx,
	}
}

// NewRefreshSecParamsWithHTTPClient creates a new RefreshSecParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefreshSecParamsWithHTTPClient(client *http.Client) *RefreshSecParams {
	return &RefreshSecParams{
		HTTPClient: client,
	}
}

/* RefreshSecParams contains all the parameters to send to the API endpoint
   for the refresh sec operation.

   Typically these are written to a http.Request.
*/
type RefreshSecParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* RobotSec.

	   The JSON object of a robot account.
	*/
	RobotSec *models.RobotSec

	/* RobotID.

	   Robot ID
	*/
	RobotID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refresh sec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshSecParams) WithDefaults() *RefreshSecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refresh sec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshSecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the refresh sec params
func (o *RefreshSecParams) WithTimeout(timeout time.Duration) *RefreshSecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh sec params
func (o *RefreshSecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh sec params
func (o *RefreshSecParams) WithContext(ctx context.Context) *RefreshSecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh sec params
func (o *RefreshSecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh sec params
func (o *RefreshSecParams) WithHTTPClient(client *http.Client) *RefreshSecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh sec params
func (o *RefreshSecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the refresh sec params
func (o *RefreshSecParams) WithXRequestID(xRequestID *string) *RefreshSecParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the refresh sec params
func (o *RefreshSecParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRobotSec adds the robotSec to the refresh sec params
func (o *RefreshSecParams) WithRobotSec(robotSec *models.RobotSec) *RefreshSecParams {
	o.SetRobotSec(robotSec)
	return o
}

// SetRobotSec adds the robotSec to the refresh sec params
func (o *RefreshSecParams) SetRobotSec(robotSec *models.RobotSec) {
	o.RobotSec = robotSec
}

// WithRobotID adds the robotID to the refresh sec params
func (o *RefreshSecParams) WithRobotID(robotID int64) *RefreshSecParams {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the refresh sec params
func (o *RefreshSecParams) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshSecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.RobotSec != nil {
		if err := r.SetBodyParam(o.RobotSec); err != nil {
			return err
		}
	}

	// path param robot_id
	if err := r.SetPathParam("robot_id", swag.FormatInt64(o.RobotID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
