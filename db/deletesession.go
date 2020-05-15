package db

// DeleteSession ...
func DeleteSession(sessionid string) error {
	createUser, err := db.Prepare(`
		DELETE FROM sessions
		WHERE sessionid = ?
	`)
	if err != nil {
		return err
	}

	_, err = createUser.Exec(
		sessionid,
	)
	return err
}
