package controllers

import "net/http"

// NotFoundHandler for handling 404 route
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotFound, "404 Not Found")
}
