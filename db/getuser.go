package db

import (
	"cozy-forum/models"
)

// GetUserByID gets the user from userid
func GetUserByID(userID int) (models.User, error) {

	var user models.User

	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE userid = ?
	`, userID)
	defer row.Close()

	if err != nil {
		return user, err
	}

	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user, nil
}

// GetUserByEmail gets the user from userid
func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	if err != nil {
		return user, err
	}

	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user, nil
}

// GetUserByUsername gets the user from userid
func GetUserByUsername(username string) (models.User, error) {

	var user models.User

	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	if err != nil {
		return user, err
	}

	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user, nil
}
