package system

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: diskio

fieldKey         fieldType
--------         ---------
io_time          integer
iops_in_progress integer
read_bytes       integer
read_time        integer
reads            integer
weighted_io_time integer
write_bytes      integer
write_time       integer
writes           integer

tagKey
------
name

*/

// DiskIOStats ...
type DiskIOStats struct{}

// Gather ...
func (DiskIOStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "diskio"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(io_time) as io_time from diskio WHERE time >= now() - 5m GROUP BY time(1m) LIMIT 5")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {

		for _, values := range metrics[0].Series[0].Values {

			if values[1] != nil {

				influxMetricItem := datastore.InfluxMetricItem{}
				t, _ := time.Parse(time.RFC3339, values[0].(string))
				influxMetricItem.Xline = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				payload := values[1].(json.Number).Float64
				p, err := payload()
				if err != nil {
					fmt.Println(err)
				}
				influxMetricItem.Payload = int64(p)
				influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
			}

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
