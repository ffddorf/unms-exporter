package exporter

import (
	"context"

	"github.com/ffddorf/unms-exporter/client/devices"
	"github.com/ffddorf/unms-exporter/models"
)

var defaultWithInterfaces = true

func (e *Exporter) fetchDeviceData(ctx context.Context) ([]*models.DeviceStatusOverview, error) {
	params := &devices.GetDevicesParams{
		WithInterfaces: &defaultWithInterfaces,
		Context:        ctx,
	}
	devicesResponse, err := e.api.Devices.GetDevices(params)
	if err != nil {
		return nil, err
	}

	data := make([]*models.DeviceStatusOverview, 0, len(devicesResponse.Payload))
	for _, overview := range devicesResponse.Payload {
		if overview.Identification == nil {
			continue
		}
		data = append(data, overview)
	}

	return data, nil
}
