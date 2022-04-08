package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ffddorf/unms-exporter/exporter"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	ServerAddr stringythingy       `mapstructure:"listen" split_words:"true"`
	LogLevel   logrus.Level `mapstructure:"log_level" split_words:"true"`

	TokenPerHost map[string]string `mapstructure:"token" envconfig:"-"`
}

const envPrefix = "UNMS_EXPORTER"

func (c *config) validate() error {
	if len(c.TokenPerHost) < 1 {
		return errors.New("No token configured")
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
		LogLevel:   logrus.InfoLevel,
	}

	if err := envconfig.Process(envPrefix, conf); err != nil {
		log.WithError(err).Fatal("failed to read config from env")
	}

	flags := pflag.NewFlagSet("unms_exporter", pflag.ContinueOnError)
	flags.StringP("listen", "l", conf.ServerAddr, "Address for the exporter to listen on")
	flags.StringP("config", "c", "", "Config file to use")
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.WithError(err).Fatal("failed to parse flags")
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.BindPFlags(flags)

	if path := v.GetString("config"); path != "" {
		v.SetConfigFile(path)
		if err := v.ReadInConfig(); err != nil {
			log.WithField("path", path).WithError(err).Fatal("Failed to read config")
		}
	}

	if err := v.Unmarshal(conf); err != nil {
		log.WithError(err).Fatal("failed to read config from flags")
	}

	log.SetLevel(conf.LogLevel)

	if err := conf.validate(); err != nil {
		fmt.Println(flags.FlagUsages())
		log.WithError(err).Fatal("invalid config")
	}

	targets := make(map[string]*prometheus.Registry)
	for host, token := range conf.TokenPerHost {
		host := strings.ToLower(host)
		registry := prometheus.NewPedanticRegistry()
		registry.MustRegister(
			prometheus.NewBuildInfoCollector(),
			prometheus.NewGoCollector(),
		)
		export := exporter.New(
			log.WithField("component", "exporter").WithField("host", host),
			host, token,
		)
		if err := registry.Register(export); err != nil {
			log.WithField("host", host).WithError(err).Fatal("failed to register exporter")
		}
		targets[host] = registry
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.WithFields(logrus.Fields{
			"url":    r.URL,
			"method": r.Method,
		})
		log.Debug("Starting request")

		target := r.URL.Query().Get("target")
		if target == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		registry, ok := targets[strings.ToLower(target)]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
			ErrorLog: log,
			Registry: registry,
		})
		h.ServeHTTP(w, r)
	})

	log.WithField("addr", conf.ServerAddr).Info("Server starting...")
	err := http.ListenAndServe(conf.ServerAddr, handler)
	if err != nil {
		log.WithError(err).Warn("HTTP server failed")
	}
}
