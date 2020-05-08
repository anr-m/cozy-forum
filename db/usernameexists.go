package db

// UsernameExists checks if email is already in database
func UsernameExists(username string) (bool, error) {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	if err != nil {
		return false, err
	}

	return row.Next(), nil
}
