package controllers

import (
	"log"
	"net/http"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/sessions"
	"cozy-forum/tpl"
)

// ErrorHandler ...
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	user, _ := sessions.GetUser(w, r)
	categories, _ := db.GetAllCategories()
	tpl.ExecuteTemplate(w, "errorhandler.html", models.PageData{message, categories, user, message})
}

// InternalError ...
func InternalError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		log.Println("500 Internal Server Error: ", err)
		ErrorHandler(w, r, http.StatusInternalServerError, "500 Internal Server Error")
		return true
	}
	return false
}
