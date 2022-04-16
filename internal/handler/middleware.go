package handler

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type logWrapper struct {
	size   int
	status int
	start  time.Time
	http.ResponseWriter
}

func (lw *logWrapper) WriteHeader(code int) {
	lw.status = code
	lw.ResponseWriter.WriteHeader(code)
}

func (lw *logWrapper) Write(b []byte) (int, error) {
	n, err := lw.ResponseWriter.Write(b)
	lw.size += n
	return n, err
}

func (lw *logWrapper) Log(logger logrus.FieldLogger, method, url string) {
	log := logger.WithFields(logrus.Fields{
		"url":      url,
		"method":   method,
		"duration": time.Since(lw.start),
		"size":     lw.size,
		"status":   lw.status,
	})
	switch s := lw.status; {
	case 100 <= s && s <= 299:
		log.Debug("served request")
	case 300 <= s && s <= 399:
		log.Info("served request")
	case 400 <= s && s <= 499:
		log.Warn("served request")
	default: // s < 100 or s >= 500:
		log.Error("served request")
	}
}

func Logging(logger logrus.FieldLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capture := &logWrapper{
			start:          time.Now(),
			ResponseWriter: w,
		}
		defer capture.Log(logger, r.Method, r.URL.String())

		next.ServeHTTP(capture, r)
	})
}
