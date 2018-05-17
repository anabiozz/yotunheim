package system

import (
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
	tableMetrics := datastore.InfoMetrics{}
	influxMetrics.Metric = make([]datastore.InfoMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(io_time) as io_time from diskio WHERE time >= now() - 30m GROUP BY time(1m)")

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
	metrics.Add("diskio", func() backend.Gatherer {
		return DiskIOStats{}
	})
}
