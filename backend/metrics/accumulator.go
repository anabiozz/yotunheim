package metrics

import (
	influx "github.com/influxdata/influxdb/client/v2"
	"heimdall_project/asgard"
)

type accumulator struct {
	metrics   chan asgard.Metric
	getter     MetricGetter
}


// MetricMaker ...
type MetricGetter interface {
	GetMetric(name string, metrics []influx.Result, err error) asgard.Metric
}


// NewAccumulator ...
func NewAccumulator(
	runningInput MetricGetter,
	metrics chan asgard.Metric) *accumulator {
	acc := accumulator{
		getter:     runningInput,
		metrics:   metrics,
	}
	return &acc
}

func (ac *accumulator) AddLine(name string, metrics []influx.Result, err error)  {
	if m := ac.getter.GetMetric(name, metrics, err); m != nil {
		ac.metrics <- m
	}


}