package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type NetStat struct{}

func (_ NetStat) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from net WHERE time >= now() - 5s")
	acc.AddTable("net", res, err)
}

func init() {
	metrics.Add("net", func() backend.Gatherer {
		return NetStat{}
	})
}
