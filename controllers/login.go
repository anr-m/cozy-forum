package controllers

import (
	"net/http"
	"regexp"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/sessions"
	"cozy-forum/tpl"

	"golang.org/x/crypto/bcrypt"
)

// Login route
func Login(w http.ResponseWriter, r *http.Request, data models.PageData) {

	data.PageTitle = "Login"

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
			return
		} else if password == "" {
			data.Data = "Password must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
			return
		}

		if regex.MatchString(username) {
			emailExists, err := db.EmailExists(username)
			if InternalError(w, r, err) {
				return
			}
			if !emailExists {
				data.Data = "Invalid email"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
				return
			}
			user, err := db.GetUserByEmail(username)
			if InternalError(w, r, err) {
				return
			}
			err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
				return
			}
			err = sessions.CreateSession(user.UserID, w)
			if InternalError(w, r, err) {
				return
			}
		} else {
			usernameExists, err := db.UsernameExists(username)
			if InternalError(w, r, err) {
				return
			}
			if !usernameExists {
				data.Data = "Invalid username"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
				return
			}
			user, err := db.GetUserByUsername(username)
			if InternalError(w, r, err) {
				return
			}
			err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
				return
			}
			err = sessions.CreateSession(user.UserID, w)
			if InternalError(w, r, err) {
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, tpl.ExecuteTemplate(w, "login.html", data))
	}
}
