package handlers

import (
	"fmt"
	"net/http"
)

// RedirectHandler ...
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("http://%s%s/server/common", r.Host, r.URL), http.StatusPermanentRedirect)
}
