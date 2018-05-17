package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/internal/config"
	"github.com/anabiozz/yotunheim/backend/metrics"
)

// GetCommonCharts ...
func GetCommonCharts(w http.ResponseWriter, r *http.Request, e *common.Env) {

	// Create new config
	newConfig := config.NewConfig()
	metricArray := []string{"cpu", "mem", "disk", "diskio", "kernel", "precesses"}

	for _, value := range metricArray {
		newConfig.AddInput(value)
	}

	metricChannel := make(chan datastore.InfluxMetrics, 100)
	metricsResult := make([]datastore.InfluxMetrics, 0)

	for _, input := range newConfig.Inputs {
		acc := metrics.NewAccumulator(input, metricChannel, nil)
		input.Metrics.Gather(e.DB, acc)
	}

	for met := range metricChannel {
		if len(metricChannel) == 0 {
			close(metricChannel)
		}
		metricsResult = append(metricsResult, met)
	}

	payload, err := json.Marshal(metricsResult)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
