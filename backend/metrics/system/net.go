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

name: net

fieldKey              fieldType
--------              ---------
bytes_recv            integer
bytes_sent            integer
drop_in               integer
drop_out              integer
err_in                integer
err_out               integer
icmp_inaddrmaskreps   integer
icmp_inaddrmasks      integer
icmp_incsumerrors     integer
icmp_indestunreachs   integer
icmp_inechoreps       integer
icmp_inechos          integer
icmp_inerrors         integer
icmp_inmsgs           integer
icmp_inparmprobs      integer
icmp_inredirects      integer
icmp_insrcquenchs     integer
icmp_intimeexcds      integer
icmp_intimestampreps  integer
icmp_intimestamps     integer
icmp_outaddrmaskreps  integer
icmp_outaddrmasks     integer
icmp_outdestunreachs  integer
icmp_outechoreps      integer
icmp_outechos         integer
icmp_outerrors        integer
icmp_outmsgs          integer
icmp_outparmprobs     integer
icmp_outredirects     integer
icmp_outsrcquenchs    integer
icmp_outtimeexcds     integer
icmp_outtimestampreps integer
icmp_outtimestamps    integer
ip_defaultttl         integer
ip_forwarding         integer
ip_forwdatagrams      integer
ip_fragcreates        integer
ip_fragfails          integer
ip_fragoks            integer
ip_inaddrerrors       integer
ip_indelivers         integer
ip_indiscards         integer
ip_inhdrerrors        integer
ip_inreceives         integer
ip_inunknownprotos    integer
ip_outdiscards        integer
ip_outnoroutes        integer
ip_outrequests        integer
ip_reasmfails         integer
ip_reasmoks           integer
ip_reasmreqds         integer
ip_reasmtimeout       integer
packets_recv          integer
packets_sent          integer
tcp_activeopens       integer
tcp_attemptfails      integer
tcp_currestab         integer
tcp_estabresets       integer
tcp_incsumerrors      integer
tcp_inerrs            integer
tcp_insegs            integer
tcp_maxconn           integer
tcp_outrsts           integer
tcp_outsegs           integer
tcp_passiveopens      integer
tcp_retranssegs       integer
tcp_rtoalgorithm      integer
tcp_rtomax            integer
tcp_rtomin            integer
udp_ignoredmulti      integer
udp_incsumerrors      integer
udp_indatagrams       integer
udp_inerrors          integer
udp_noports           integer
udp_outdatagrams      integer
udp_rcvbuferrors      integer
udp_sndbuferrors      integer
udplite_ignoredmulti  integer
udplite_incsumerrors  integer
udplite_indatagrams   integer
udplite_inerrors      integer
udplite_noports       integer
udplite_outdatagrams  integer
udplite_rcvbuferrors  integer
udplite_sndbuferrors  integer

tagKey
------
interface

*/

// NetStat ...
type NetStat struct{}

// Gather ...
func (NetStat) Gather(c datastore.Datastore, acc backend.Accumulator) {

	name := "net"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT mean(tcp_maxconn) as tcp_maxconn from net WHERE time >= now() - 5m GROUP BY time(1m) LIMIT 5")

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
	metrics.Add("net", func() backend.Gatherer {
		return NetStat{}
	})
}
