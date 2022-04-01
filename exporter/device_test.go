package exporter

import (
	"testing"
	"time"

	"github.com/ffddorf/unms-exporter/models"
)

const (
	ms = time.Millisecond
	µs = time.Microsecond //nolint:asciicheck
)

type metricExpectation map[string]struct {
	actual    interface{}
	satisfied bool
}

func comparePingMetrics(t *testing.T, expectations metricExpectation, actual *PingMetrics) {
	t.Helper()

	anyFailure := false
	for field, expectation := range expectations {
		if !expectation.satisfied {
			anyFailure = true
			t.Errorf("unexpected value for field %q: %v", field, expectation.actual)
		}
	}
	if anyFailure {
		t.FailNow()
	}
}

func TestDevice_PingMetrics_connected(t *testing.T) {
	t.Parallel()

	subject := Device{
		Statistics: &models.DeviceStatistics{
			Ping: models.ListOfCoordinates{{Y: 5}, {Y: 10}, {Y: 25}, {Y: 15}, {Y: 1}}, // x values are ignored
		},
	}

	actual := subject.PingMetrics()
	if actual == nil {
		t.Error("expected PingMetrics() to return somthing, got nil")
	}

	comparePingMetrics(t, metricExpectation{
		"packets sent": {actual.PacketsSent, actual.PacketsSent == 5},
		"packets lost": {actual.PacketsLost, actual.PacketsLost == 0},
		"rtt best":     {actual.Best, actual.Best == 1*ms},
		"rtt worst":    {actual.Worst, actual.Worst == 25*ms},
		"rtt median":   {actual.Median, actual.Median == 10*ms},
		"rtt meain":    {actual.Mean, actual.Mean == 11200*µs},                              // 11.2ms
		"rtt std dev":  {actual.StdDev, 8350*µs < actual.StdDev && actual.StdDev < 8360*µs}, // ~8.352245ms
	}, actual)
}

func TestDevice_PingMetrics_missingPackets(t *testing.T) {
	t.Parallel()

	subject := Device{
		Statistics: &models.DeviceStatistics{
			Ping: models.ListOfCoordinates{nil, {Y: 100}, {Y: 250}, nil, {Y: 120}},
		},
	}

	actual := subject.PingMetrics()
	if actual == nil {
		t.Error("expected PingMetrics() to return somthing, got nil")
	}

	comparePingMetrics(t, metricExpectation{
		"packets sent": {actual.PacketsSent, actual.PacketsSent == 5},
		"packets lost": {actual.PacketsLost, actual.PacketsLost == 2},
		"rtt best":     {actual.Best, actual.Best == 100*ms},
		"rtt worst":    {actual.Worst, actual.Worst == 250*ms},
		"rtt median":   {actual.Median, actual.Median == 120*ms},
		"rtt meain":    {actual.Mean, 156666*µs < actual.Mean && actual.Mean < 156667*µs},     // 156.66666ms
		"rtt std dev":  {actual.StdDev, 66499*µs < actual.StdDev && actual.StdDev < 66500*µs}, // ~66.499791ms
	}, actual)
}

func TestDevice_PingMetrics_disconnected(t *testing.T) {
	t.Parallel()

	if actual := (&Device{}).PingMetrics(); actual != nil {
		t.Errorf("expected PingMetrics() to return nil, got %+v", actual)
	}
}
