package exporter

import (
	"context"
	"time"

	"github.com/ffddorf/unms-exporter/client/devices"
	"github.com/ffddorf/unms-exporter/models"
)

var defaultWithInterfaces = true

type Device struct {
	Statistics *models.DeviceStatistics
	*models.DeviceStatusOverview
}

func (e *Exporter) fetchDeviceData() ([]Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	params := &devices.GetDevicesParams{
		WithInterfaces: &defaultWithInterfaces,
		Context:        ctx,
	}
	devicesResponse, err := e.api.Devices.GetDevices(params)
	if err != nil {
		return nil, err
	}

	data := make([]Device, 0, len(devicesResponse.Payload))
	for _, overview := range devicesResponse.Payload {
		if overview.Identification == nil {
			continue
		}
		dev := Device{nil, overview}

		if id := derefOrEmpty(overview.Identification.ID); id != "" {
			params := &devices.GetDevicesIDStatisticsParams{
				ID:       id,
				Interval: "hour", // smallest interval possible
				Context:  ctx,
			}
			statisticsResponse, err := e.api.Devices.GetDevicesIDStatistics(params)
			if err != nil {
				return nil, err
			}
			dev.Statistics = statisticsResponse.Payload
		}
		data = append(data, dev)
	}

	return data, nil
}

func (dev *Device) PingMetrics() *PingMetrics {
	if dev.Statistics == nil || len(dev.Statistics.Ping) == 0 {
		return nil
	}

	m := NewHistory(len(dev.Statistics.Ping))
	for _, xy := range dev.Statistics.Ping {
		if xy == nil {
			m.Add(0, true)
			continue
		}

		rtt := time.Duration(xy.Y * float64(time.Millisecond))
		m.Add(rtt, false)
	}

	return m.Compute()
}
