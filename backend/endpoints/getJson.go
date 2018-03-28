package endpoints

import (
	"fmt"
	"heimdall_project/yotunheim/backend/common"
	"heimdall_project/yotunheim/backend/common/datastore"
	"time"

	"github.com/kataras/iris"
)

// GetJSONnEndpoint ...
func GetJSONnEndpoint(e *common.Env) iris.Handler {
	return func(ctx iris.Context) {
		influxMetrics := datastore.InfluxMetrics{}
		influxMetrics.Response = make(map[string][]interface{})

		influxCPUUsage, _ := datastore.CPUUsageInfluxQuery(e.DB)
		influxMemUsage, _ := datastore.MemUsageInfluxQuery(e.DB)
		// if strings.HasSuffix(err.Error(), "getsockopt: connection refused") {
		// 	log.Println("connection to localhost:8086 does not exist")
		// }

		if len(influxCPUUsage) > 0 && len(influxCPUUsage[0].Series) > 0 {
			for _, ser := range influxCPUUsage[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				t, _ := time.Parse(time.RFC3339, ser[0].(string))
				influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				influxMetricItem.Payload = ser[1]
				influxMetrics.Batch = append(influxMetrics.Batch, influxMetricItem)
			}
			influxMetrics.Response["cpu"] = influxMetrics.Batch
		}

		influxMetrics.Batch = nil

		if len(influxMemUsage) > 0 && len(influxMemUsage[0].Series) > 0 {
			for _, ser := range influxMemUsage[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				t, _ := time.Parse(time.RFC3339, ser[0].(string))
				influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				influxMetricItem.Payload = ser[1]
				influxMetrics.Batch = append(influxMetrics.Batch, influxMetricItem)
			}
			influxMetrics.Response["mem"] = influxMetrics.Batch
		}

		if _, err := ctx.JSON(influxMetrics.Response); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	}
}
