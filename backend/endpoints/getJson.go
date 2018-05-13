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

	getherTime := r.FormValue("time")
	groupby := r.FormValue("groupby")

	if getherTime != "" {
		newConfig.Gather.Time = getherTime
	}

	if groupby != "" {
		newConfig.Gather.GroupeBy = groupby
	}

	InfluxResult := datastore.InfluxResult{}

	metricChannel := make(chan datastore.InfluxMetrics, 100)

	for _, input := range newConfig.Inputs {
		acc := metrics.NewAccumulator(input, metricChannel)
		input.Metrics.Gather(e.DB, acc, newConfig.Gather.Time, newConfig.Gather.GroupeBy)
		InfluxResult.Metrics = append(InfluxResult.Metrics, <-metricChannel)
	}

	payload, err := json.Marshal(InfluxResult)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
