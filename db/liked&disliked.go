package db

import "cozy-forum/models"

func postLikedByUser(post *models.Post, userid int) error {
	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ? AND liked = 1)
	`, userid, post.PostID)
	defer exists.Close()

	if err != nil {
		return err
	}

	post.Liked = exists.Next()

	return nil
}

func postDislikedByUser(post *models.Post, userid int) error {
	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ? AND liked = 0)
	`, userid, post.PostID)
	defer exists.Close()

	if err != nil {
		return err
	}

	post.Disliked = exists.Next()

	return nil
}

func commentLikedByUser(comment *models.Comment, userid int) error {
	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ? AND liked = 1)
	`, userid, comment.CommentID)
	defer exists.Close()

	if err != nil {
		return err
	}

	comment.Liked = exists.Next()

	return nil
}

func commentDislikedByUser(comment *models.Comment, userid int) error {
	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ? AND liked = 0)
	`, userid, comment.CommentID)
	defer exists.Close()

	if err != nil {
		return err
	}

	comment.Disliked = exists.Next()

	return nil
}
