// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIngredientVersionParams creates a new GetIngredientVersionParams object
// with the default values initialized.
func NewGetIngredientVersionParams() *GetIngredientVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetIngredientVersionParams{
		AllowUnstable: &allowUnstableDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIngredientVersionParamsWithTimeout creates a new GetIngredientVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIngredientVersionParamsWithTimeout(timeout time.Duration) *GetIngredientVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetIngredientVersionParams{
		AllowUnstable: &allowUnstableDefault,

		timeout: timeout,
	}
}

// NewGetIngredientVersionParamsWithContext creates a new GetIngredientVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIngredientVersionParamsWithContext(ctx context.Context) *GetIngredientVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetIngredientVersionParams{
		AllowUnstable: &allowUnstableDefault,

		Context: ctx,
	}
}

// NewGetIngredientVersionParamsWithHTTPClient creates a new GetIngredientVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIngredientVersionParamsWithHTTPClient(client *http.Client) *GetIngredientVersionParams {
	var (
		allowUnstableDefault = bool(false)
	)
	return &GetIngredientVersionParams{
		AllowUnstable: &allowUnstableDefault,
		HTTPClient:    client,
	}
}

/*GetIngredientVersionParams contains all the parameters to send to the API endpoint
for the get ingredient version operation typically these are written to a http.Request
*/
type GetIngredientVersionParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*IngredientID*/
	IngredientID strfmt.UUID
	/*IngredientVersionID*/
	IngredientVersionID strfmt.UUID
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ingredient version params
func (o *GetIngredientVersionParams) WithTimeout(timeout time.Duration) *GetIngredientVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ingredient version params
func (o *GetIngredientVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ingredient version params
func (o *GetIngredientVersionParams) WithContext(ctx context.Context) *GetIngredientVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ingredient version params
func (o *GetIngredientVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ingredient version params
func (o *GetIngredientVersionParams) WithHTTPClient(client *http.Client) *GetIngredientVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ingredient version params
func (o *GetIngredientVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get ingredient version params
func (o *GetIngredientVersionParams) WithAllowUnstable(allowUnstable *bool) *GetIngredientVersionParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get ingredient version params
func (o *GetIngredientVersionParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithIngredientID adds the ingredientID to the get ingredient version params
func (o *GetIngredientVersionParams) WithIngredientID(ingredientID strfmt.UUID) *GetIngredientVersionParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the get ingredient version params
func (o *GetIngredientVersionParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersionID adds the ingredientVersionID to the get ingredient version params
func (o *GetIngredientVersionParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *GetIngredientVersionParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the get ingredient version params
func (o *GetIngredientVersionParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WithStateAt adds the stateAt to the get ingredient version params
func (o *GetIngredientVersionParams) WithStateAt(stateAt *strfmt.DateTime) *GetIngredientVersionParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get ingredient version params
func (o *GetIngredientVersionParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetIngredientVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool
		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {
			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}

	}

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}

	// path param ingredient_version_id
	if err := r.SetPathParam("ingredient_version_id", o.IngredientVersionID.String()); err != nil {
		return err
	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime
		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {
			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
