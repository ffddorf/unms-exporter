# UNMS Exporter for Prometheus

Small daemon offering UNMS device statistics as Prometheus metrics.

## Deployment

Premade Docker Images are available at [quay.io](https://quay.io/repository/ffddorf/unms-exporter).

```bash
docker pull quay.io/ffddorf/unms-exporter
```

## Configuration

### Options

Config can be specified via a YAML file, as args or from environment variables.

### Listen Address

- Config: `listen`
- Args: `--listen` or `-l`
- Env: `UNMS_EXPORTER_SERVER_ADDR`

Address the exporter should listen on. Defaults to `[::]:9806`.

### Config File Location

- Args: `--config` or `-c`

Location of the YAML config file to load.

### Log Verbosity

- Config: `log_level`
- Env: `UNMS_EXPORTER_LOG_LEVEL`

Log verbosity level. Defaults to `info`. Use `debug` to get more details.

### UNMS API Tokens

- Config: `token`
- Env: `UNMS_EXPORTER_TOKEN`
  - use a comma-separated list of `instance=token` values

Configures an API token per UNMS instance.

Example:

```yaml
token:
  my-unms-instance.example.org: "my token"
  unms.example.com: "token123"
```

```console
$ UNMS_EXPORTER_TOKEN="my-unms-instance.example.org=my token,unms.example.com=token123" \
    unms-exporter
```

## Prometheus Scrape Setup

The exporter follows the convention for exporters. The UNMS instance to target should be specified using the `target` query parameter.

Here is how to achieve this using a static prometheus config:

```yaml
scrape_configs:
- job_name: unms_exporter
  metrics_path: /
  relabel_configs:
    - source_labels: [__address__]
      target_label: instance
    - source_labels: [__address__]
      target_label: __param_target
    - replacement: 'unms_exporter:9806'
      target_label: __address__

  static_configs:
    - targets:
        - my-unms-instance.example.org
```

## Available Metrics

### Device wide

- `device_cpu`: CPU load average in percent
- `device_ram`: RAM usage in percent
- `device_enabled`: Indicating if device is enabled in UNMS
- `device_maintenance`: Indicating if device is in maintenance mode (useful for muting alerts)
- `device_uptime`: Uptime in seconds
- `device_last_seen`: Last seen as unix timestamp
- `device_last_backup`: Time of last backup as unix timestamp

### Per Interface

- `interface_enabled`: Indicating if interface is enabled
- `interface_plugged`: Indicating if interface has a cable plugged
- `interface_up`: Indicating if interface is considered up
- `interface_dropped`: Number of packets dropped
- `interface_errors`: Number of interface errors
- `interface_rx_bytes`: Bytes received since last reset
- `interface_tx_bytes`: Bytes transmitted since last reset
- `interface_rx_rate`: Bytes received rate (momentarily)
- `interface_tx_rate`: Bytes transmitted rate (momentarily)
- `interface_poe_power`: POE power consumption

### WAN Interface

If an interface is marked as the WAN interface, these metrics are populated.

- `wan_rx_bytes`: Bytes received since last reset
- `wan_tx_bytes`: Bytes transmitted since last reset
- `wan_rx_rate`: Bytes received rate (momentarily)
- `wan_tx_rate`: Bytes transmitted rate (momentarily)
