package system

import (
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type KernelStats struct{}

func (_ KernelStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * as kernel from kernel WHERE time >= now() - 5s")
	acc.AddTable("kernel", res, err)
}

func init() {
	metrics.Add("kernel", func() backend.Gatherer {
		return KernelStats{}
	})
}
