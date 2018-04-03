package system

import (
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type ProcessesStats struct{}

func (_ ProcessesStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from processes WHERE time >= now() - 5s")
	acc.AddTable("processes", res, err)
}

func init() {
	metrics.Add("processes", func() backend.Gatherer {
		return ProcessesStats{}
	})
}
