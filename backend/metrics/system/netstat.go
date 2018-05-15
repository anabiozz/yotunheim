package system

import (
	"encoding/json"
	"fmt"
	"time"

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
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(tcp_listen) as tcp_listen from netstat WHERE time >= now() - 30m GROUP BY time(1m)")

	if len(metrics) > 0 && len(metrics[0].Series) > 0 {

		for _, values := range metrics[0].Series[0].Values {

			if values[1] != nil {

				influxMetricItem := datastore.InfluxMetricItem{}
				t, _ := time.Parse(time.RFC3339, values[0].(string))
				influxMetricItem.Xline = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				payload := values[1].(json.Number).Float64
				p, err := payload()
				if err != nil {
					fmt.Println(err)
				}
				influxMetricItem.Payload = int64(p)
				influxMetrics.Metric[name] = append(influxMetrics.Metric[name], influxMetricItem)
			}
		}
		influxMetrics.ChartType = backend.Counter
	}
	acc.AddMetric(influxMetrics)
}

func init() {
	metrics.Add("netstat", func() backend.Gatherer {
		return Netstat{}
	})
}
