// Code generated by go-swagger; DO NOT EDIT.

package secrets_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSecretChange user secret change
// swagger:model UserSecretChange
type UserSecretChange struct {

	// is user
	// Required: true
	IsUser *bool `json:"is_user"`

	// name
	// Required: true
	Name *string `json:"name"`

	// project id
	// Format: uuid
	ProjectID strfmt.UUID `json:"project_id,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this user secret change
func (m *UserSecretChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSecretChange) validateIsUser(formats strfmt.Registry) error {

	if err := validate.Required("is_user", "body", m.IsUser); err != nil {
		return err
	}

	return nil
}

func (m *UserSecretChange) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserSecretChange) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("project_id", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserSecretChange) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSecretChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSecretChange) UnmarshalBinary(b []byte) error {
	var res UserSecretChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
