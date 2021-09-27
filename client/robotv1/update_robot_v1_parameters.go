// Code generated by go-swagger; DO NOT EDIT.

package robotv1

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

// NewUpdateRobotV1Params creates a new UpdateRobotV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRobotV1Params() *UpdateRobotV1Params {
	return &UpdateRobotV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRobotV1ParamsWithTimeout creates a new UpdateRobotV1Params object
// with the ability to set a timeout on a request.
func NewUpdateRobotV1ParamsWithTimeout(timeout time.Duration) *UpdateRobotV1Params {
	return &UpdateRobotV1Params{
		timeout: timeout,
	}
}

// NewUpdateRobotV1ParamsWithContext creates a new UpdateRobotV1Params object
// with the ability to set a context for a request.
func NewUpdateRobotV1ParamsWithContext(ctx context.Context) *UpdateRobotV1Params {
	return &UpdateRobotV1Params{
		Context: ctx,
	}
}

// NewUpdateRobotV1ParamsWithHTTPClient creates a new UpdateRobotV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRobotV1ParamsWithHTTPClient(client *http.Client) *UpdateRobotV1Params {
	return &UpdateRobotV1Params{
		HTTPClient: client,
	}
}

/* UpdateRobotV1Params contains all the parameters to send to the API endpoint
   for the update robot v1 operation.

   Typically these are written to a http.Request.
*/
type UpdateRobotV1Params struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	/* Robot.

	   The JSON object of a robot account.
	*/
	Robot *models.Robot

	/* RobotID.

	   Robot ID
	*/
	RobotID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRobotV1Params) WithDefaults() *UpdateRobotV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRobotV1Params) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := UpdateRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update robot v1 params
func (o *UpdateRobotV1Params) WithTimeout(timeout time.Duration) *UpdateRobotV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update robot v1 params
func (o *UpdateRobotV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update robot v1 params
func (o *UpdateRobotV1Params) WithContext(ctx context.Context) *UpdateRobotV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update robot v1 params
func (o *UpdateRobotV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update robot v1 params
func (o *UpdateRobotV1Params) WithHTTPClient(client *http.Client) *UpdateRobotV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update robot v1 params
func (o *UpdateRobotV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the update robot v1 params
func (o *UpdateRobotV1Params) WithXIsResourceName(xIsResourceName *bool) *UpdateRobotV1Params {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update robot v1 params
func (o *UpdateRobotV1Params) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update robot v1 params
func (o *UpdateRobotV1Params) WithXRequestID(xRequestID *string) *UpdateRobotV1Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update robot v1 params
func (o *UpdateRobotV1Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the update robot v1 params
func (o *UpdateRobotV1Params) WithProjectNameOrID(projectNameOrID string) *UpdateRobotV1Params {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update robot v1 params
func (o *UpdateRobotV1Params) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithRobot adds the robot to the update robot v1 params
func (o *UpdateRobotV1Params) WithRobot(robot *models.Robot) *UpdateRobotV1Params {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the update robot v1 params
func (o *UpdateRobotV1Params) SetRobot(robot *models.Robot) {
	o.Robot = robot
}

// WithRobotID adds the robotID to the update robot v1 params
func (o *UpdateRobotV1Params) WithRobotID(robotID int64) *UpdateRobotV1Params {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the update robot v1 params
func (o *UpdateRobotV1Params) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRobotV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}
	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
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
