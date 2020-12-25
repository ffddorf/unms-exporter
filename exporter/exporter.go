package exporter

import (
	"context"
	"fmt"
	"time"

	"github.com/ffddorf/unms-exporter/client"
	"github.com/ffddorf/unms-exporter/client/devices"
	openapi "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var _ prom.Collector = (*Exporter)(nil)

const namespace = "unms"

type metricSpec struct {
	help   string
	labels []string
}

func newSpec(help string, labels []string) metricSpec {
	return metricSpec{help, labels}
}

var defaultLabels = []string{
	"deviceId", "deviceName", "deviceMac", "role", "siteId", "siteName",
}

func (s metricSpec) intoDesc(name string) *prom.Desc {
	labels := make([]string, 0, len(s.labels)+2)
	labels = append(labels, defaultLabels...)
	labels = append(labels, s.labels...)
	return prom.NewDesc(namespace+"_"+name, s.help, labels, prom.Labels{})
}

var metricSpecs = map[string]metricSpec{
	"device_cpu": newSpec("CPU usage in percent", []string{}),
	"device_ram": newSpec("RAM usage in percent", []string{}),
}

type Exporter struct {
	api     *client.UNMSAPI
	metrics map[string]*prom.Desc

	// Internal metrics about the exporter
	im  internalMetrics
	log logrus.FieldLogger
}

func New(log logrus.FieldLogger, host string, token string) *Exporter {
	conf := client.DefaultTransportConfig()
	conf.Schemes = []string{"https"}
	conf.Host = host
	api := client.NewHTTPClientWithConfig(strfmt.Default, conf)

	client, ok := api.Transport.(*openapi.Runtime)
	if !ok {
		panic(fmt.Errorf("Invalid openapi transport: %T", api.Transport))
	}
	auth := openapi.APIKeyAuth("x-auth-token", "header", token)
	client.DefaultAuthentication = auth

	metrics := make(map[string]*prom.Desc)
	for name, spec := range metricSpecs {
		metrics[name] = spec.intoDesc(name)
	}

	im := newInternalMetrics()

	return &Exporter{api, metrics, im, log}
}

func (e *Exporter) Describe(out chan<- *prom.Desc) {
	for _, desc := range e.metrics {
		out <- desc
	}
	e.im.Describe(out)
}

func (e *Exporter) Collect(out chan<- prom.Metric) {
	defer e.im.Collect(out)

	err := e.collectImpl(out)
	if err != nil {
		e.log.WithError(err).Warn("Metric collection failed")
		e.im.errors.Inc()
	} else {
		e.im.success.Inc()
	}
}

var (
	defaultWithInterfaces = true
	defaultDevicesParams  = &devices.GetDevicesParams{
		WithInterfaces: &defaultWithInterfaces,
	}
)

func (e *Exporter) collectImpl(out chan<- prom.Metric) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := &devices.GetDevicesParams{
		WithInterfaces: new(bool),
		Context:        ctx,
	}
	devices, err := e.api.Devices.GetDevices(params, nil)
	if err != nil {
		return err
	}

	for _, device := range devices.Payload {
		deviceLabels := []string{
			*device.Identification.ID,       // deviceId
			device.Identification.Name,      // deviceName
			device.Identification.Mac,       // mac
			device.Identification.Role,      // role
			*device.Identification.Site.ID,  // siteId
			device.Identification.Site.Name, // siteName
		}
		out <- prom.MustNewConstMetric(e.metrics["device_cpu"], prom.GaugeValue, device.Overview.CPU, deviceLabels...)
		out <- prom.MustNewConstMetric(e.metrics["device_ram"], prom.GaugeValue, device.Overview.RAM, deviceLabels...)
	}

	return nil
}
