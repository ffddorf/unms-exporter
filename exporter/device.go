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

func (e *Exporter) fetchDeviceData(ctx context.Context) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
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

		if e.extras.Ping {
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
		}
		data = append(data, dev)
	}

	return data, nil
}

func (dev *Device) PingMetrics() *PingMetrics {
	if dev.Statistics == nil || dev.Statistics.Ping == nil || len(dev.Statistics.Ping.Avg) == 0 {
		return nil
	}

	coords := dev.Statistics.Ping.Avg
	m := NewHistory(len(coords))
	for _, xy := range coords {
		if xy == nil {
			m.Add(0, true)
			continue
		}

		// UISP 3.x returns RTT in seconds, convert to Duration
		rtt := time.Duration(xy.Y * float64(time.Second))
		m.Add(rtt, false)
	}

	return m.Compute()
}
