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
}

var pingMetrics = map[string]metricSpec{
	"ping_loss_ratio":                newSpec("Ping packet loss ratio", nil),
	"ping_rtt_best_seconds":          newSpec("Best ping round trip time in seconds", nil),
	"ping_rtt_mean_seconds":          newSpec("Mean ping round trip time in seconds", nil),
	"ping_rtt_worst_seconds":         newSpec("Worst ping round trip time in seconds", nil),
	"ping_rtt_std_deviation_seconds": newSpec("Standard deviation for ping round trip time in seconds", nil),
}

func (e *Exporter) SetExtras(extras []string) error {
	e.extras = ExtraMetrics{} // reset all values
	for _, x := range extras {
		switch x {
		case "ping":
			e.extras.Ping = true
		default:
			return fmt.Errorf("unknown extra metric: %q", x)
		}
	}

	for name, spec := range pingMetrics {
		if _, exists := e.metrics[name]; !exists && e.extras.Ping {
			e.metrics[name] = spec.intoDesc(name)
		} else if !e.extras.Ping {
			delete(e.metrics, name)
		}
	}

	return nil
}
