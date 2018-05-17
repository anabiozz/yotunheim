package metrics

import (
	"github.com/anabiozz/yotunheim/backend/common/datastore"
)

type accumulator struct {
	metrics    chan datastore.InfluxMetrics
	mapmetrics chan datastore.Response
	getter     MetricGetter
}

// MetricGetter ...
type MetricGetter interface {
	GetMetric(influxMetrics datastore.InfluxMetrics) datastore.InfluxMetrics
	GetMetrics(influxMetrics datastore.Response) datastore.Response
}

// NewAccumulator ...
func NewAccumulator(
	runningInput MetricGetter,
	metrics chan datastore.InfluxMetrics,
	mapmetrics chan datastore.Response) *accumulator {
	acc := accumulator{
		getter:     runningInput,
		metrics:    metrics,
		mapmetrics: mapmetrics,
	}
	return &acc
}

func (ac *accumulator) AddMetric(influxMetrics datastore.InfluxMetrics) {
	ac.metrics <- ac.getter.GetMetric(influxMetrics)
}

func (ac *accumulator) AddMetrics(influxMetrics datastore.Response) {
	ac.mapmetrics <- ac.getter.GetMetrics(influxMetrics)
}

func (ac *accumulator) AddTable(influxMetrics datastore.InfluxMetrics) {
	ac.metrics <- ac.getter.GetMetric(influxMetrics)
}
