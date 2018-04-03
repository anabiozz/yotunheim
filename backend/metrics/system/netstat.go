package system

import (
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type Netstat struct{}

func (_ Netstat) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from netstat WHERE time >= now() - 5s")
	acc.AddTable("netstat", res, err)
}

func init() {
	metrics.Add("netstat", func() backend.Gatherer {
		return Netstat{}
	})
}
