package controllers

import (
	"net/http"

	"../models"
	"../sessions"
)

// AuthorizationMW is middleware to check if the user is logged in
func AuthorizationMW(loginrequired bool, handler func(http.ResponseWriter, *http.Request, models.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := sessions.GetUser(w, r)
		if internalError(w, r, err) {
			return
		}
		if loginrequired && user.UserID == 0 {
			Login(w, r)
			return
		}
		handler(w, r, user)
	}
}
