// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewAddLabelParams creates a new AddLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddLabelParams() *AddLabelParams {
	return &AddLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddLabelParamsWithTimeout creates a new AddLabelParams object
// with the ability to set a timeout on a request.
func NewAddLabelParamsWithTimeout(timeout time.Duration) *AddLabelParams {
	return &AddLabelParams{
		timeout: timeout,
	}
}

// NewAddLabelParamsWithContext creates a new AddLabelParams object
// with the ability to set a context for a request.
func NewAddLabelParamsWithContext(ctx context.Context) *AddLabelParams {
	return &AddLabelParams{
		Context: ctx,
	}
}

// NewAddLabelParamsWithHTTPClient creates a new AddLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddLabelParamsWithHTTPClient(client *http.Client) *AddLabelParams {
	return &AddLabelParams{
		HTTPClient: client,
	}
}

/* AddLabelParams contains all the parameters to send to the API endpoint
   for the add label operation.

   Typically these are written to a http.Request.
*/
type AddLabelParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Label.

	   The label that added to the artifact. Only the ID property is needed.
	*/
	Label *models.Label

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	/* Reference.

	   The reference of the artifact, can be digest or tag
	*/
	Reference string

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLabelParams) WithDefaults() *AddLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add label params
func (o *AddLabelParams) WithTimeout(timeout time.Duration) *AddLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add label params
func (o *AddLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add label params
func (o *AddLabelParams) WithContext(ctx context.Context) *AddLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add label params
func (o *AddLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add label params
func (o *AddLabelParams) WithHTTPClient(client *http.Client) *AddLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add label params
func (o *AddLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the add label params
func (o *AddLabelParams) WithXRequestID(xRequestID *string) *AddLabelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the add label params
func (o *AddLabelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithLabel adds the label to the add label params
func (o *AddLabelParams) WithLabel(label *models.Label) *AddLabelParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the add label params
func (o *AddLabelParams) SetLabel(label *models.Label) {
	o.Label = label
}

// WithProjectName adds the projectName to the add label params
func (o *AddLabelParams) WithProjectName(projectName string) *AddLabelParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the add label params
func (o *AddLabelParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the add label params
func (o *AddLabelParams) WithReference(reference string) *AddLabelParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the add label params
func (o *AddLabelParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the add label params
func (o *AddLabelParams) WithRepositoryName(repositoryName string) *AddLabelParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the add label params
func (o *AddLabelParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *AddLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
