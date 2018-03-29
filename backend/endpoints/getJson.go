package endpoints

import (
	"heimdall_project/yotunheim/backend/common"
	"heimdall_project/yotunheim/backend/common/datastore"
	"log"
	"github.com/kataras/iris"
	"heimdall_project/yotunheim/backend/internal/config"

	_"heimdall_project/yotunheim/backend/metrics/all"
	"heimdall_project/yotunheim/backend/metrics"
	"heimdall_project/yotunheim/backend"
	"heimdall_project/asgard"
)

// GetJSONnEndpoint ...
func GetJSONnEndpoint(e *common.Env) iris.Handler {
	return func(ctx iris.Context) {

		// Create new config
		newConfig := config.NewConfig()
		// Filling new config getting data from default config
		err := newConfig.LoadConfig()
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

		inputs := newConfig.InputFilters["inputs"].([]interface{})

		// Filling InputFilters
		for _, value := range inputs {
			newConfig.AddInput(value.(string))
		}

		if len(newConfig.Inputs) == 0 {
			log.Fatalf("ERROR: no inputs found, did you provide a valid config file?")
		}

		influxMetrics := datastore.InfluxMetrics{}
		influxMetrics.Metrics = make(map[string][]interface{})

		metricChannel := make(chan asgard.Metric, 100)

		for _, input := range newConfig.Inputs {
			acc := metrics.NewAccumulator(input, metricChannel)
			input.Metrics.Gather(e.DB, acc)
		}

		//influxCPUUsage, err := datastore.CPUUsageInfluxQuery(e.DB)
		//influxMemUsage, err := datastore.MemUsageInfluxQuery(e.DB)
		//influxDiskUsage, err := datastore.DiskUsageInfluxQuery(e.DB)
		//
		//if err != nil {
		//	if strings.HasSuffix(err.Error(), "getsockopt: connection refused") {
		//		log.Println("connection to localhost:8086 does not exist")
		//	}
		//}

		//if len(influxCPUUsage) > 0 && len(influxCPUUsage[0].Series) > 0 {
		//	for _, ser := range influxCPUUsage[0].Series[0].Values {
		//		influxMetricItem := datastore.InfluxMetricItem{}
		//		t, _ := time.Parse(time.RFC3339, ser[0].(string))
		//		influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
		//		influxMetricItem.Payload = ser[1]
		//		influxMetrics.Metrics["cpu"] = append(influxMetrics.Metrics["cpu"], influxMetricItem)
		//	}
		//	influxMetrics.ChartType = append(influxMetrics.ChartType, "line")
		//}
		//
		//if len(influxMemUsage) > 0 && len(influxMemUsage[0].Series) > 0 {
		//	for _, ser := range influxMemUsage[0].Series[0].Values {
		//		influxMetricItem := datastore.InfluxMetricItem{}
		//		t, _ := time.Parse(time.RFC3339, ser[0].(string))
		//		influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
		//		influxMetricItem.Payload = ser[1]
		//		influxMetricItem.Type = "bar"
		//		influxMetrics.Metrics["mem"] = append(influxMetrics.Metrics["mem"], influxMetricItem)
		//	}
		//}
		//
		//if len(influxDiskUsage) > 0 && len(influxDiskUsage[0].Series) > 0 {
		//	for _, ser := range influxDiskUsage[0].Series[0].Values {
		//		influxMetricItem := datastore.InfluxMetricItem{}
		//		t, _ := time.Parse(time.RFC3339, ser[0].(string))
		//		influxMetricItem.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
		//		influxMetricItem.Payload = ser[1]
		//		influxMetrics.Metrics["disk"] = append(influxMetrics.Metrics["disk"], influxMetricItem)
		//	}
		//	influxMetrics.ChartType = append(influxMetrics.ChartType, "line")
		//}

		if _, err := ctx.JSON(influxMetrics); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	}
}
