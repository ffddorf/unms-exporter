package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ffddorf/unms-exporter/exporter"
	"github.com/ffddorf/unms-exporter/internal/cli/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type handler struct {
	base    *prometheus.Registry
	targets map[string]*exporter.Exporter
	log     logrus.FieldLogger
}

func New(logger logrus.FieldLogger, cfg *config.Config) (http.Handler, error) {
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewGoCollector(),
	)

	exporters := make(map[string]*exporter.Exporter)
	for host, token := range cfg.TokenPerHost {
		host := strings.ToLower(host)
		exporter := exporter.New(logger, host, token)
		if err := exporter.SetExtras(cfg.ExtraMetrics); err != nil {
			return nil, err
		}
		exporters[host] = exporter
	}

	return &handler{
		base:    reg,
		targets: exporters,
		log:     logger.WithField("component", "exporter"),
	}, nil
}

// ServeHTTP realizes a very rudimentary routing.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorResponse(w, http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/":
		h.getIndex(w, r)
	case "/metrics":
		h.getMetrics(w, r)
	case "/favicon.ico":
		h.getFavicon(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *handler) errorResponse(w http.ResponseWriter, code int) {
	text := fmt.Sprintf("%d %s", code, http.StatusText(code))
	http.Error(w, text, code)
}

func (h *handler) getMetrics(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	log := h.log
	reg := h.base

	if target != "" { // /metrics?target=<name>
		exporter, ok := h.targets[target]
		if !ok {
			h.errorResponse(w, http.StatusNotFound)
			return
		}

		log = log.WithField("target", target)
		reg = prometheus.NewPedanticRegistry()
		withCtx := withContext{
			ctx:      r.Context(),
			exporter: exporter,
		}
		if err := reg.Register(&withCtx); err != nil {
			log.WithError(err).Error("invalid exporter")
			h.errorResponse(w, http.StatusInternalServerError)
		}
	}

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog: log,
	}).ServeHTTP(w, r)
}
