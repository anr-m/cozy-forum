package sessions

import (
	"net/http"
	"time"

	"../db"
	"../models"
	uuid "github.com/satori/go.uuid"
)

// CreateSession creates session for the userid
func CreateSession(userID int, w http.ResponseWriter) error {
	sID, err := uuid.NewV4()
	if err != nil {
		return err
	}

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

	err = db.CreateSession(session)
	if err != nil {
		return err
	}

	http.SetCookie(w, c)

	return nil
}

// GetUser gets user from the database from the session
func GetUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User

	c, err := r.Cookie("session")
	if err != nil {
		return user, nil
	}

	session, err := db.GetSession(c.Value)
	if session.UserID == 0 || err != nil {
		return user, err
	}

	user, err = db.GetUserByID(session.UserID)
	if user.UserID == 0 || err != nil {
		return user, err
	}

	// session.TimeCreated = time.Now()
	// db.UpdateSession(session)

	c.Path = "/"
	c.MaxAge = 24 * 60 * 60
	http.SetCookie(w, c)

	return user, nil
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
