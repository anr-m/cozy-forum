package db

import (
	"../models"
)

// GetSession looks up session by sessionID
func GetSession(sessionID string) (models.Session, error) {

	var session models.Session

	row, err := db.Query(`
		SELECT *
		FROM sessions
		WHERE sessionid = ?
	`, sessionID)
	defer row.Close()

	if err != nil {
		return session, err
	}

	for row.Next() {
		row.Scan(&session.SessionID, &session.UserID, &session.TimeCreated)
	}

	return session, nil
}
