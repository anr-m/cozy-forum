package controllers

import "net/http"

// NotFoundHandler for handling 404 route
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
}
