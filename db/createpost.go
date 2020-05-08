package db

import (
	"log"

	"../models"
)

// CreatePost to create a new post
func CreatePost(newPost *models.Post) error {
	log.Printf("Creating new post for userid %d...\n", newPost.UserID)
	createPost, err := db.Prepare(`
		INSERT INTO posts
		(userid, username, category, title, content, image, timecreated, timestring)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}

	res, err := createPost.Exec(
		newPost.UserID,
		newPost.Username,
		newPost.Category,
		newPost.Title,
		newPost.Content,
		newPost.HTMLImage,
		newPost.TimeCreated,
		newPost.TimeString,
	)
	if err != nil {
		return err
	}

	postid, err := res.LastInsertId()
	newPost.PostID = int(postid)
	if err != nil {
		return err
	}

	log.Printf("Created a new post with id %d\n", postid)

	return nil
}
