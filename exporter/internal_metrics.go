package exporter

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type internalMetrics struct {
	success prom.Counter
	errors  prom.Counter
}

func newInternalMetrics() internalMetrics {
	success := prom.NewCounter(prom.CounterOpts{
		Namespace: namespace,
		Subsystem: "collector",
		Name:      "success",
		Help:      "Indicating if the scrape was successful",
	})
	errors := prom.NewCounter(prom.CounterOpts{
		Namespace: namespace,
		Subsystem: "collector",
		Name:      "errors",
		Help:      "Errors encountered while exporting metrics",
	})

	return internalMetrics{success, errors}
}

func (m *internalMetrics) Describe(out chan<- *prom.Desc) {
	m.success.Describe(out)
	m.errors.Describe(out)
}

func (m *internalMetrics) Collect(out chan<- prom.Metric) {
	m.success.Collect(out)
	m.errors.Collect(out)
}
