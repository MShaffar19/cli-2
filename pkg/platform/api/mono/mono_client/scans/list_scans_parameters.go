// Code generated by go-swagger; DO NOT EDIT.

package scans

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
)

// NewListScansParams creates a new ListScansParams object
// with the default values initialized.
func NewListScansParams() *ListScansParams {
	var ()
	return &ListScansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListScansParamsWithTimeout creates a new ListScansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListScansParamsWithTimeout(timeout time.Duration) *ListScansParams {
	var ()
	return &ListScansParams{

		timeout: timeout,
	}
}

// NewListScansParamsWithContext creates a new ListScansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListScansParamsWithContext(ctx context.Context) *ListScansParams {
	var ()
	return &ListScansParams{

		Context: ctx,
	}
}

// NewListScansParamsWithHTTPClient creates a new ListScansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListScansParamsWithHTTPClient(client *http.Client) *ListScansParams {
	var ()
	return &ListScansParams{
		HTTPClient: client,
	}
}

/*ListScansParams contains all the parameters to send to the API endpoint
for the list scans operation typically these are written to a http.Request
*/
type ListScansParams struct {

	/*Limit
	  Limit output to this many rows

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list scans params
func (o *ListScansParams) WithTimeout(timeout time.Duration) *ListScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list scans params
func (o *ListScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list scans params
func (o *ListScansParams) WithContext(ctx context.Context) *ListScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list scans params
func (o *ListScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list scans params
func (o *ListScansParams) WithHTTPClient(client *http.Client) *ListScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list scans params
func (o *ListScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list scans params
func (o *ListScansParams) WithLimit(limit *int64) *ListScansParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list scans params
func (o *ListScansParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ListScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
