package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ffddorf/unms-exporter/client"
	"github.com/ffddorf/unms-exporter/models"
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
	labels := make([]string, 0, len(s.labels)+len(defaultLabels))
	labels = append(labels, defaultLabels...)
	labels = append(labels, s.labels...)
	return prom.NewDesc(namespace+"_"+name, s.help, labels, prom.Labels{})
}

var interfaceLabels = []string{"ifName", "ifDescr", "ifPos", "ifType"}

var metricSpecs = map[string]metricSpec{
	"device_cpu": newSpec("CPU usage in percent", nil),
	"device_ram": newSpec("RAM usage in percent", nil),

	"device_enabled":     newSpec("Whether device is enabled", nil),
	"device_maintenance": newSpec("Whether device is in maintenance", nil),

	"device_uptime":      newSpec("Duration the device is up in seconds", nil),
	"device_last_seen":   newSpec("Unix epoch when device was last seen", nil),
	"device_last_backup": newSpec("Unix epoch when last backup was made", nil),

	"interface_enabled": newSpec("Whether interface is enabled", interfaceLabels),
	"interface_plugged": newSpec("Whether interface has a plugged link", interfaceLabels),
	"interface_up":      newSpec("Whether interface is up", interfaceLabels),

	"interface_dropped":   newSpec("Number of packets dropped on an interface", interfaceLabels),
	"interface_errors":    newSpec("Number of interface errors", interfaceLabels),
	"interface_rx_bytes":  newSpec("Bytes received on an interface", interfaceLabels),
	"interface_tx_bytes":  newSpec("Bytes sent on an interface", interfaceLabels),
	"interface_rx_rate":   newSpec("Receive rate on an interface", interfaceLabels),
	"interface_tx_rate":   newSpec("Transmit rate on an interface", interfaceLabels),
	"interface_poe_power": newSpec("POE power output on an interface", interfaceLabels),

	"wan_rx_bytes": newSpec("Bytes received on WAN interface", nil),
	"wan_tx_bytes": newSpec("Bytes sent on WAN interface", nil),
	"wan_rx_rate":  newSpec("Receive rate on WAN interface", nil),
	"wan_tx_rate":  newSpec("Transmit rate on WAN interface", nil),
}

type Exporter struct {
	api     *client.UNMSAPI
	metrics map[string]*prom.Desc
	extras  ExtraMetrics

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
	return &Exporter{
		api:     api,
		metrics: metrics,
		im:      im,
		log:     log,
	}
}

func (e *Exporter) Describe(out chan<- *prom.Desc) {
	e.DescribeContext(context.Background(), out)
}

func (e *Exporter) DescribeContext(ctx context.Context, out chan<- *prom.Desc) {
	for _, desc := range e.metrics {
		out <- desc
	}
	e.im.Describe(out)
}

func (e *Exporter) Collect(out chan<- prom.Metric) {
	e.CollectContext(context.Background(), out)
}

func (e *Exporter) CollectContext(ctx context.Context, out chan<- prom.Metric) {
	defer e.im.Collect(out)

	if err := e.collectImpl(ctx, out); err != nil {
		e.log.WithError(err).Warn("Metric collection failed")
		e.im.errors.Inc()
	} else {
		e.im.success.Inc()
	}
}

func derefOrFalse(in *bool) bool {
	if in == nil {
		return false
	}
	return *in
}

func boolToGauge(in bool) float64 {
	if in {
		return 1
	}
	return 0
}

func timeToGauge(ts strfmt.DateTime) float64 {
	return float64(time.Time(ts).Unix())
}

func (e *Exporter) newMetric(name string, typ prom.ValueType, val float64, labels ...string) prom.Metric {
	return prom.MustNewConstMetric(e.metrics[name], typ, val, labels...)
}

func (e *Exporter) collectImpl(ctx context.Context, out chan<- prom.Metric) error {
	devices, err := e.fetchDeviceData(ctx)
	if err != nil {
		return err
	}

	for _, device := range devices {
		siteID := "no-site-id"
		siteName := "no-site"
		if s := device.Identification.Site; s != nil {
			if s.ID != nil {
				siteID = *s.ID
			}
			if s.Name != "" {
				siteName = s.Name
			}
		}

		deviceLabels := []string{
			*device.Identification.ID,  // deviceId
			device.Identification.Name, // deviceName
			device.Identification.Mac,  // mac
			device.Identification.Role, // role
			siteID,
			siteName,
		}

		out <- e.newMetric("device_enabled", prom.GaugeValue, boolToGauge(derefOrFalse(device.Enabled)), deviceLabels...)
		if device.Meta != nil {
			out <- e.newMetric("device_maintenance", prom.GaugeValue, boolToGauge(derefOrFalse(device.Meta.Maintenance)), deviceLabels...)
		}
		if device.Overview != nil {
			out <- e.newMetric("device_cpu", prom.GaugeValue, device.Overview.CPU, deviceLabels...)
			out <- e.newMetric("device_ram", prom.GaugeValue, device.Overview.RAM, deviceLabels...)
			out <- e.newMetric("device_uptime", prom.GaugeValue, device.Overview.Uptime, deviceLabels...)
			out <- e.newMetric("device_last_seen", prom.CounterValue, timeToGauge(device.Overview.LastSeen), deviceLabels...)
		}
		if device.LatestBackup != nil && device.LatestBackup.Timestamp != nil {
			out <- e.newMetric("device_last_backup", prom.GaugeValue, timeToGauge(*device.LatestBackup.Timestamp), deviceLabels...)
		}

		seenInterfaces := make(map[string]struct{})

		var wanIF *models.DeviceInterfaceSchema
		for _, intf := range device.Interfaces {
			if intf.Identification == nil {
				continue
			}

			// sometimes UNMS duplicates an interface in the list.
			// skip it so we don't send duplicate metrics.
			if _, ok := seenInterfaces[intf.Identification.Name]; ok {
				continue
			}
			seenInterfaces[intf.Identification.Name] = struct{}{}

			if intf.Identification.Name == device.Identification.WanInterfaceID {
				wanIF = intf
			}

			intfLabels := make([]string, 0, len(deviceLabels)+len(interfaceLabels))
			intfLabels = append(intfLabels, deviceLabels...)
			intfLabels = append(intfLabels,
				intf.Identification.Name,                            // ifName
				derefOrEmpty(intf.Identification.Description),       // ifDescr
				strconv.FormatInt(intf.Identification.Position, 10), // ifPos
				intf.Identification.Type,                            // ifType
			)

			out <- e.newMetric("interface_enabled", prom.GaugeValue, boolToGauge(intf.Enabled), intfLabels...)
			if intf.Status != nil {
				out <- e.newMetric("interface_plugged", prom.GaugeValue, boolToGauge(intf.Status.Plugged), intfLabels...)
				out <- e.newMetric("interface_up", prom.GaugeValue, boolToGauge(intf.Status.Status == "active"), intfLabels...)
			}

			if intf.Statistics != nil {
				out <- e.newMetric("interface_dropped", prom.CounterValue, intf.Statistics.Dropped, intfLabels...)
				out <- e.newMetric("interface_errors", prom.CounterValue, intf.Statistics.Errors, intfLabels...)
				out <- e.newMetric("interface_rx_bytes", prom.CounterValue, intf.Statistics.Rxbytes, intfLabels...)
				out <- e.newMetric("interface_tx_bytes", prom.CounterValue, intf.Statistics.Txbytes, intfLabels...)
				out <- e.newMetric("interface_rx_rate", prom.GaugeValue, intf.Statistics.Rxrate, intfLabels...)
				out <- e.newMetric("interface_tx_rate", prom.GaugeValue, intf.Statistics.Txrate, intfLabels...)
				out <- e.newMetric("interface_poe_power", prom.GaugeValue, intf.Statistics.PoePower, intfLabels...)
			}
		}

		// WAN metrics
		if wanIF != nil && wanIF.Statistics != nil {
			out <- e.newMetric("wan_rx_bytes", prom.CounterValue, wanIF.Statistics.Rxbytes, deviceLabels...)
			out <- e.newMetric("wan_tx_bytes", prom.CounterValue, wanIF.Statistics.Txbytes, deviceLabels...)
			out <- e.newMetric("wan_rx_rate", prom.GaugeValue, wanIF.Statistics.Rxrate, deviceLabels...)
			out <- e.newMetric("wan_tx_rate", prom.GaugeValue, wanIF.Statistics.Txrate, deviceLabels...)
		}

		// Ping metrics, if enabled
		if e.extras.Ping {
			ratio := 1.0
			if ping := device.PingMetrics(); ping != nil {
				if ping.PacketsSent > 0 {
					ratio = float64(ping.PacketsLost) / float64(ping.PacketsSent)
				}
				out <- e.newMetric("ping_rtt_best_seconds", prom.GaugeValue, ping.Best.Seconds(), deviceLabels...)
				out <- e.newMetric("ping_rtt_mean_seconds", prom.GaugeValue, ping.Mean.Seconds(), deviceLabels...)
				out <- e.newMetric("ping_rtt_worst_seconds", prom.GaugeValue, ping.Worst.Seconds(), deviceLabels...)
				out <- e.newMetric("ping_rtt_std_deviation_seconds", prom.GaugeValue, ping.StdDev.Seconds(), deviceLabels...)
			}
			out <- e.newMetric("ping_loss_ratio", prom.GaugeValue, ratio, deviceLabels...)
		}
	}

	return nil
}

func derefOrEmpty(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}
