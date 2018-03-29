package system

import (
	"heimdall_project/yotunheim/backend"
	"heimdall_project/yotunheim/backend/common/datastore"
	"heimdall_project/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type DiskStats struct{}

func (_ DiskStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT used_percent, fstype as disk_usage from disk WHERE time >= now() - 5s and device =~ /sda*/")
	acc.AddBar("disk", res, err)
}

func init() {
	metrics.Add("disk", func() backend.Gatherer {
		return DiskStats{}
	})
}
