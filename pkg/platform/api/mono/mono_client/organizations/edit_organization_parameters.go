// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewEditOrganizationParams creates a new EditOrganizationParams object
// with the default values initialized.
func NewEditOrganizationParams() *EditOrganizationParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditOrganizationParams{
		IdentifierType: &identifierTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEditOrganizationParamsWithTimeout creates a new EditOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditOrganizationParamsWithTimeout(timeout time.Duration) *EditOrganizationParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditOrganizationParams{
		IdentifierType: &identifierTypeDefault,

		timeout: timeout,
	}
}

// NewEditOrganizationParamsWithContext creates a new EditOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditOrganizationParamsWithContext(ctx context.Context) *EditOrganizationParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditOrganizationParams{
		IdentifierType: &identifierTypeDefault,

		Context: ctx,
	}
}

// NewEditOrganizationParamsWithHTTPClient creates a new EditOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditOrganizationParamsWithHTTPClient(client *http.Client) *EditOrganizationParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditOrganizationParams{
		IdentifierType: &identifierTypeDefault,
		HTTPClient:     client,
	}
}

/*EditOrganizationParams contains all the parameters to send to the API endpoint
for the edit organization operation typically these are written to a http.Request
*/
type EditOrganizationParams struct {

	/*IdentifierType
	  what kind of thing the provided organizationIdentifier is

	*/
	IdentifierType *string
	/*Organization*/
	Organization *mono_models.OrganizationEditable
	/*OrganizationIdentifier
	  identifier (URLname, by default) of the desired organization

	*/
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit organization params
func (o *EditOrganizationParams) WithTimeout(timeout time.Duration) *EditOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit organization params
func (o *EditOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit organization params
func (o *EditOrganizationParams) WithContext(ctx context.Context) *EditOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit organization params
func (o *EditOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit organization params
func (o *EditOrganizationParams) WithHTTPClient(client *http.Client) *EditOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit organization params
func (o *EditOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifierType adds the identifierType to the edit organization params
func (o *EditOrganizationParams) WithIdentifierType(identifierType *string) *EditOrganizationParams {
	o.SetIdentifierType(identifierType)
	return o
}

// SetIdentifierType adds the identifierType to the edit organization params
func (o *EditOrganizationParams) SetIdentifierType(identifierType *string) {
	o.IdentifierType = identifierType
}

// WithOrganization adds the organization to the edit organization params
func (o *EditOrganizationParams) WithOrganization(organization *mono_models.OrganizationEditable) *EditOrganizationParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the edit organization params
func (o *EditOrganizationParams) SetOrganization(organization *mono_models.OrganizationEditable) {
	o.Organization = organization
}

// WithOrganizationIdentifier adds the organizationIdentifier to the edit organization params
func (o *EditOrganizationParams) WithOrganizationIdentifier(organizationIdentifier string) *EditOrganizationParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the edit organization params
func (o *EditOrganizationParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *EditOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IdentifierType != nil {

		// query param identifierType
		var qrIdentifierType string
		if o.IdentifierType != nil {
			qrIdentifierType = *o.IdentifierType
		}
		qIdentifierType := qrIdentifierType
		if qIdentifierType != "" {
			if err := r.SetQueryParam("identifierType", qIdentifierType); err != nil {
				return err
			}
		}

	}

	if o.Organization != nil {
		if err := r.SetBodyParam(o.Organization); err != nil {
			return err
		}
	}

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
