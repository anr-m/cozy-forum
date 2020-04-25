package controllers

import (
	"net/http"
	"regexp"

	"../database"
	"../models"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser models.User
		regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		r.ParseForm()

		newUser.Email = r.FormValue("email")
		newUser.Username = r.FormValue("username")
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")

		if newUser.Email == "" {
			w.Write([]byte("Email must not be empty"))
		} else if !regex.MatchString(newUser.Email) {
			w.Write([]byte("Enter valid email"))
		} else if database.DataBase.EmailExists(newUser.Email) {
			w.Write([]byte("Email exists"))
		} else if newUser.Username == "" {
			w.Write([]byte("Username must not be empty"))
		} else if database.DataBase.UsernameExists(newUser.Username) {
			w.Write([]byte("Username exists"))
		} else if newUser.FirstName == "" {
			w.Write([]byte("First Name must not be empty"))
		} else if newUser.LastName == "" {
			w.Write([]byte("Last Name must not be empty"))
		} else {
			w.Write([]byte("Everything is ok"))
		}
	} else if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<form method="POST">
			<input type="text" name="username" placeholder="username" required><br>
			<input type="email" name="email" placeholder="email" required><br>
			<input type="text" name="firstname" placeholder="first name" required><br>
			<input type="text" name="lastname" placeholder="last name" required><br>
			<button type="submit">Submit</button>
		</form>`))
	}
}
