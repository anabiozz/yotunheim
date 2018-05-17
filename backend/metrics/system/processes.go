package system

import (
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
	tableMetrics := datastore.InfoMetrics{}
	influxMetrics.Metric = make([]datastore.InfoMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(*) from processes WHERE time >= now() - 30m GROUP BY time(1m)")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {

		if metrics[0].Series[0].Values[1] != nil {
			tableMetrics.Titles = make([]string, len(metrics[0].Series[0].Columns))
			tableMetrics.Value = make([][]interface{}, len(metrics[0].Series[0].Values))

			copy(tableMetrics.Titles, metrics[0].Series[0].Columns)
			copy(tableMetrics.Value, metrics[0].Series[0].Values)
			influxMetrics.Metric = append(influxMetrics.Metric, tableMetrics)

			influxMetrics.InfoType = backend.Counter
			influxMetrics.InfoName = name
		}
	}
	acc.AddMetric(influxMetrics)
}

func init() {
	metrics.Add("processes", func() backend.Gatherer {
		return ProcessesStats{}
	})
}
