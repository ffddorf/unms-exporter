package exporter

import "fmt"

// ExtraMetrics is used to instruct Exporter to fetch additional metrics,
// not captured by the UNMS API /devices endpoint.
//
// These metrics may require extra HTTP requests, usually one per device,
// so it might not be desirable to have them included by default.
//
// Expect this type to gain additional fields in the future. Currently,
// enabling ping metrics will also fetch (but discard) additional metrics
// from the /devices/{id}/statitics API endpoint, like temperature, link
// capacity, and many more values.
type ExtraMetrics struct {
	Ping bool
	Link bool
}

// NeedStatistics is true, if any field is true that require data from
// the same device statistics endpoint.
func (x ExtraMetrics) NeedStatistics() bool {
	return x.Ping || x.Link
}

var pingMetrics = map[string]metricSpec{
	"ping_loss_ratio":                newSpec("Ping packet loss ratio", nil),
	"ping_rtt_best_seconds":          newSpec("Best ping round trip time in seconds", nil),
	"ping_rtt_mean_seconds":          newSpec("Mean ping round trip time in seconds", nil),
	"ping_rtt_worst_seconds":         newSpec("Worst ping round trip time in seconds", nil),
	"ping_rtt_std_deviation_seconds": newSpec("Standard deviation for ping round trip time in seconds", nil),
}

var linkMetrics = map[string]metricSpec{
	"uplink_capacity_rate":       newSpec("Uplink capacity in Bit/s", nil),
	"uplink_utilization_ratio":   newSpec("Uplink utilization ratio", nil),
	"downlink_capacity_rate":     newSpec("Downlink capacity in Bit/s", nil),
	"downlink_utilization_ratio": newSpec("Downlink utilization ratio", nil),
}

func (e *Exporter) SetExtras(extras []string) error {
	e.extras = ExtraMetrics{} // reset all values
	for _, x := range extras {
		switch x {
		case "ping":
			e.extras.Ping = true
		case "link":
			e.extras.Link = true
		default:
			return fmt.Errorf("unknown extra metric: %q", x)
		}
	}

	e.configureMetrics(e.extras.Ping, pingMetrics)
	e.configureMetrics(e.extras.Link, linkMetrics)
	return nil
}

func (e *Exporter) configureMetrics(enable bool, metrics map[string]metricSpec) {
	for name, spec := range metrics {
		if _, exists := e.metrics[name]; !exists && enable {
			e.metrics[name] = spec.intoDesc(name)
		} else if !enable {
			delete(e.metrics, name)
		}
	}
}
