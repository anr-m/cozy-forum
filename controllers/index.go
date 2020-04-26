package controllers

import (
	"net/http"
	"regexp"

	"../database"
	"../errorhandle"
	"../models"
	"../sessions"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	if sessions.AlreadyLoggedIn(w, r) {
		user := sessions.GetUser(w, r)
		w.Write([]byte("Welcome " + user.FirstName + " " + user.LastName))
	} else {
		w.Write([]byte("You are not logged in"))
	}
}

// Register route
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser models.User
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		r.ParseForm()

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		password := r.FormValue("password")

		if newUser.Email == "" {
			w.Write([]byte("Email must not be empty"))
			return
		} else if !regex.MatchString(newUser.Email) {
			w.Write([]byte("Enter valid email"))
			return
		} else if database.DataBase.EmailExists(newUser.Email) {
			w.Write([]byte("Email exists"))
			return
		} else if newUser.Username == "" {
			w.Write([]byte("Username must not be empty"))
			return
		} else if database.DataBase.UsernameExists(newUser.Username) {
			w.Write([]byte("Username exists"))
			return
		} else if len(password) < 8 {
			w.Write([]byte("Password must be at least 8 characters"))
			return
		} else if newUser.FirstName == "" {
			w.Write([]byte("First Name must not be empty"))
			return
		} else if newUser.LastName == "" {
			w.Write([]byte("Last Name must not be empty"))
			return
		}

		salt, err := uuid.NewV4()
		errorhandle.Check(err)
		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		errorhandle.Check(err)

		newUser.Hash = hash
		newUser.Salt = salt.String()

		database.DataBase.CreateUser(&newUser)
		sessions.CreateSession(newUser.UserID, w)

		w.Write([]byte("Successfully registered"))
	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST">
			<input type="text" name="username" placeholder="username" required><br>
			<input type="email" name="email" placeholder="email" required><br>
			<input type="password" name="password" placeholder="password" required><br>
			<input type="text" name="firstname" placeholder="first name" required><br>
			<input type="text" name="lastname" placeholder="last name" required><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}

// Login route
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			w.Write([]byte("Username must not be empty"))
			return
		} else if password == "" {
			w.Write([]byte("Password must not be empty"))
			return
		}

		if regex.MatchString(username) {
			if !database.DataBase.EmailExists(username) {
				w.Write([]byte("Invalid email"))
				return
			}
			user := database.DataBase.GetUserByEmail(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			errorhandle.Check(err)
			sessions.CreateSession(user.UserID, w)
		} else {
			if !database.DataBase.UsernameExists(username) {
				w.Write([]byte("Invalid username"))
				return
			}
			user := database.DataBase.GetUserByUsername(username)
			err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			errorhandle.Check(err)
			sessions.CreateSession(user.UserID, w)
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST">
			<input type="text" name="username" placeholder="username or email" required><br>
			<input type="password" name="password" placeholder="password" required><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}
