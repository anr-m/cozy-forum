package controllers

import (
	"net/http"
	"regexp"

	"../db"
	"../sessions"
	"../tpl"
	"golang.org/x/crypto/bcrypt"
)

// Login route
func Login(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Login", sessions.GetUser(w, r), nil}

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		} else if password == "" {
			data.Data = "Password must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		}

		if regex.MatchString(username) {
			if !db.EmailExists(username) {
				data.Data = "Invalid email"
				w.WriteHeader(http.StatusUnprocessableEntity)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			user := db.GetUserByEmail(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			sessions.CreateSession(user.UserID, w)
		} else {
			if !db.UsernameExists(username) {
				data.Data = "Invalid username"
				w.WriteHeader(http.StatusUnprocessableEntity)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			user := db.GetUserByUsername(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Incorrect password"
				w.WriteHeader(http.StatusUnauthorized)
				tpl.ExecuteTemplate(w, "login.html", data)
				return
			}
			sessions.CreateSession(user.UserID, w)
		}

		http.Redirect(w, r, "/index", http.StatusFound)

	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "login.html", data)
	}
}
