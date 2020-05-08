package controllers

import (
	"net/http"
	"regexp"

	"../db"
	"../errorhandle"
	"../models"
	"../sessions"
	"../tpl"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register route
func Register(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Register", sessions.GetUser(w, r), nil}

	if r.Method == http.MethodPost {
		var newUser models.User
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		password := r.FormValue("password")

		if newUser.Email == "" {
			data.Data = "Email must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if !regex.MatchString(newUser.Email) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if db.EmailExists(newUser.Email) {
			data.Data = "Email already exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.Username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if db.UsernameExists(newUser.Username) {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if len(password) < 8 {
			data.Data = "Password must be at least 8 characters"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.FirstName == "" {
			data.Data = "First Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		} else if newUser.LastName == "" {
			data.Data = "Last Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "register.html", data)
			return
		}

		salt, err := uuid.NewV4()
		errorhandle.Check(err)
		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		errorhandle.Check(err)

		newUser.Hash = hash
		newUser.Salt = salt.String()

		db.CreateUser(&newUser)
		sessions.CreateSession(newUser.UserID, w)

		http.Redirect(w, r, "/index", http.StatusFound)

	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "register.html", data)
	}
}
