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

<details><summary>Example: config file (click to open)</summary>

```yaml
# config.yaml
token:
  my-unms-instance.example.org: "my token"
  unms.example.com: "token123"
```

```console
$ unms-exporter --config config.yaml
```

</details>
<details><summary>Example: environment variable (click to open)</summary>

```console
$ UNMS_EXPORTER_TOKEN="my-unms-instance.example.org=my token,unms.example.com=token123" \
    unms-exporter
```

</details>

### Extra metrics

- Config: `extra_metrics` (as Array)
- Args: `--extra-metrics` (as comma-separated list)
- Env: `UNMS_EXPORTER_EXTRA_METRICS` (as comma-separated list)

Enable additional metrics to be exported. These metrics may require extra
HTTP requests, usually one per device, so they are disabled by default.

<details><summary>Example: config file (click to open)</summary>

```yaml
# config.yaml
extras:
- ping
```

```console
$ unms-exporter --config config.yaml
```

</details>
<details><summary>Example: environment variable (click to open)</summary>

```console
$ UNMS_EXPORTER_EXTRA_METRICS="ping" \
    unms-exporter
```

</details>
<details><summary>Example: command line argument (click to open)</summary>

```console
$ unms-exporter --extra-metrics="ping"
```

</details>

#### Available metrics

- `ping`: Fetch statistical data from UNMS and extract and export
  Ping RTT measurements between UNMS and the device.

  <details><summary>Exported metrics (click to open)</summary>

  - `ping_loss_ratio`: Packet loss ratio (range 0-1, with 0.33 meaning 33% packet loss)
  - `ping_rtt_best_seconds`: Best round trip time, in seconds
  - `ping_rtt_mean_seconds`: Mean round trip time, in seconds
  - `ping_rtt_worst_seconds`: Worst round trip time, in seconds
  - `ping_rtt_std_deviation_seconds`: Standard deviation for round trip time, in seconds

  </details>

Further data is available, but not currently exported (see the API
documentation for the `/devices/{id}/statistics` endpoint on your UNMS
installation to get an overview). Feel free to [open a new issue][] to
inquire whether an integration into the exporter is feasable.

[open a new issue]: https://github.com/ffddorf/unms-exporter/issues/new

## Prometheus Scrape Setup

The exporter follows the convention for exporters. The UNMS instance to target should be specified using the `target` query parameter.

Here is how to achieve this using a static prometheus config:

```yaml
scrape_configs:
- job_name: exporters
  static_configs:
    - exporter.example.org:9806 # UNMS exporter
    - exporter.example.org:9100 # node exporter
    - ...

- job_name: unms_exporter
  # for a static target "unms.example.org", rewrite it to
  # "exporter.example.org:9806/metrics?target=unms.example.org",
  # but keep "unms.example.org" as instance label
  relabel_configs:
    - source_labels: [__address__]
      target_label: instance
    - source_labels: [__address__]
      target_label: __param_target
    - replacement: 'exporter.example.org:9806'
      target_label: __address__
  static_configs:
    - targets:
      - my-unms-instance.example.org
```

<details><summary>Upgrade from v0.1.2 or earlier (click to open)</summary>

Previous versions did expose the UNMS metrics under any path on the exporter,
i.e. the following URLs were handled identically:

- `http://localhost:9806/?target=my-unms-instance.example.org`
- `http://localhost:9806/metrics?target=my-unms-instance.example.org`
- `http://localhost:9806/this/is/all/ignored?target=my-unms-instance.example.org`

Additionally, the UNMS exporter has returned a mixed set of internal and
instance-specific metrics.

This has changed and now follows best practices. All UNMS-specific metrics
are now available *only* on the following URL:

- `http://localhost:9806/metrics?target=my-unms-instance.example.org`

Additionally, internal metrics (e.g. Go runtime statistics) can be retrieved
by omitting the `target` parameter:

- `http://localhost:9806/metrics`

</details>

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
