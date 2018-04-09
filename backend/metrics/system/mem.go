package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type MemStats struct{}

func (_ MemStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT used_percent as mem_usage from mem WHERE time >= now() - 60s")
	acc.AddLine("mem", res, err)
}

func init() {
	metrics.Add("mem", func() backend.Gatherer {
		return MemStats{}
	})
}
