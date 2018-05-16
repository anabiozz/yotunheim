package metrics

import (
	"github.com/anabiozz/yotunheim/backend/common/datastore"
)

type accumulator struct {
	metrics chan datastore.InfluxMetrics
	getter  MetricGetter
}

// MetricGetter ...
type MetricGetter interface {
	GetMetric(influxMetrics datastore.InfluxMetrics) datastore.InfluxMetrics
}

// NewAccumulator ...
func NewAccumulator(runningInput MetricGetter, metrics chan datastore.InfluxMetrics) *accumulator {
	acc := accumulator{
		getter:  runningInput,
		metrics: metrics,
	}
	return &acc
}

func (ac *accumulator) AddMetric(influxMetrics datastore.InfluxMetrics) {
	ac.metrics <- ac.getter.GetMetric(influxMetrics)
}

func (ac *accumulator) AddTable(influxMetrics datastore.InfluxMetrics) {
	ac.metrics <- ac.getter.GetMetric(influxMetrics)
}
