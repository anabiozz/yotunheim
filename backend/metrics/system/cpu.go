package system

import (
	influx "github.com/influxdata/influxdb/client/v2"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"
	"heimdall_project/yotunheim/backend"
)

type CPUStats struct {}

func (_ CPUStats) Gather(c datastore.Datastore, acc backend.Accumulator)  {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT (100 - usage_idle) as cpu_usage from cpu WHERE time >= now() - 60s AND cpu = 'cpu-total'")
	acc.AddLine("cpu", res, err)
}

func init()  {
	metrics.Add("cpu", func() backend.Gatherer {
		return  CPUStats{}
	})
}