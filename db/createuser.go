package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// CreateUser to create a new user
func CreateUser(newUser *models.User) {
	log.Println("Creating new user...")
	createUser, err := db.Prepare(`
		INSERT INTO users
		(hash, salt, firstname, lastname, username, email)
		VALUES (?, ?, ?, ?, ?, ?);
	`)
	errorhandle.Check(err)
	res, err := createUser.Exec(
		newUser.Hash,
		newUser.Salt,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
	)
	errorhandle.Check(err)
	userid, _ := res.LastInsertId()
	newUser.UserID = int(userid)
	log.Printf("Created a new user with id %d\n", userid)
}
