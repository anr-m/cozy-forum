package controllers

import (
	"net/http"
	"regexp"

	"../db"
	"../models"
	"../sessions"
	"../tpl"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register route
func Register(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Register", models.User{}, nil}

	if r.Method == http.MethodPost {
		var newUser models.User
		var err error
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		password := r.FormValue("password")

		emailExists, err := db.EmailExists(newUser.Email)
		if internalError(w, r, err) {
			return
		}
		usernameExists, err := db.UsernameExists(newUser.Username)
		if internalError(w, r, err) {
			return
		}

		if newUser.Email == "" {
			data.Data = "Email must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if !regex.MatchString(newUser.Email) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if emailExists {
			data.Data = "Email already exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.Username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if usernameExists {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if len(password) < 8 {
			data.Data = "Password must be at least 8 characters"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.FirstName == "" {
			data.Data = "First Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.LastName == "" {
			data.Data = "Last Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		}

		salt, err := uuid.NewV4()
		if internalError(w, r, err) {
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		if internalError(w, r, err) {
			return
		}

		newUser.Hash = hash
		newUser.Salt = salt.String()

		err = db.CreateUser(&newUser)
		if internalError(w, r, err) {
			return
		}

		err = sessions.CreateSession(newUser.UserID, w)
		if internalError(w, r, err) {
			return
		}

		http.Redirect(w, r, "/index", http.StatusFound)

	} else if r.Method == http.MethodGet {
		internalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
	}
}
