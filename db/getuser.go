package db

import (
	"../errorhandle"
	"../models"
)

// GetUserByID gets the user from userid
func GetUserByID(userID int) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE userid = ?
	`, userID)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User

	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}

// GetUserByEmail gets the user from userid
func GetUserByEmail(email string) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}

// GetUserByUsername gets the user from userid
func GetUserByUsername(username string) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}
