package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// CreateSession to create a new session
func CreateSession(newSession models.Session) {
	log.Printf("Creating new session for userid %d...\n", newSession.UserID)
	createSession, err := db.Prepare(`
		INSERT INTO sessions
		(sessionid, userid, timecreated)
		VALUES (?, ?, ?);
	`)

	errorhandle.Check(err)
	_, err = createSession.Exec(
		newSession.SessionID,
		newSession.UserID,
		newSession.TimeCreated,
	)
	errorhandle.Check(err)
	log.Printf("Created a new session for userid %d with id %v\n", newSession.UserID, newSession.SessionID)
}
