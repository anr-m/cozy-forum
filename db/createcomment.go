package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// CreateComment to create a new comment
func CreateComment(newComment *models.Comment) {
	log.Printf("Creating new comment from username %v for post %d...\n", newComment.Username, newComment.PostID)
	createComment, err := db.Prepare(`
		INSERT INTO comments
		(username, postid, text, timecreated)
		VALUES (?, ?, ?, ?);
	`)

	errorhandle.Check(err)
	res, err := createComment.Exec(
		newComment.Username,
		newComment.PostID,
		newComment.Text,
		newComment.TimeCreated,
	)
	errorhandle.Check(err)
	commentid, _ := res.LastInsertId()
	newComment.CommentID = int(commentid)
	log.Printf("Created a new comment with id %d\n", commentid)
}
