package sessions

import (
	"net/http"
	"time"

	"../database"
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
	}

	session := models.Session{
		SessionID:   c.Value,
		UserID:      userID,
		TimeCreated: time.Now(),
	}

	database.DataBase.CreateSession(session)
	http.SetCookie(w, c)
}

// GetUser gets user from the database from the session
func GetUser(w http.ResponseWriter, r *http.Request) models.User {
	var user models.User

	c, err := r.Cookie("session")
	if err != nil {
		return user
	}

	session := database.DataBase.GetSession(c.Value)
	if session.UserID == 0 {
		return user
	}

	user = database.DataBase.GetUserByID(session.UserID)
	if user.UserID == 0 {
		return user
	}

	// session.TimeCreated = time.Now()
	// database.DataBase.UpdateSession(session)

	c.MaxAge = 24 * 60 * 60
	http.SetCookie(w, c)

	return user
}

// IsLoggedIn checks if user is logged in
func IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	user := GetUser(w, r)
	return user.UserID != 0
}
