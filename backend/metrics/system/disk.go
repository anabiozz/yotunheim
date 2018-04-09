package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type DiskStats struct{}

func (_ DiskStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from disk WHERE time >= now() - 5s")
	acc.AddTable("disk", res, err)
}

func init() {
	metrics.Add("disk", func() backend.Gatherer {
		return DiskStats{}
	})
}
