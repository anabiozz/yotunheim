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

// GetNetworkCharts ...
func GetNetworkCharts(w http.ResponseWriter, r *http.Request, e *common.Env) {

	// Create new config
	newConfig := config.NewConfig()
	metricArray := []string{"net", "netstat"}

	for _, value := range metricArray {
		newConfig.AddInput(value)
	}

	metricChannel := make(chan datastore.InfluxMetrics, 100)
	metricsResult := make([]datastore.InfluxMetrics, 0)

	for _, input := range newConfig.Inputs {
		acc := metrics.NewAccumulator(input, metricChannel, nil)
		input.Metrics.Gather(e.DB, acc)
		metricsResult = append(metricsResult, <-metricChannel)
	}

	payload, err := json.Marshal(metricsResult)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
