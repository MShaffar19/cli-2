// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildRequestRecipePlatformOperatingSystemAllOf0 v1 build request recipe platform operating system all of0
// swagger:model v1BuildRequestRecipePlatformOperatingSystemAllOf0
type V1BuildRequestRecipePlatformOperatingSystemAllOf0 struct {

	// links
	// Required: true
	Links *V1BuildRequestRecipePlatformOperatingSystemAllOf0Links `json:"links"`

	// operating system id
	// Required: true
	// Format: uuid
	OperatingSystemID *strfmt.UUID `json:"operating_system_id"`
}

// Validate validates this v1 build request recipe platform operating system all of0
func (m *V1BuildRequestRecipePlatformOperatingSystemAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildRequestRecipePlatformOperatingSystemAllOf0) validateLinks(formats strfmt.Registry) error {

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

func (m *V1BuildRequestRecipePlatformOperatingSystemAllOf0) validateOperatingSystemID(formats strfmt.Registry) error {

	if err := validate.Required("operating_system_id", "body", m.OperatingSystemID); err != nil {
		return err
	}

	if err := validate.FormatOf("operating_system_id", "body", "uuid", m.OperatingSystemID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemAllOf0) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformOperatingSystemAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
