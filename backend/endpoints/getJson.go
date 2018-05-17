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

// GetJSONnEndpoint ...
func GetJSONnEndpoint(
	w http.ResponseWriter,
	r *http.Request,
	e *common.Env,
	newConfig *config.Config) {

	metricChannel := make(chan datastore.InfluxMetrics, 100)

	for _, input := range newConfig.Inputs {
		acc := metrics.NewAccumulator(input, metricChannel, nil)
		input.Metrics.Gather(e.DB, acc)
	}

	payload, err := json.Marshal(<-metricChannel)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
