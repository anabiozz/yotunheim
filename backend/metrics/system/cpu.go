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

name: cpu

fieldKey         fieldType
--------         ---------
usage_guest      float
usage_guest_nice float
usage_idle       float
usage_iowait     float
usage_irq        float
usage_nice       float
usage_softirq    float
usage_steal      float
usage_system     float
usage_user       float

tagKey
------
cpu

*/

// CPUStats ...
type CPUStats struct{}

// Gather ...
func (CPUStats) Gather(c datastore.Datastore, acc backend.Accumulator, getherTime string, groupby string) {

	name := "cpu"

	influxMetrics := datastore.InfluxMetrics{}
	influxMetrics.Metric = make(map[string][]datastore.InfluxMetricItem, 0)

	metrics, _ := datastore.QueryDB(c.(influx.Client), "SELECT 100-MEAN(usage_idle) AS usage_idle FROM cpu WHERE time >= now() - "+getherTime+" GROUP BY time("+groupby+")")

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
	metrics.Add("cpu", func() backend.Gatherer {
		return CPUStats{}
	})
}
