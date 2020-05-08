package db

import (
	"log"

	"../models"
)

// CreateUser to create a new user
func CreateUser(newUser *models.User) error {
	log.Println("Creating new user...")
	createUser, err := db.Prepare(`
		INSERT INTO users
		(hash, salt, firstname, lastname, username, email)
		VALUES (?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}

	res, err := createUser.Exec(
		newUser.Hash,
		newUser.Salt,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
	)
	if err != nil {
		return err
	}

	userid, err := res.LastInsertId()
	newUser.UserID = int(userid)

	if err != nil {
		return err
	}

	log.Printf("Created a new user with id %d\n", userid)

	return nil
}
