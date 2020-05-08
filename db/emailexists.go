package db

// EmailExists checks if email is already in database
func EmailExists(email string) (bool, error) {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	if err != nil {
		return false, err
	}

	return row.Next(), nil
}
