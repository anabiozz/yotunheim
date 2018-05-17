package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: kernel

fieldKey         fieldType
--------         ---------
boot_time        integer
context_switches integer
entropy_avail    integer
interrupts       integer
processes_forked integer

*/

// KernelStats ...
type KernelStats struct{}

// Gather ...
func (KernelStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "kernel"

	influxMetrics := datastore.InfluxMetrics{}
	tableMetrics := datastore.TableMetrics{}
	influxMetrics.Metric = make([]datastore.TableMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(context_switches) as context_switches from kernel WHERE time >= now() - 30m GROUP BY time(1m)")

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
	metrics.Add("kernel", func() backend.Gatherer {
		return KernelStats{}
	})
}
