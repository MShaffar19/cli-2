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

// V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0 v1 recipe response recipes items platform gpu architecture all of0
// swagger:model v1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0
type V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0 struct {

	// gpu architecture id
	// Required: true
	// Format: uuid
	GpuArchitectureID *strfmt.UUID `json:"gpu_architecture_id"`

	// links
	// Required: true
	Links *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0Links `json:"links"`
}

// Validate validates this v1 recipe response recipes items platform gpu architecture all of0
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGpuArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0) validateGpuArchitectureID(formats strfmt.Registry) error {

	if err := validate.Required("gpu_architecture_id", "body", m.GpuArchitectureID); err != nil {
		return err
	}

	if err := validate.FormatOf("gpu_architecture_id", "body", "uuid", m.GpuArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
