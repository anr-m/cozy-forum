package controllers

import (
	"net/http"

	"../sessions"
)

// AuthorizationMW is middleware to check if the user is logged in
func AuthorizationMW(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(w, r) {
			Login(w, r)
			return
		}
		handler(w, r)
	}
}
