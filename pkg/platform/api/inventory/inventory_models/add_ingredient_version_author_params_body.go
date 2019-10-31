// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddIngredientVersionAuthorParamsBody add ingredient version author params body
// swagger:model addIngredientVersionAuthorParamsBody
type AddIngredientVersionAuthorParamsBody struct {

	// The ID of the author of this ingredient version
	// Required: true
	// Format: uuid
	AuthorID *strfmt.UUID `json:"author_id"`
}

// Validate validates this add ingredient version author params body
func (m *AddIngredientVersionAuthorParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddIngredientVersionAuthorParamsBody) validateAuthorID(formats strfmt.Registry) error {

	if err := validate.Required("author_id", "body", m.AuthorID); err != nil {
		return err
	}

	if err := validate.FormatOf("author_id", "body", "uuid", m.AuthorID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddIngredientVersionAuthorParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddIngredientVersionAuthorParamsBody) UnmarshalBinary(b []byte) error {
	var res AddIngredientVersionAuthorParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
