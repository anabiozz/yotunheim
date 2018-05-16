package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: linux_sysctl_fs

fieldKey           fieldType
--------           ---------
aio-max-nr         integer
aio-nr             integer
dentry-age-limit   integer
dentry-nr          integer
dentry-unused-nr   integer
dentry-want-pages  integer
file-max           integer
file-nr            integer
inode-free-nr      integer
inode-nr           integer
inode-preshrink-nr integer

*/

// LinuxSysctlFsStats ...
type LinuxSysctlFsStats struct{}

// Gather ...
func (LinuxSysctlFsStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "linux_sysctl_fs"

	influxMetrics := datastore.InfluxMetrics{}
	tableMetrics := datastore.TableMetrics{}
	influxMetrics.Metric = make([]datastore.TableMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(file-max) as file-max from linux_sysctl_fs WHERE time >= now() - 30m GROUP BY time(1m)")

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
	metrics.Add("linux_sysctl_fs", func() backend.Gatherer {
		return LinuxSysctlFsStats{}
	})
}
