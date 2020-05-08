package controllers

import (
	"net/http"
)

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}
