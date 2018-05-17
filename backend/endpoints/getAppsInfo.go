package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	influx "github.com/influxdata/influxdb/client/v2"
)

// Info ...
type Info struct {
	Title string
	Value interface{}
}

// InfoResponse ...
type InfoResponse struct {
	Name    string
	Metrics []Info
}

// GetAppsInfo ...
func GetAppsInfo(w http.ResponseWriter, r *http.Request, e *common.Env) {

	// Create new config
	metricArray := []string{"docker", "cpu"}
	metricChannel := make(chan []influx.Result, len(metricArray))
	infoResponse := InfoResponse{}
	resultArray := make([]InfoResponse, 0)

	var wg sync.WaitGroup

	wg.Add(len(metricArray))

	for _, app := range metricArray {
		go func(app string) {
			defer wg.Done()
			metric, _ := datastore.QueryDB(e.DB.(influx.Client), "SELECT last(*) FROM "+app+" WHERE time >= now() - 1m")
			metricChannel <- metric
		}(app)
	}

	wg.Wait()

	for metric := range metricChannel {
		if len(metricChannel) == 0 {
			close(metricChannel)
		}
		if len(metric[0].Series) > 0 {

			metrics := make([]Info, 0)
			name := metric[0].Series[0].Name
			columns := metric[0].Series[0].Columns

			switch name {
			case "docker":
				for index, title := range columns {

					jsonNumber, _ := metric[0].Series[0].Values[0][index].(json.Number)
					value := jsonNumber.String()
					info := Info{}

					switch title {

					case "time":
						continue

					case "last_memory_total":
						info.Title = "memory total"
						value, _ := strconv.ParseFloat(value, 32)
						info.Value = fmt.Sprintf("%f Gb", value/1000/1000/1000)

					case "last_n_containers":
						info.Title = "containers number"
						info.Value = value

					case "last_n_containers_paused":
						info.Title = "containers paused"
						info.Value = value

					case "last_n_containers_running":
						info.Title = "containers running"
						info.Value = value

					case "last_n_containers_stopped":
						info.Title = "containers stopped"
						info.Value = value

					case "last_n_cpus":
						continue

					case "last_n_goroutines":
						info.Title = "goroutines number"
						info.Value = value

					case "last_n_images":
						info.Title = "images number"
						info.Value = value

					case "last_n_listener_events":
						info.Title = "listener events number"
						info.Value = value

					case "last_n_used_file_descriptors":
						info.Title = "used file descriptors number"
						info.Value = value
					}

					metrics = append(metrics, info)
				}
			}
			infoResponse.Metrics = metrics
			infoResponse.Name = name
		}
		resultArray = append(resultArray, infoResponse)
	}

	payload, err := json.Marshal(resultArray)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
