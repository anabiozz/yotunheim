package handlers

import (
	"net/http"

	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/gorilla/mux"
)

// ExposeHandlers ...
func ExposeHandlers(router *mux.Router, env common.Env) *mux.Router {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RedirectHandler(w, r)
	}).Methods("GET")

	// Start point for reactjs

	router.HandleFunc("/{category}", func(w http.ResponseWriter, r *http.Request) {
		DashboardHandler(w, r)
	}).Methods("GET")

	subrouter := router.PathPrefix("/{category}/").Subrouter()
	subrouter.HandleFunc("/{category}", func(w http.ResponseWriter, r *http.Request) {
		DashboardHandler(w, r)
	}).Methods("GET")

	subsubrouter := subrouter.PathPrefix("/{category}/").Subrouter()
	subsubrouter.HandleFunc("/{category}", func(w http.ResponseWriter, r *http.Request) {
		DashboardHandler(w, r)
	}).Methods("GET")
	return router
}
