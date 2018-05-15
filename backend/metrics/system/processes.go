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

name: processes

fieldKey      fieldType
--------      ---------
blocked       integer
dead          integer
idle          integer
paging        integer
running       integer
sleeping      integer
stopped       integer
total         integer
total_threads integer
unknown       integer
zombies       integer

*/

// ProcessesStats ...
type ProcessesStats struct{}

// Gather ...
func (ProcessesStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "processes"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(total) as total from processes WHERE time >= now() - 30m GROUP BY time(1m)")

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
	metrics.Add("processes", func() backend.Gatherer {
		return ProcessesStats{}
	})
}
