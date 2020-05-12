package db

import (
	"log"

	"cozy-forum/models"
)

// CreateSession to create a new session
func CreateSession(newSession models.Session) error {
	log.Printf("Creating new session for userid %d...\n", newSession.UserID)
	createSession, err := db.Prepare(`
		REPLACE INTO sessions
		(sessionid, userid, timecreated)
		VALUES (?, ?, ?);
	`)

	if err != nil {
		return err
	}

	_, err = createSession.Exec(
		newSession.SessionID,
		newSession.UserID,
		newSession.TimeCreated,
	)
	if err != nil {
		return err
	}

	log.Printf("Created a new session for userid %d with id %v\n", newSession.UserID, newSession.SessionID)

	return nil
}
