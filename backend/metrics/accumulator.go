package metrics

import (
	"github.com/anabiozz/yotunheim/backend/common/datastore"

	"github.com/anabiozz/yotunheim/backend"

	influx "github.com/influxdata/influxdb/client/v2"
)

type accumulator struct {
	metrics chan datastore.InfluxMetrics
	getter  MetricGetter
}

// MetricMaker ...
type MetricGetter interface {
	GetMetric(name string, chartType string, metrics []influx.Result, err error) datastore.InfluxMetrics
}

// NewAccumulator ...
func NewAccumulator(runningInput MetricGetter, metrics chan datastore.InfluxMetrics) *accumulator {
	acc := accumulator{
		getter:  runningInput,
		metrics: metrics,
	}
	return &acc
}

func (ac *accumulator) AddLine(name string, metrics []influx.Result, err error) {
	ac.metrics <- ac.getter.GetMetric(name, backend.Counter, metrics, err)
}

// AddBar
func (ac *accumulator) AddBar(name string, metrics []influx.Result, err error) {
	ac.metrics <- ac.getter.GetMetric(name, backend.Histogram, metrics, err)
}

func (ac *accumulator) AddTable(name string, metrics []influx.Result, err error) {
	ac.metrics <- ac.getter.GetMetric(name, backend.Table, metrics, err)
}
