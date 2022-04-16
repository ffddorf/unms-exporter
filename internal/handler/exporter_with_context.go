package handler

import (
	"context"

	"github.com/ffddorf/unms-exporter/exporter"
	prom "github.com/prometheus/client_golang/prometheus"
)

type withContext struct {
	ctx      context.Context
	exporter *exporter.Exporter
}

var _ prom.Collector = (*withContext)(nil)

func (e *withContext) Describe(out chan<- *prom.Desc) {
	e.exporter.DescribeContext(e.ctx, out)
}

func (e *withContext) Collect(out chan<- prom.Metric) {
	e.exporter.CollectContext(e.ctx, out)
}
