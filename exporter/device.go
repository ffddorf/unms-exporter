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

		if e.extras.NeedStatistics() {
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

type LinkMetrics struct {
	UplinkCapacity      float64
	UplinkUtilization   float64
	DownlinkCapacity    float64
	DownlinkUtilization float64
}

// LinkMetricsWindow limits the data returned from the statistics
// endpoint from which we compute the average. The smallest interval
// allowed by UNMS is 1 hour, but we don't want to wait this long for
// anomalies to become visible.
const LinkMetricsWindow = 10 * time.Minute

func (dev *Device) LinkMetrics() *LinkMetrics {
	s := dev.Statistics
	if s == nil {
		return nil
	}

	max := s.Interval.End
	min := float64(time.UnixMilli(int64(max)).Add(-LinkMetricsWindow).UnixMilli())

	return &LinkMetrics{
		UplinkCapacity:      weightedMean(min, max, s.UplinkCapacity),
		UplinkUtilization:   weightedMean(min, max, s.UplinkUtilization),
		DownlinkCapacity:    weightedMean(min, max, s.DownlinkCapacity),
		DownlinkUtilization: weightedMean(min, max, s.DownlinkUtilization),
	}
}
