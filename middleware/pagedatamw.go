package middleware

import (
	"net/http"

	"cozy-forum/controllers"
	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/sessions"
)

// PageDataMW is middleware to check if the user is logged in
func PageDataMW(loginrequired bool, handler func(http.ResponseWriter, *http.Request, models.PageData)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := sessions.GetUser(w, r)
		if controllers.InternalError(w, r, err) {
			return
		}
		categories, err := db.GetAllCategories()
		if controllers.InternalError(w, r, err) {
			return
		}
		data := models.PageData{"", categories, user, nil}
		if loginrequired && user.UserID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handler(w, r, data)
	}
}
