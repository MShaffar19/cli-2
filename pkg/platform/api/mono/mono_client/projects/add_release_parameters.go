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
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewAddReleaseParams creates a new AddReleaseParams object
// with the default values initialized.
func NewAddReleaseParams() *AddReleaseParams {
	var ()
	return &AddReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddReleaseParamsWithTimeout creates a new AddReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddReleaseParamsWithTimeout(timeout time.Duration) *AddReleaseParams {
	var ()
	return &AddReleaseParams{

		timeout: timeout,
	}
}

// NewAddReleaseParamsWithContext creates a new AddReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddReleaseParamsWithContext(ctx context.Context) *AddReleaseParams {
	var ()
	return &AddReleaseParams{

		Context: ctx,
	}
}

// NewAddReleaseParamsWithHTTPClient creates a new AddReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddReleaseParamsWithHTTPClient(client *http.Client) *AddReleaseParams {
	var ()
	return &AddReleaseParams{
		HTTPClient: client,
	}
}

/*AddReleaseParams contains all the parameters to send to the API endpoint
for the add release operation typically these are written to a http.Request
*/
type AddReleaseParams struct {

	/*OrganizationName
	  desired organization

	*/
	OrganizationName string
	/*ProjectName
	  desired project

	*/
	ProjectName string
	/*Release
	  Release details

	*/
	Release *mono_models.Release

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add release params
func (o *AddReleaseParams) WithTimeout(timeout time.Duration) *AddReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add release params
func (o *AddReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add release params
func (o *AddReleaseParams) WithContext(ctx context.Context) *AddReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add release params
func (o *AddReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add release params
func (o *AddReleaseParams) WithHTTPClient(client *http.Client) *AddReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add release params
func (o *AddReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the add release params
func (o *AddReleaseParams) WithOrganizationName(organizationName string) *AddReleaseParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the add release params
func (o *AddReleaseParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the add release params
func (o *AddReleaseParams) WithProjectName(projectName string) *AddReleaseParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the add release params
func (o *AddReleaseParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithRelease adds the release to the add release params
func (o *AddReleaseParams) WithRelease(release *mono_models.Release) *AddReleaseParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the add release params
func (o *AddReleaseParams) SetRelease(release *mono_models.Release) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *AddReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if o.Release != nil {
		if err := r.SetBodyParam(o.Release); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
