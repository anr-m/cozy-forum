package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// UpdateSession updates session's created time
func UpdateSession(updatedSession models.Session) {
	log.Printf("Updating session with id %v...\n", updatedSession.SessionID)
	updateSession, err := db.Prepare(`
		UPDATE sessions
		SET timecreated = ?
		WHERE sessionid = ?;
	`)

	errorhandle.Check(err)
	_, err = updateSession.Exec(
		updatedSession.TimeCreated,
		updatedSession.SessionID,
	)
	errorhandle.Check(err)
	log.Printf("Updated session with id %v\n", updatedSession.SessionID)
}
