package system

import (
	"github.com/anabiozz/yotunheim/backend"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/metrics"
	influx "github.com/influxdata/influxdb/client/v2"
)

/*

name: netstat

fieldKey        fieldType
--------        ---------
tcp_close       integer
tcp_close_wait  integer
tcp_closing     integer
tcp_established integer
tcp_fin_wait1   integer
tcp_fin_wait2   integer
tcp_last_ack    integer
tcp_listen      integer
tcp_none        integer
tcp_syn_recv    integer
tcp_syn_sent    integer
tcp_time_wait   integer
udp_socket      integer

*/

// Netstat ...
type Netstat struct{}

// Gather ...
func (Netstat) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "netstat"

	influxMetrics := datastore.InfluxMetrics{}
	tableMetrics := datastore.InfoMetrics{}
	influxMetrics.Metric = make([]datastore.InfoMetrics, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(tcp_listen) as tcp_listen from netstat WHERE time >= now() - 30m GROUP BY time(1m)")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {

		if metrics[0].Series[0].Values[1] != nil {
			tableMetrics.Titles = make([]string, len(metrics[0].Series[0].Columns))
			tableMetrics.Value = make([][]interface{}, len(metrics[0].Series[0].Values))

			copy(tableMetrics.Titles, metrics[0].Series[0].Columns)
			copy(tableMetrics.Value, metrics[0].Series[0].Values)
			influxMetrics.Metric = append(influxMetrics.Metric, tableMetrics)

			influxMetrics.InfoType = backend.Counter
			influxMetrics.InfoName = name
		}
	}
	acc.AddMetric(influxMetrics)
}

func init() {
	metrics.Add("netstat", func() backend.Gatherer {
		return Netstat{}
	})
}
