package system

import (
	"fmt"
	"time"

	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

// DiskIOStats ...
type DiskIOStats struct{}

// Gather ...
func (DiskIOStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "diskio"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]interface{}, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT * from diskio WHERE time >= now() - 20m GROUP BY time(2m) LIMIT 20")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {
		for _, values := range metrics[0].Series[0].Values {
			influxMetricItem := datastore.InfluxMetricItem{}
			t, _ := time.Parse(time.RFC3339, values[0].(string))
			influxMetricItem.Xline = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
			influxMetricItem.Payload = values[1]
			influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
		}
		influxMetrics.ChartType = backend.Counter
	}
	acc.AddMetric(influxMetrics)
}

func init() {
	metrics.Add("diskio", func() backend.Gatherer {
		return DiskIOStats{}
	})
}
