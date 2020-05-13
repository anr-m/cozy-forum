package controllers

import (
	"net/http"
)

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}
	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}
