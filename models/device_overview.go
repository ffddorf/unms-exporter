// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceOverview Read-only basic device/client overview attributes.
//
// swagger:model DeviceOverview
type DeviceOverview struct {

	// antenna
	Antenna *Antenna `json:"antenna,omitempty"`

	// battery capacity
	BatteryCapacity float64 `json:"batteryCapacity,omitempty"`

	// battery time
	BatteryTime float64 `json:"batteryTime,omitempty"`

	// Nullable property in milliamperes.
	BiasCurrent float64 `json:"biasCurrent,omitempty"`

	// TRUE if device can be upgraded.
	CanUpgrade bool `json:"canUpgrade,omitempty"`

	// channel width
	ChannelWidth float64 `json:"channelWidth,omitempty"`

	// Current cpu load.
	CPU float64 `json:"cpu,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Nullable property in meters.
	Distance float64 `json:"distance,omitempty"`

	// downlink capacity
	DownlinkCapacity int64 `json:"downlinkCapacity,omitempty"`

	// Nullable prop; current frequency (only for airmax devices).
	Frequency float64 `json:"frequency,omitempty"`

	// TRUE if device is in location mode.
	IsLocateRunning bool `json:"isLocateRunning,omitempty"`

	// Last seen timestamp in ISO format.
	// Example: 2018-11-14T15:20:32.004Z
	// Format: date-time
	LastSeen strfmt.DateTime `json:"lastSeen,omitempty"`

	// link score
	LinkScore *LinkScore `json:"linkScore,omitempty"`

	// power status
	PowerStatus float64 `json:"powerStatus,omitempty"`

	// Current memory usage.
	RAM float64 `json:"ram,omitempty"`

	// Theoretical max remote signal level.
	// Example: -55
	RemoteSignalMax float64 `json:"remoteSignalMax,omitempty"`

	// TRUE if device is running on battery
	RunningOnBattery bool `json:"runningOnBattery,omitempty"`

	// Nullable prop; current signal level (only for airmax devices), for example -55 dBm.
	// Example: -55
	Signal float64 `json:"signal,omitempty"`

	// Theoretical max local signal level.
	// Example: -55
	SignalMax float64 `json:"signalMax,omitempty"`

	// Count of stations (only for airmax and aircube).
	StationsCount float64 `json:"stationsCount,omitempty"`

	// Read-only value generated by UNMS.
	Status string `json:"status,omitempty"`

	// temperature
	Temperature float64 `json:"temperature,omitempty"`

	// theoretical downlink capacity
	TheoreticalDownlinkCapacity int64 `json:"theoreticalDownlinkCapacity,omitempty"`

	// theoretical max downlink capacity
	TheoreticalMaxDownlinkCapacity int64 `json:"theoreticalMaxDownlinkCapacity,omitempty"`

	// theoretical max uplink capacity
	TheoreticalMaxUplinkCapacity int64 `json:"theoreticalMaxUplinkCapacity,omitempty"`

	// theoretical uplink capacity
	TheoreticalUplinkCapacity int64 `json:"theoreticalUplinkCapacity,omitempty"`

	// transmit power
	TransmitPower float64 `json:"transmitPower,omitempty"`

	// uplink capacity
	UplinkCapacity int64 `json:"uplinkCapacity,omitempty"`

	// Uptime in seconds.
	Uptime float64 `json:"uptime,omitempty"`

	// System input voltage in V.
	Voltage float64 `json:"voltage,omitempty"`

	// wireless active interface ids
	WirelessActiveInterfaceIds WirelessActiveInterfaceIds `json:"wirelessActiveInterfaceIds,omitempty"`

	// wireless mode
	// Enum: [ap ap-ptp ap-ptmp ap-ptmp-airmax ap-ptmp-airmax-mixed ap-ptmp-airmax-ac sta sta-ptp sta-ptmp aprepeater repeater mesh]
	WirelessMode string `json:"wirelessMode,omitempty"`
}

// Validate validates this device overview
func (m *DeviceOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntenna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWirelessActiveInterfaceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWirelessMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceOverview) validateAntenna(formats strfmt.Registry) error {
	if swag.IsZero(m.Antenna) { // not required
		return nil
	}

	if m.Antenna != nil {
		if err := m.Antenna.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antenna")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceOverview) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceOverview) validateLastSeen(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceOverview) validateLinkScore(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkScore) { // not required
		return nil
	}

	if m.LinkScore != nil {
		if err := m.LinkScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linkScore")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceOverview) validateWirelessActiveInterfaceIds(formats strfmt.Registry) error {
	if swag.IsZero(m.WirelessActiveInterfaceIds) { // not required
		return nil
	}

	if err := m.WirelessActiveInterfaceIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wirelessActiveInterfaceIds")
		}
		return err
	}

	return nil
}

var deviceOverviewTypeWirelessModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ap","ap-ptp","ap-ptmp","ap-ptmp-airmax","ap-ptmp-airmax-mixed","ap-ptmp-airmax-ac","sta","sta-ptp","sta-ptmp","aprepeater","repeater","mesh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceOverviewTypeWirelessModePropEnum = append(deviceOverviewTypeWirelessModePropEnum, v)
	}
}

const (

	// DeviceOverviewWirelessModeAp captures enum value "ap"
	DeviceOverviewWirelessModeAp string = "ap"

	// DeviceOverviewWirelessModeApMinusptp captures enum value "ap-ptp"
	DeviceOverviewWirelessModeApMinusptp string = "ap-ptp"

	// DeviceOverviewWirelessModeApMinusptmp captures enum value "ap-ptmp"
	DeviceOverviewWirelessModeApMinusptmp string = "ap-ptmp"

	// DeviceOverviewWirelessModeApMinusptmpMinusairmax captures enum value "ap-ptmp-airmax"
	DeviceOverviewWirelessModeApMinusptmpMinusairmax string = "ap-ptmp-airmax"

	// DeviceOverviewWirelessModeApMinusptmpMinusairmaxMinusmixed captures enum value "ap-ptmp-airmax-mixed"
	DeviceOverviewWirelessModeApMinusptmpMinusairmaxMinusmixed string = "ap-ptmp-airmax-mixed"

	// DeviceOverviewWirelessModeApMinusptmpMinusairmaxMinusac captures enum value "ap-ptmp-airmax-ac"
	DeviceOverviewWirelessModeApMinusptmpMinusairmaxMinusac string = "ap-ptmp-airmax-ac"

	// DeviceOverviewWirelessModeSta captures enum value "sta"
	DeviceOverviewWirelessModeSta string = "sta"

	// DeviceOverviewWirelessModeStaMinusptp captures enum value "sta-ptp"
	DeviceOverviewWirelessModeStaMinusptp string = "sta-ptp"

	// DeviceOverviewWirelessModeStaMinusptmp captures enum value "sta-ptmp"
	DeviceOverviewWirelessModeStaMinusptmp string = "sta-ptmp"

	// DeviceOverviewWirelessModeAprepeater captures enum value "aprepeater"
	DeviceOverviewWirelessModeAprepeater string = "aprepeater"

	// DeviceOverviewWirelessModeRepeater captures enum value "repeater"
	DeviceOverviewWirelessModeRepeater string = "repeater"

	// DeviceOverviewWirelessModeMesh captures enum value "mesh"
	DeviceOverviewWirelessModeMesh string = "mesh"
)

// prop value enum
func (m *DeviceOverview) validateWirelessModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceOverviewTypeWirelessModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceOverview) validateWirelessMode(formats strfmt.Registry) error {
	if swag.IsZero(m.WirelessMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateWirelessModeEnum("wirelessMode", "body", m.WirelessMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device overview based on the context it is used
func (m *DeviceOverview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntenna(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWirelessActiveInterfaceIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceOverview) contextValidateAntenna(ctx context.Context, formats strfmt.Registry) error {

	if m.Antenna != nil {
		if err := m.Antenna.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antenna")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceOverview) contextValidateLinkScore(ctx context.Context, formats strfmt.Registry) error {

	if m.LinkScore != nil {
		if err := m.LinkScore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linkScore")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceOverview) contextValidateWirelessActiveInterfaceIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.WirelessActiveInterfaceIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wirelessActiveInterfaceIds")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceOverview) UnmarshalBinary(b []byte) error {
	var res DeviceOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
