// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceFirmware device firmware
//
// swagger:model DeviceFirmware
type DeviceFirmware struct {

	// Is firmware compatible with UNMS
	// Required: true
	Compatible *bool `json:"compatible"`

	// Current firmware version.
	// Required: true
	Current *string `json:"current"`

	// Latest known firmware version.
	// Required: true
	Latest *string `json:"latest"`

	// semver
	Semver *Semver `json:"semver,omitempty"`
}

// Validate validates this device firmware
func (m *DeviceFirmware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFirmware) validateCompatible(formats strfmt.Registry) error {

	if err := validate.Required("compatible", "body", m.Compatible); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware) validateLatest(formats strfmt.Registry) error {

	if err := validate.Required("latest", "body", m.Latest); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFirmware) validateSemver(formats strfmt.Registry) error {
	if swag.IsZero(m.Semver) { // not required
		return nil
	}

	if m.Semver != nil {
		if err := m.Semver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semver")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device firmware based on the context it is used
func (m *DeviceFirmware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSemver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFirmware) contextValidateSemver(ctx context.Context, formats strfmt.Registry) error {

	if m.Semver != nil {

		if swag.IsZero(m.Semver) { // not required
			return nil
		}

		if err := m.Semver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("semver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("semver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceFirmware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceFirmware) UnmarshalBinary(b []byte) error {
	var res DeviceFirmware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
