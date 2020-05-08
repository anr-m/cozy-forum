package db

import "../errorhandle"

// UsernameExists checks if email is already in database
func UsernameExists(username string) bool {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	errorhandle.Check(err)
	return row.Next()
}
