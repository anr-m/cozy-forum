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
