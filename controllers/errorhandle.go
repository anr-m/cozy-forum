package controllers

import (
	"log"
	"net/http"

	"cozy-forum/sessions"
	"cozy-forum/tpl"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	user, _ := sessions.GetUser(w, r)
	tpl.ExecuteTemplate(w, "errorhandler.html", pageData{message, user, message})
}

func internalError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		log.Println("500 Internal Server Error: ", err)
		errorHandler(w, r, http.StatusInternalServerError, "500 Internal Server Error")
		return true
	}
	return false
}
