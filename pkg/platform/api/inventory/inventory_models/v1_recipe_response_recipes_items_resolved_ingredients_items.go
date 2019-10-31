// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResponseRecipesItemsResolvedIngredientsItems Resolved Ingredient
//
// An ingredient that is part of a recipe's resolved requirements
// swagger:model v1RecipeResponseRecipesItemsResolvedIngredientsItems
type V1RecipeResponseRecipesItemsResolvedIngredientsItems struct {

	// Alternative ingredient versions which can also satisfy the order's requirement. Each entry in the array is the ID of an ingredient version which could satisfy these requirements.
	Alternatives []strfmt.UUID `json:"alternatives"`

	// The custom build scripts for building this ingredient, if any
	BuildScripts []*V1RecipeResponseRecipesItemsResolvedIngredientsItemsBuildScriptsItems `json:"build_scripts"`

	// This dependencies in the recipe for this ingredient version. Each item contains an ingredient version UUID which maps to an ingredient version in this recipe.
	Dependencies []*V1RecipeResponseRecipesItemsResolvedIngredientsItemsDependenciesItems `json:"dependencies"`

	// ingredient
	// Required: true
	Ingredient *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredient `json:"ingredient"`

	// ingredient version
	// Required: true
	IngredientVersion *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersion `json:"ingredient_version"`

	// The patches to apply to this ingredient's source before building, if any
	Patches []*V1RecipeResponseRecipesItemsResolvedIngredientsItemsPatchesItems `json:"patches"`

	// The original requirement(s) in the order that were resolved to this ingredient version. This list will be empty if an ingredient was added to the recipe to fulfill a dependency of something else in the order.
	ResolvedRequirements []*V1RecipeResponseRecipesItemsResolvedIngredientsItemsResolvedRequirementsItems `json:"resolved_requirements"`
}

// Validate validates this v1 recipe response recipes items resolved ingredients items
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternatives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildScripts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolvedRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateAlternatives(formats strfmt.Registry) error {

	if swag.IsZero(m.Alternatives) { // not required
		return nil
	}

	for i := 0; i < len(m.Alternatives); i++ {

		if err := validate.FormatOf("alternatives"+"."+strconv.Itoa(i), "body", "uuid", m.Alternatives[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateBuildScripts(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildScripts) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildScripts); i++ {
		if swag.IsZero(m.BuildScripts[i]) { // not required
			continue
		}

		if m.BuildScripts[i] != nil {
			if err := m.BuildScripts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_scripts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateIngredient(formats strfmt.Registry) error {

	if err := validate.Required("ingredient", "body", m.Ingredient); err != nil {
		return err
	}

	if m.Ingredient != nil {
		if err := m.Ingredient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingredient")
			}
			return err
		}
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateIngredientVersion(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version", "body", m.IngredientVersion); err != nil {
		return err
	}

	if m.IngredientVersion != nil {
		if err := m.IngredientVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingredient_version")
			}
			return err
		}
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validatePatches(formats strfmt.Registry) error {

	if swag.IsZero(m.Patches) { // not required
		return nil
	}

	for i := 0; i < len(m.Patches); i++ {
		if swag.IsZero(m.Patches[i]) { // not required
			continue
		}

		if m.Patches[i] != nil {
			if err := m.Patches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) validateResolvedRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolvedRequirements) { // not required
		return nil
	}

	for i := 0; i < len(m.ResolvedRequirements); i++ {
		if swag.IsZero(m.ResolvedRequirements[i]) { // not required
			continue
		}

		if m.ResolvedRequirements[i] != nil {
			if err := m.ResolvedRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolved_requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsResolvedIngredientsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
