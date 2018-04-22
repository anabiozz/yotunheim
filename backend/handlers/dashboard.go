package handlers

import (
	"io/ioutil"
	"net/http"
)

// DashboardHandler ..
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(string("../go/src/github.com/anabiozz/yotunheim/backend/public/index.html"))
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}
