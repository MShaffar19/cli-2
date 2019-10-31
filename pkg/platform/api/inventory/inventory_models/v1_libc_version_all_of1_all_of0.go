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

// V1LibcVersionAllOf1AllOf0 Version Info
//
// Properties shared by all versioned resources
// swagger:model v1LibcVersionAllOf1AllOf0
type V1LibcVersionAllOf1AllOf0 struct {

	// An array of decimal values representing all segments of a version, ordered from most to least significant. How a version string is rendered into a list of decimals will vary depending on the format of the source string and is therefore left up to the caller, but it must be done consistently across all versions of the same resource for sorting to work properly. This is represented as a string to avoid losing precision when converting to a floating point number.
	// Required: true
	// Min Length: 1
	SortableVersion []string `json:"sortable_version"`

	// The canonical version string for the resource. Should be as specific as possible (e.g. '10.9.6' of macOS instead of just '10.9'). May contain non-numeric version segments and other formatting characters if necessary.
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this v1 libc version all of1 all of0
func (m *V1LibcVersionAllOf1AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortableVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LibcVersionAllOf1AllOf0) validateSortableVersion(formats strfmt.Registry) error {

	if err := validate.Required("sortable_version", "body", m.SortableVersion); err != nil {
		return err
	}

	for i := 0; i < len(m.SortableVersion); i++ {

		if err := validate.MinLength("sortable_version"+"."+strconv.Itoa(i), "body", string(m.SortableVersion[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1LibcVersionAllOf1AllOf0) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LibcVersionAllOf1AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibcVersionAllOf1AllOf0) UnmarshalBinary(b []byte) error {
	var res V1LibcVersionAllOf1AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
