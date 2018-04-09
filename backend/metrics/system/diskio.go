package system

import (
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type DiskIOStats struct{}

func (_ DiskIOStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from diskio WHERE time >= now() - 5s")
	acc.AddTable("diskio", res, err)
}

func init() {
	metrics.Add("diskio", func() backend.Gatherer {
		return DiskIOStats{}
	})
}