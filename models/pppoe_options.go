// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PppoeOptions pppoe options
//
// swagger:model pppoeOptions
type PppoeOptions struct {

	// account
	Account string `json:"account,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// parent
	Parent string `json:"parent,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this pppoe options
func (m *PppoeOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pppoe options based on context it is used
func (m *PppoeOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PppoeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PppoeOptions) UnmarshalBinary(b []byte) error {
	var res PppoeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
