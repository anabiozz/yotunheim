package models

import (
	"fmt"
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
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

	switch chartType {

	case backend.Counter:
		if len(metrics) > 0 && len(metrics[0].Series) > 0 {
			for _, values := range metrics[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				t, _ := time.Parse(time.RFC3339, values[0].(string))
				influxMetricItem.Xline = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				influxMetricItem.Payload = values[1]
				influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
			}
			influxMetrics.ChartType = chartType
		}

	case backend.Histogram:
		if len(metrics) > 0 && len(metrics[0].Series) > 0 {
			for _, values := range metrics[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				influxMetricItem.Xline = values[2]
				influxMetricItem.Payload = values[1]
				influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
			}
			influxMetrics.ChartType = chartType
		}

	case backend.Table:
		if len(metrics) > 0 && len(metrics[0].Series) > 0 {

			for _, values := range metrics[0].Series[0].Values {
				valueMap := make(map[string]interface{}, 0)

				influxMetricItem := datastore.InfluxMetricItem{}

				for i, value := range values {
					if metrics[0].Series[0].Columns[i] == "time" || value == 0 {
						continue
					}
					fmt.Println(value)
					valueMap[metrics[0].Series[0].Columns[i]] = value
				}

				influxMetricItem.PayloadArray = append(influxMetricItem.PayloadArray, valueMap)
				influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
			}

			influxMetrics.ChartType = chartType
		}
	}
	return influxMetrics
}

// NewRunningInput ...
func NewRunningInput(metrics backend.Gatherer) *RunningInput {
	return &RunningInput{
		Metrics: metrics,
	}
}
