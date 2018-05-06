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
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(file-max) as file-max from linux_sysctl_fs WHERE time >= now() - 5m GROUP BY time(1m) LIMIT 5")

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
	metrics.Add("linux_sysctl_fs", func() backend.Gatherer {
		return LinuxSysctlFsStats{}
	})
}
