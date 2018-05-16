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
type DockerStats struct{}

// Gather ...
func (DockerStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	// name := "docker"

	allDocker := []string{"docker", "docker_container_blkio", "docker_container_cpu", "docker_container_mem", "docker_container_net"}
	influxMetrics := datastore.InfluxMetrics{}

	for _, val := range allDocker {
		tableMetrics := datastore.TableMetrics{}
		influxMetrics.Metric = make([]datastore.TableMetrics, 0)

		metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT MEAN(*) FROM "+val+" WHERE time >= now() - 10m GROUP BY time(5m)")

		if len(metrics) > 0 && len(metrics[0].Series) > 0 {

			tableMetrics.Titles = make([]string, len(metrics[0].Series[0].Columns))
			tableMetrics.Value = make([][]interface{}, len(metrics[0].Series[0].Values))
			copy(tableMetrics.Titles, metrics[0].Series[0].Columns)
			copy(tableMetrics.Value, metrics[0].Series[0].Values)

			influxMetrics.Metric = append(influxMetrics.Metric, tableMetrics)
			influxMetrics.ChartType = backend.Table
			influxMetrics.ChartName = val
		}
		acc.AddMetric(influxMetrics)
	}
}

func init() {
	metrics.Add("docker", func() backend.Gatherer {
		return DockerStats{}
	})
}
