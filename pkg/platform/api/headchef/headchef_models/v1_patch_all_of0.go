// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PatchAllOf0 v1 patch all of0
//
// swagger:model v1PatchAllOf0
type V1PatchAllOf0 struct {

	// creation timestamp
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

	// links
	// Required: true
	Links *V1SubSchemaSelfLink `json:"links"`

	// patch id
	// Required: true
	// Format: uuid
	PatchID *strfmt.UUID `json:"patch_id"`
}

// Validate validates this v1 patch all of0
func (m *V1PatchAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatchID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PatchAllOf0) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PatchAllOf0) validateLinks(formats strfmt.Registry) error {

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

func (m *V1PatchAllOf0) validatePatchID(formats strfmt.Registry) error {

	if err := validate.Required("patch_id", "body", m.PatchID); err != nil {
		return err
	}

	if err := validate.FormatOf("patch_id", "body", "uuid", m.PatchID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PatchAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PatchAllOf0) UnmarshalBinary(b []byte) error {
	var res V1PatchAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
