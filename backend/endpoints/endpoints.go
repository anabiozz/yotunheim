package endpoints

import (
	"fmt"
	"net/http"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/gorilla/mux"
)

// ExposeEndpoints ...
func ExposeEndpoints(router *mux.Router, env common.Env) *mux.Router {
	router.HandleFunc("/api/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/settings")
	}).Methods("PUT")

	router.HandleFunc("/api/get-common-charts", func(w http.ResponseWriter, r *http.Request) {
		GetCommonCharts(w, r, &env)
		// send config with charts initial state in request body?
	}).Methods("GET")

	router.HandleFunc("/api/get-network-charts", func(w http.ResponseWriter, r *http.Request) {
		GetNetworkCharts(w, r, &env)
	}).Methods("GET")

	router.HandleFunc("/api/get-apps-info", func(w http.ResponseWriter, r *http.Request) {
		GetAppsInfo(w, r, &env)
	}).Methods("GET")

	router.HandleFunc("/api/get-apps-count", func(w http.ResponseWriter, r *http.Request) {
		GetAppsCount(w, r, &env)
	}).Methods("GET")

	return router
}
