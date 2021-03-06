package db

import (
	"cozy-forum/models"
)

// GetCommentsByPostID ...
func GetCommentsByPostID(postid int, userid int) ([]models.Comment, error) {

	var comments []models.Comment

	row, err := db.Query(`
		SELECT *
		FROM comments
		WHERE postid = ?
	`, postid)
	defer row.Close()

	if err != nil {
		return comments, err
	}

	for row.Next() {
		var comment models.Comment
		row.Scan(&comment.CommentID, &comment.PostID, &comment.Username, &comment.Text, &comment.ImageExist, &comment.TimeCreated, &comment.TimeString)
		err = getCommentLikesAndDislikes(&comment)
		if err != nil {
			return comments, err
		}
		if userid != 0 {
			err = commentLikedByUser(&comment, userid)
			if err != nil {
				return comments, err
			}
			err = commentDislikedByUser(&comment, userid)
			if err != nil {
				return comments, err
			}
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func getPostCommentCount(post *models.Post) error {
	row, err := db.Query(`
		SELECT COUNT(*)
		FROM comments
		WHERE postid = ?
	`, post.PostID)
	defer row.Close()

	if err != nil {
		return err
	}

	for row.Next() {
		row.Scan(&post.CommentCount)
	}

	return nil
}
