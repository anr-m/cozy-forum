package controllers

import (
	"net/http"

	"cozy-forum/sessions"
)

// Logout route
func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
