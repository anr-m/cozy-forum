package db

import (
	"../errorhandle"
	"../models"
)

// GetSession looks up session by sessionID
func GetSession(sessionID string) models.Session {
	row, err := db.Query(`
		SELECT *
		FROM sessions
		WHERE sessionid = ?
	`, sessionID)
	defer row.Close()

	errorhandle.Check(err)

	var session models.Session

	for row.Next() {
		row.Scan(&session.SessionID, &session.UserID, &session.TimeCreated)
	}

	return session
}
