package db

import (
	"log"

	"cozy-forum/models"
)

// UpdateSession updates session's created time
func UpdateSession(updatedSession models.Session) error {
	log.Printf("Updating session with id %v...\n", updatedSession.SessionID)
	updateSession, err := db.Prepare(`
		UPDATE sessions
		SET timecreated = ?
		WHERE sessionid = ?;
	`)
	if err != nil {
		return err
	}

	_, err = updateSession.Exec(
		updatedSession.TimeCreated,
		updatedSession.SessionID,
	)
	if err != nil {
		return err
	}

	log.Printf("Updated session with id %v\n", updatedSession.SessionID)

	return nil
}
