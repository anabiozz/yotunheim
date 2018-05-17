package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: disk

fieldKey     fieldType
--------     ---------
free         integer
inodes_free  integer
inodes_total integer
inodes_used  integer
total        integer
used         integer
used_percent float

tagKey
------
device
fstype
mode
path

*/

// DiskStats ...
type DiskStats struct{}

// Gather ...
func (DiskStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "disk"

	influxMetrics := datastore.InfluxMetrics{}
	tableMetrics := datastore.TableMetrics{}
	influxMetrics.Metric = make([]datastore.TableMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(used_percent) as used_percent from disk WHERE time >= now() - 30m GROUP BY time(1m)")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {

		if metrics[0].Series[0].Values[1] != nil {
			tableMetrics.Titles = make([]string, len(metrics[0].Series[0].Columns))
			tableMetrics.Value = make([][]interface{}, len(metrics[0].Series[0].Values))

			copy(tableMetrics.Titles, metrics[0].Series[0].Columns)
			copy(tableMetrics.Value, metrics[0].Series[0].Values)
			influxMetrics.Metric = append(influxMetrics.Metric, tableMetrics)

			influxMetrics.ChartType = backend.Counter
			influxMetrics.ChartName = name
		}
	}
	acc.AddMetric(influxMetrics)
}

func init() {
	metrics.Add("disk", func() backend.Gatherer {
		return DiskStats{}
	})
}
