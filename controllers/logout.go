package controllers

import (
	"net/http"

	"../sessions"
)

// Logout route
func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.Logout(w, r)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
