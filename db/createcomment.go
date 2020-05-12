package db

import (
	"log"

	"cozy-forum/models"
)

// CreateComment to create a new comment
func CreateComment(newComment *models.Comment) error {
	log.Printf("Creating new comment from username %v for post %d...\n", newComment.Username, newComment.PostID)
	createComment, err := db.Prepare(`
		INSERT INTO comments
		(username, postid, text, timecreated, timestring)
		VALUES (?, ?, ?, ?, ?);
	`)

	if err != nil {
		return err
	}

	res, err := createComment.Exec(
		newComment.Username,
		newComment.PostID,
		newComment.Text,
		newComment.TimeCreated,
		newComment.TimeString,
	)

	if err != nil {
		return err
	}

	commentid, _ := res.LastInsertId()
	newComment.CommentID = int(commentid)
	log.Printf("Created a new comment with id %d\n", commentid)

	return nil
}
