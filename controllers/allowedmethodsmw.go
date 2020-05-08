package controllers

import "net/http"

// AllowedMethodsMW is middleware to check for the correct method
func AllowedMethodsMW(methods []string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range methods {
			if r.Method == method {
				handler(w, r)
				return
			}
		}
		errorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
