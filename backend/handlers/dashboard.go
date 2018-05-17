package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// DashboardHandler ..
func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	fp := filepath.Join("../go/src/github.com/anabiozz/yotunheim/backend/public", "index.html")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "server", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	// data, err := ioutil.ReadFile(string("../go/src/github.com/anabiozz/yotunheim/backend/public/index.html"))
	// if err == nil {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	w.Write(data)
	// } else {
	// 	w.WriteHeader(404)
	// 	w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	// }
}
