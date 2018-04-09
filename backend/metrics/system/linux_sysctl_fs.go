package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"

	influx "github.com/influxdata/influxdb/client/v2"
)

type LinuxSysctlFsStats struct{}

func (_ LinuxSysctlFsStats) Gather(c datastore.Datastore, acc backend.Accumulator) {
	res, err := datastore.QueryDB(c.(influx.Client), "SELECT * from linux_sysctl_fs WHERE time >= now() - 5s")
	acc.AddTable("linux_sysctl_fs", res, err)
}

func init() {
	metrics.Add("linux_sysctl_fs", func() backend.Gatherer {
		return LinuxSysctlFsStats{}
	})
}
