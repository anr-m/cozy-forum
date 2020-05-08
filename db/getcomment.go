package db

import (
	"../errorhandle"
	"../models"
)

// GetCommentsByPostID ...
func GetCommentsByPostID(postid int) []models.Comment {
	row, err := db.Query(`
		SELECT *
		FROM comments
		WHERE postid = ?
	`, postid)
	defer row.Close()

	errorhandle.Check(err)

	var comments []models.Comment

	for row.Next() {
		var comment models.Comment
		row.Scan(&comment.CommentID, &comment.PostID, &comment.Username, &comment.Text, &comment.TimeCreated)
		getCommentLikesAndDislikes(&comment)
		comments = append(comments, comment)
	}

	return comments
}
