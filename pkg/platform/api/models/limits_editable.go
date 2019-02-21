// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LimitsEditable limits editable
// swagger:model LimitsEditable
type LimitsEditable struct {

	// users limit
	UsersLimit *int64 `json:"usersLimit,omitempty"`
}

// Validate validates this limits editable
func (m *LimitsEditable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LimitsEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LimitsEditable) UnmarshalBinary(b []byte) error {
	var res LimitsEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
