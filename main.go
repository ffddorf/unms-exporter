package main

import (
	"net/http"
	"os"

	"github.com/ffddorf/unms-exporter/internal/cli/config"
	"github.com/ffddorf/unms-exporter/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	conf, err := config.New(os.Args[1:])
	if err != nil {
		log.WithError(err).Fatal("configuration failure")
	}

	log.SetLevel(conf.LogLevel)

	h := handler.New(log, conf.TokenPerHost)
	h = handler.Logging(log, h)

	log.WithField("addr", conf.ServerAddr).Info("Server starting...")
	if err := http.ListenAndServe(conf.ServerAddr, h); err != nil {
		log.WithError(err).Warn("HTTP server failed")
	}
}
