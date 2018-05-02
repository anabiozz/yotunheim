package models

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
)

// RunningInput ...
type RunningInput struct {
	Metrics backend.Gatherer
}

// GetMetric ...
func (r *RunningInput) GetMetric(influxMetrics datastore.InfluxMetrics) datastore.InfluxMetrics {
	return influxMetrics
}

// NewRunningInput ...
func NewRunningInput(metrics backend.Gatherer) *RunningInput {
	return &RunningInput{
		Metrics: metrics,
	}
}
