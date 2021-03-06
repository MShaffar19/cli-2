// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1PlatformCPUArchitecture CPU Architecture
//
// The full CPU architecture data model
// swagger:model v1PlatformCpuArchitecture
type V1PlatformCPUArchitecture struct {
	V1PlatformCPUArchitectureAllOf0

	V1PlatformCPUArchitectureAllOf1

	V1PlatformCPUArchitectureAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1PlatformCPUArchitecture) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1PlatformCPUArchitectureAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1PlatformCPUArchitectureAllOf0 = aO0

	// AO1
	var aO1 V1PlatformCPUArchitectureAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1PlatformCPUArchitectureAllOf1 = aO1

	// AO2
	var aO2 V1PlatformCPUArchitectureAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1PlatformCPUArchitectureAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1PlatformCPUArchitecture) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1PlatformCPUArchitectureAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1PlatformCPUArchitectureAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1PlatformCPUArchitectureAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 platform Cpu architecture
func (m *V1PlatformCPUArchitecture) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1PlatformCPUArchitectureAllOf0
	if err := m.V1PlatformCPUArchitectureAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformCPUArchitectureAllOf1
	if err := m.V1PlatformCPUArchitectureAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformCPUArchitectureAllOf2
	if err := m.V1PlatformCPUArchitectureAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1PlatformCPUArchitecture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlatformCPUArchitecture) UnmarshalBinary(b []byte) error {
	var res V1PlatformCPUArchitecture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
