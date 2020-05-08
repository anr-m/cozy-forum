package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// CreatePost to create a new post
func CreatePost(newPost *models.Post) {
	log.Printf("Creating new post for userid %d...\n", newPost.UserID)
	createPost, err := db.Prepare(`
		INSERT INTO posts
		(userid, category, title, content, image, timecreated)
		VALUES (?, ?, ?, ?, ?, ?);
	`)

	errorhandle.Check(err)
	res, err := createPost.Exec(
		newPost.UserID,
		newPost.Category,
		newPost.Title,
		newPost.Content,
		newPost.HTMLImage,
		newPost.TimeCreated,
	)
	errorhandle.Check(err)
	postid, _ := res.LastInsertId()
	newPost.PostID = int(postid)
	log.Printf("Created a new post with id %d\n", postid)
}
