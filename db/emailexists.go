package db

import "../errorhandle"

// EmailExists checks if email is already in database
func EmailExists(email string) bool {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	errorhandle.Check(err)
	return row.Next()
}
