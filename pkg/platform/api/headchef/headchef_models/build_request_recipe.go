// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildRequestRecipe Recipe
//
// A recipe contains the exact ingredient versions, including dependencies, needed to build a project for a single platform.
// swagger:model buildRequestRecipe
type BuildRequestRecipe struct {

	// List of build options which are selected for this recipe.
	BuildOptions []*BuildRequestRecipeBuildOptionsItems `json:"build_options"`

	// The name of the image that will be used to build this recipe.
	Image string `json:"image,omitempty"`

	// The type of image that will be used to build this recipe.
	// Enum: [Docker WindowsInstance]
	ImageType string `json:"image_type,omitempty"`

	// Platform ID for the recipe.
	// Required: true
	// Format: uuid
	PlatformID *strfmt.UUID `json:"platform_id"`

	// Recipe UUID
	// Required: true
	// Format: uuid
	RecipeID *strfmt.UUID `json:"recipe_id"`

	// Resolved list of requirements. For each requirement in the original order there will be a corresponding ingredient version.
	// Required: true
	ResolvedRequirements []*BuildRequestRecipeResolvedRequirementsItems `json:"resolved_requirements"`
}

// Validate validates this build request recipe
func (m *BuildRequestRecipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeID(formats); err != nil {
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

func (m *BuildRequestRecipe) validateBuildOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildOptions); i++ {
		if swag.IsZero(m.BuildOptions[i]) { // not required
			continue
		}

		if m.BuildOptions[i] != nil {
			if err := m.BuildOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var buildRequestRecipeTypeImageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Docker","WindowsInstance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildRequestRecipeTypeImageTypePropEnum = append(buildRequestRecipeTypeImageTypePropEnum, v)
	}
}

const (

	// BuildRequestRecipeImageTypeDocker captures enum value "Docker"
	BuildRequestRecipeImageTypeDocker string = "Docker"

	// BuildRequestRecipeImageTypeWindowsInstance captures enum value "WindowsInstance"
	BuildRequestRecipeImageTypeWindowsInstance string = "WindowsInstance"
)

// prop value enum
func (m *BuildRequestRecipe) validateImageTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildRequestRecipeTypeImageTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildRequestRecipe) validateImageType(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateImageTypeEnum("image_type", "body", m.ImageType); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipe) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platform_id", "body", m.PlatformID); err != nil {
		return err
	}

	if err := validate.FormatOf("platform_id", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipe) validateRecipeID(formats strfmt.Registry) error {

	if err := validate.Required("recipe_id", "body", m.RecipeID); err != nil {
		return err
	}

	if err := validate.FormatOf("recipe_id", "body", "uuid", m.RecipeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipe) validateResolvedRequirements(formats strfmt.Registry) error {

	if err := validate.Required("resolved_requirements", "body", m.ResolvedRequirements); err != nil {
		return err
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
func (m *BuildRequestRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildRequestRecipe) UnmarshalBinary(b []byte) error {
	var res BuildRequestRecipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
