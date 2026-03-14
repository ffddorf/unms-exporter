package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag/jsonutils"
)

// TimeSeriesData holds avg and max time series returned by the
// UISP v3 /devices/{id}/statistics endpoint.
type TimeSeriesData struct {
	Avg ListOfCoordinates `json:"avg,omitempty"`
	Max ListOfCoordinates `json:"max,omitempty"`
}

// DeviceStatistics device statistics
//
// swagger:model DeviceStatistics
type DeviceStatistics struct {
	Ping        *TimeSeriesData `json:"ping,omitempty"`
	CPU         *TimeSeriesData `json:"cpu,omitempty"`
	RAM         *TimeSeriesData `json:"ram,omitempty"`
	Temperature *TimeSeriesData `json:"temperature,omitempty"`
	Errors      *TimeSeriesData `json:"errors,omitempty"`
}

// Validate validates this device statistics
func (m *DeviceStatistics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device statistics based on the context it is used
func (m *DeviceStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return jsonutils.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatistics) UnmarshalBinary(b []byte) error {
	var res DeviceStatistics
	if err := jsonutils.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
