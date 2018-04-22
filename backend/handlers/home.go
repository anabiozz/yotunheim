package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("http://%s%s/dashboard", r.Host, r.URL), http.StatusPermanentRedirect)
	return
}
