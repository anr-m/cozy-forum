package sessions

import (
	"net/http"
	"time"

	"../db"
	"../errorhandle"
	"../models"
	uuid "github.com/satori/go.uuid"
)

// CreateSession creates session for the userid
func CreateSession(userID int, w http.ResponseWriter) {
	sID, err := uuid.NewV4()
	errorhandle.Check(err)

	c := &http.Cookie{
		Name:   "session",
		Value:  sID.String(),
		MaxAge: 24 * 60 * 60,
		Path:   "/",
	}

	session := models.Session{
		SessionID:   c.Value,
		UserID:      userID,
		TimeCreated: time.Now(),
	}

	db.CreateSession(session)
	http.SetCookie(w, c)
}

// GetUser gets user from the database from the session
func GetUser(w http.ResponseWriter, r *http.Request) models.User {
	var user models.User

	c, err := r.Cookie("session")
	if err != nil {
		return user
	}

	session := db.GetSession(c.Value)
	if session.UserID == 0 {
		return user
	}

	user = db.GetUserByID(session.UserID)
	if user.UserID == 0 {
		return user
	}

	// session.TimeCreated = time.Now()
	// db.UpdateSession(session)

	c.Path = "/"
	c.MaxAge = 24 * 60 * 60
	http.SetCookie(w, c)

	return user
}

// IsLoggedIn checks if user is logged in
func IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	user := GetUser(w, r)
	return user.UserID != 0
}

// Logout to log the user out
func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
}
