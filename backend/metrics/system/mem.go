package system

import (
	"fmt"
	"time"

	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: mem

fieldKey          fieldType
--------          ---------
active            integer
available         integer
available_percent float
buffered          integer
cached            integer
free              integer
inactive          integer
slab              integer
total             integer
used              integer
used_percent      float
wired             integer

*/

// MemStats ...
type MemStats struct{}

// Gather ...
func (MemStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "mem"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]interface{}, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(used_percent) as mem_usage from mem WHERE time >= now() - 5m GROUP BY time(30s) LIMIT 10")

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
	metrics.Add("mem", func() backend.Gatherer {
		return MemStats{}
	})
}
