// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListFormatsParams creates a new ListFormatsParams object
// with the default values initialized.
func NewListFormatsParams() *ListFormatsParams {
	var ()
	return &ListFormatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFormatsParamsWithTimeout creates a new ListFormatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFormatsParamsWithTimeout(timeout time.Duration) *ListFormatsParams {
	var ()
	return &ListFormatsParams{

		timeout: timeout,
	}
}

// NewListFormatsParamsWithContext creates a new ListFormatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFormatsParamsWithContext(ctx context.Context) *ListFormatsParams {
	var ()
	return &ListFormatsParams{

		Context: ctx,
	}
}

// NewListFormatsParamsWithHTTPClient creates a new ListFormatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFormatsParamsWithHTTPClient(client *http.Client) *ListFormatsParams {
	var ()
	return &ListFormatsParams{
		HTTPClient: client,
	}
}

/*ListFormatsParams contains all the parameters to send to the API endpoint
for the list formats operation typically these are written to a http.Request
*/
type ListFormatsParams struct {

	/*DistroID
	  desired distro

	*/
	DistroID strfmt.UUID
	/*OrganizationName
	  desired organization

	*/
	OrganizationName string
	/*ProjectName
	  desired project

	*/
	ProjectName string
	/*ReleaseID
	  desired release

	*/
	ReleaseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list formats params
func (o *ListFormatsParams) WithTimeout(timeout time.Duration) *ListFormatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list formats params
func (o *ListFormatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list formats params
func (o *ListFormatsParams) WithContext(ctx context.Context) *ListFormatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list formats params
func (o *ListFormatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list formats params
func (o *ListFormatsParams) WithHTTPClient(client *http.Client) *ListFormatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list formats params
func (o *ListFormatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDistroID adds the distroID to the list formats params
func (o *ListFormatsParams) WithDistroID(distroID strfmt.UUID) *ListFormatsParams {
	o.SetDistroID(distroID)
	return o
}

// SetDistroID adds the distroId to the list formats params
func (o *ListFormatsParams) SetDistroID(distroID strfmt.UUID) {
	o.DistroID = distroID
}

// WithOrganizationName adds the organizationName to the list formats params
func (o *ListFormatsParams) WithOrganizationName(organizationName string) *ListFormatsParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the list formats params
func (o *ListFormatsParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the list formats params
func (o *ListFormatsParams) WithProjectName(projectName string) *ListFormatsParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the list formats params
func (o *ListFormatsParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReleaseID adds the releaseID to the list formats params
func (o *ListFormatsParams) WithReleaseID(releaseID strfmt.UUID) *ListFormatsParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the list formats params
func (o *ListFormatsParams) SetReleaseID(releaseID strfmt.UUID) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *ListFormatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param distroID
	if err := r.SetPathParam("distroID", o.DistroID.String()); err != nil {
		return err
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	// path param releaseID
	if err := r.SetPathParam("releaseID", o.ReleaseID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
