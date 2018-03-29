package models

import (
	"heimdall_project/asgard"
	"time"
	"heimdall_project/yotunheim/backend"
	influx "github.com/influxdata/influxdb/client/v2"
)

// RunningInput ...
type RunningInput struct {
	Metrics       backend.Gatherer
	Config      *InputConfig
}

// InputConfig containing a name, interval, and filter
type InputConfig struct {
	Name              string
	NameOverride      string
	MeasurementPrefix string
	MeasurementSuffix string
	Tags              map[string]string
	Interval          time.Duration
}

func (r *RunningInput) GetMetric(name string, metrics []influx.Result, err error) asgard.Metric {
	m := getmetrics(name, metrics, err)
	return m
}

// NewRunningInput ...
func NewRunningInput(metrics backend.Gatherer) *RunningInput {
	return &RunningInput{
		Metrics: metrics,
	}
}
