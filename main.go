package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/ffddorf/unms-exporter/exporter"
	"github.com/ffddorf/unms-exporter/internal/cli/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	conf, err := config.New(os.Args[1:])
	if err != nil {
		log.WithError(err).Fatal("configuration failure")
	}

	log.SetLevel(conf.LogLevel)

	targets := make(map[string]*prometheus.Registry)
	for host, token := range conf.TokenPerHost {
		host := strings.ToLower(host)
		registry := prometheus.NewPedanticRegistry()
		registry.MustRegister(
			prometheus.NewBuildInfoCollector(),
			prometheus.NewGoCollector(),
		)
		logger := log.WithFields(logrus.Fields{
			"component": "exporter",
			"host":      host,
		})
		export := exporter.New(logger, host, token)
		if err := registry.Register(export); err != nil {
			log.WithFields(logrus.Fields{
				logrus.ErrorKey: err,
				"host":          host,
			}).Fatal("failed to register exporter")
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
	if err := http.ListenAndServe(conf.ServerAddr, handler); err != nil {
		log.WithError(err).Warn("HTTP server failed")
	}
}
