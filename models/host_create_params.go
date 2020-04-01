// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostCreateParams host create params
//
// swagger:model host-create-params
type HostCreateParams struct {

	// host Id
	// Required: true
	// Format: uuid
	HostID *strfmt.UUID `json:"hostId"`
}

// Validate validates this host create params
func (m *HostCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreateParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("hostId", "body", m.HostID); err != nil {
		return err
	}

	if err := validate.FormatOf("hostId", "body", "uuid", m.HostID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostCreateParams) UnmarshalBinary(b []byte) error {
	var res HostCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}