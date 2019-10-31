// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1 v1 build request recipe resolved ingredients items ingredient version all of3 all of0 dependency sets items dependencies items conditions items all of1
// swagger:model v1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1
type V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1 struct {

	// Whatever text or data structure we parsed to generate this condition
	OriginalCondition string `json:"original_condition,omitempty"`
}

// Validate validates this v1 build request recipe resolved ingredients items ingredient version all of3 all of0 dependency sets items dependencies items conditions items all of1
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItemsDependenciesItemsConditionsItemsAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
