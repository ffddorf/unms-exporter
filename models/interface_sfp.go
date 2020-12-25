// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterfaceSfp interface sfp
//
// swagger:model interfaceSfp
type InterfaceSfp struct {

	// include vlans
	IncludeVlans interface{} `json:"includeVlans,omitempty"`

	// max speed
	MaxSpeed int64 `json:"maxSpeed,omitempty"`

	// olt
	Olt bool `json:"olt,omitempty"`

	// part
	Part string `json:"part,omitempty"`

	// present
	Present bool `json:"present,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// vlan native
	VlanNative interface{} `json:"vlanNative,omitempty"`
}

// Validate validates this interface sfp
func (m *InterfaceSfp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this interface sfp based on context it is used
func (m *InterfaceSfp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceSfp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceSfp) UnmarshalBinary(b []byte) error {
	var res InterfaceSfp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
