package endpoints

import (
	"heimdall_project/yotunheim/backend/common"
	"heimdall_project/yotunheim/backend/common/datastore"
	"log"
	"strings"

	"github.com/kataras/iris"
)

// GetJSONnEndpoint ...
func GetJSONnEndpoint(e *common.Env) iris.Handler {
	return func(ctx iris.Context) {
		influxMetrics := datastore.InfluxMetrics{}
		influxMetrics.Response = make(map[string][]interface{})

		influxCPUUsage, err := datastore.CPUUsageInfluxQuery(e.DB)
		influxMemUsage, err := datastore.MemUsageInfluxQuery(e.DB)
		if strings.HasSuffix(err.Error(), "getsockopt: connection refused") {
			log.Println("connection to localhost:8086 does not exist")
		}

		if len(influxCPUUsage) > 0 && len(influxCPUUsage[0].Series) > 0 {
			for _, ser := range influxCPUUsage[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				influxMetricItem.Timestamp = ser[0]
				influxMetricItem.Payload = ser[1]
				influxMetrics.Batch = append(influxMetrics.Batch, influxMetricItem)
			}
			influxMetrics.Response["cpu"] = influxMetrics.Batch
		}

		influxMetrics.Batch = append(influxMetrics.Batch[:0])

		if len(influxMemUsage) > 0 && len(influxMemUsage[0].Series) > 0 {
			for _, ser := range influxMemUsage[0].Series[0].Values {
				influxMetricItem := datastore.InfluxMetricItem{}
				influxMetricItem.Timestamp = ser[0]
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
