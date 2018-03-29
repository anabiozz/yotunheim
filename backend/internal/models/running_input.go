package models

import (
	"fmt"
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"log"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

// RunningInput ...
type RunningInput struct {
	Metrics backend.Gatherer
}

// GetMetric ...
func (r *RunningInput) GetMetric(name string, chartType string, metrics []influx.Result, err error) datastore.InfluxMetrics {
	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]interface{}, 0)
	if len(metrics) > 0 && len(metrics[0].Series) > 0 {
		for _, ser := range metrics[0].Series[0].Values {
			log.Println(ser)
			influxMetricItem := datastore.InfluxMetricItem{}
			t, _ := time.Parse(time.RFC3339, ser[0].(string))
			influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
			influxMetricItem.Payload = ser[1]
			influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
		}
		influxMetrics.ChartType = chartType
	}
	return influxMetrics
}

// NewRunningInput ...
func NewRunningInput(metrics backend.Gatherer) *RunningInput {
	return &RunningInput{
		Metrics: metrics,
	}
}
