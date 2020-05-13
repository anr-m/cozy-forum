package controllers

import (
	"net/http"
	"regexp"
	"unicode"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/sessions"
	"cozy-forum/tpl"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register route
func Register(w http.ResponseWriter, r *http.Request, data models.PageData) {

	data.PageTitle = "Register"

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
		if InternalError(w, r, err) {
			return
		}
		usernameExists, err := db.UsernameExists(newUser.Username)
		if InternalError(w, r, err) {
			return
		}

		if newUser.Email == "" {
			data.Data = "Email must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if !regex.MatchString(newUser.Email) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if emailExists {
			data.Data = "Email already exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.Username == "" {
			data.Data = "Username must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if usernameExists {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if !validPassword(password) {
			data.Data = "Password must have at least 8 characters: 1 upper-case, 1 lower-case, and 1 number."
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.FirstName == "" {
			data.Data = "First Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		} else if newUser.LastName == "" {
			data.Data = "Last Name must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
			return
		}

		salt := uuid.NewV4()

		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		if InternalError(w, r, err) {
			return
		}

		newUser.Hash = hash
		newUser.Salt = salt.String()

		err = db.CreateUser(&newUser)
		if InternalError(w, r, err) {
			return
		}

		err = sessions.CreateSession(newUser.UserID, w)
		if InternalError(w, r, err) {
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, tpl.ExecuteTemplate(w, "register.html", data))
	}
}

func validPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var upper bool
	var lower bool
	var number bool
	for _, c := range password {
		if unicode.IsUpper(c) {
			upper = true
		} else if unicode.IsLower(c) {
			lower = true
		} else if unicode.IsNumber(c) {
			number = true
		}
	}
	return upper && lower && number
}
