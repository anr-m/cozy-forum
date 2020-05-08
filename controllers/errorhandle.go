package controllers

import (
	"net/http"

	"../sessions"
	"../tpl"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	tpl.ExecuteTemplate(w, "errorhandler.html", pageData{message, sessions.GetUser(w, r), message})
}
