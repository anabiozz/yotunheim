package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anabiozz/yotunheim/backend/common"
)

// GetAppsInfo ...
func GetAppsCount(w http.ResponseWriter, r *http.Request, e *common.Env) {

	// Create new config
	metricArray := []string{"docker", "influxdb"}

	payload, err := json.Marshal(metricArray)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
