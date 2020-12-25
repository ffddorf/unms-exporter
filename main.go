package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ffddorf/unms-exporter/exporter"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	Host       string `mapstructure:"host"`
	APIToken   string `mapstructure:"api_token" envconfig:"api_token"`
	ServerAddr string `mapstructure:"listen" envconfig:"server_addr"`
}

const EnvPrefix = "UNMS_EXPORTER"

func (c *config) validate() error {
	if c.Host == "" {
		return errors.New("Host cannot be empty")
	}

	if c.APIToken == "" {
		return fmt.Errorf("API token cannot be empty, set %s_API_TOKEN", EnvPrefix)
	}

	if c.ServerAddr == "" {
		return errors.New("Server addr cannot be nil")
	}

	return nil
}

func main() {
	log := logrus.New()

	conf := &config{
		ServerAddr: "[::]:9806",
	}

	if err := envconfig.Process(EnvPrefix, conf); err != nil {
		log.WithError(err).Fatal("failed to read config from env")
	}

	flags := pflag.NewFlagSet("unms_exporter", pflag.ContinueOnError)
	flags.String("host", conf.Host, "UNMS host to connect to")
	flags.String("listen", conf.ServerAddr, "Address for the exporter to listen on")
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.WithError(err).Fatal("failed to parse flags")
	}

	v := viper.New()
	v.BindPFlags(flags)
	if err := v.Unmarshal(conf); err != nil {
		log.WithError(err).Fatal("failed to read config from flags")
	}

	if err := conf.validate(); err != nil {
		fmt.Println(flags.FlagUsages())
		log.WithError(err).Fatal("invalid config")
	}

	export := exporter.New(log.WithField("component", "exporter"), conf.Host, conf.APIToken)

	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(
		prometheus.NewBuildInfoCollector(),
		prometheus.NewGoCollector(),
	)

	if err := registry.Register(export); err != nil {
		log.WithError(err).Fatal("failed to register exporter")
	}

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog: log,
		Registry: registry,
	})

	log.WithField("addr", conf.ServerAddr).Info("Server starting...")
	err := http.ListenAndServe(conf.ServerAddr, handler)
	if err != nil {
		log.WithError(err).Warn("HTTP server failed")
	}
}
