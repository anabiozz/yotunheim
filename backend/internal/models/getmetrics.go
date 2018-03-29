package models

import (
	"heimdall_project/yotunheim/backend/common/datastore"
	influx "github.com/influxdata/influxdb/client/v2"
	"time"
	"fmt"
	"heimdall_project/asgard"
)

func getmetrics(name string, metrics []influx.Result, err error) asgard.Metric  {
	if len(metrics) > 0 && len(metrics[0].Series) > 0 {
		for _, ser := range metrics[0].Series[0].Values {
			influxMetricItem := datastore.InfluxMetricItem{}
			t, _ := time.Parse(time.RFC3339, ser[0].(string))
			influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
			influxMetricItem.Payload = ser[1]
			//influxMetrics.Metrics["cpu"] = append(influxMetrics.Metrics["cpu"], influxMetricItem)
		}
		//influxMetrics.ChartType = append(influxMetrics.ChartType, "line")
	}
}