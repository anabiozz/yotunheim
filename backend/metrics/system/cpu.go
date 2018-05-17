package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: cpu

fieldKey         fieldType
--------         ---------
usage_guest      float
usage_guest_nice float
usage_idle       float
usage_iowait     float
usage_irq        float
usage_nice       float
usage_softirq    float
usage_steal      float
usage_system     float
usage_user       float

tagKey
------
cpu

*/

// CPUStats ...
type CPUStats struct{}

// Gather ...
func (CPUStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "cpu"

	tableMetrics := datastore.InfoMetrics{}

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make([]datastore.InfoMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT 100-MEAN(usage_idle) AS usage_idle FROM cpu WHERE time >= now() - 30m GROUP BY time(1m)")

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
	metrics.Add("cpu", func() backend.Gatherer {
		return CPUStats{}
	})
}
