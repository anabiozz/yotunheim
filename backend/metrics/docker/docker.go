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

// DockerStats ...
type DockerStats struct{}

// Response ...

// Gather ...
func (DockerStats) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "docker"

	allDocker := []string{"docker", "docker_container_blkio", "docker_container_cpu", "docker_container_mem", "docker_container_net"}
	influxMetrics := datastore.InfluxMetrics{}
	response := datastore.Response{}
	response.Metrics = make([]datastore.InfluxMetrics, 0)

	for _, val := range allDocker {

		infoMetrics := datastore.InfoMetrics{}
		influxMetrics.Metric = make([]datastore.InfoMetrics, 0)

		metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(*) FROM "+val+" WHERE time >= now() - 1m")

		if len(metrics) > 0 && len(metrics[0].Series) > 0 {

			infoMetrics.Titles = make([]string, len(metrics[0].Series[0].Columns))
			infoMetrics.Value = make([][]interface{}, len(metrics[0].Series[0].Values))

			copy(infoMetrics.Titles, metrics[0].Series[0].Columns)
			copy(infoMetrics.Value, metrics[0].Series[0].Values)

			influxMetrics.Metric = append(influxMetrics.Metric, infoMetrics)
			influxMetrics.InfoType = backend.Table
			influxMetrics.InfoName = val
		}
		response.Name = name
		response.Metrics = append(response.Metrics, influxMetrics)
	}
	acc.AddMetrics(response)
}

func init() {
	metrics.Add("docker", func() backend.Gatherer {
		return DockerStats{}
	})
}
