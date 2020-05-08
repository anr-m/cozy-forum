package db

import (
	"log"

	"../models"
)

// LikePost ...
func LikePost(postid int, userid int) error {
	log.Printf("Creating new like from userid %d for postid %d...\n", userid, postid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 1);
	`)
	if err != nil {
		return err
	}

	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, userid, postid)
	defer exists.Close()

	if err != nil {
		return err
	}

	liked := 2
	for exists.Next() {
		exists.Scan(&userid, &postid, &liked)
	}

	if liked == 0 {
		sqlCommand, err = db.Prepare(`
			UPDATE postlikes
			SET liked = 1
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 1 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}

	_, err = sqlCommand.Exec(
		userid,
		postid,
	)
	if err != nil {
		return err
	}

	log.Printf("Created a new like from userid %d for postid %d\n", userid, postid)

	return nil
}

// DislikePost ...
func DislikePost(postid int, userid int) error {
	log.Printf("Creating new dislike from userid %d for postid %d...\n", userid, postid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 0);
	`)
	if err != nil {
		return err
	}

	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, userid, postid)
	defer exists.Close()

	if err != nil {
		return err
	}

	liked := 2
	for exists.Next() {
		exists.Scan(&userid, &postid, &liked)
	}

	if liked == 1 {
		sqlCommand, err = db.Prepare(`
			UPDATE postlikes
			SET liked = 0
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 0 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}

	_, err = sqlCommand.Exec(
		userid,
		postid,
	)
	if err != nil {
		return err
	}

	log.Printf("Created a new dislike from userid %d for postid %d\n", userid, postid)

	return nil
}

// LikeComment ...
func LikeComment(commentid int, userid int) error {
	log.Printf("Creating new like from userid %d for commentid %d...\n", userid, commentid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, commentid, liked)
		VALUES (?, ?, 1);
	`)
	if err != nil {
		return err
	}

	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ?)
	`, userid, commentid)
	defer exists.Close()

	if err != nil {
		return err
	}

	liked := 2
	for exists.Next() {
		exists.Scan(&userid, &commentid, &liked)
	}

	if liked == 0 {
		sqlCommand, err = db.Prepare(`
			UPDATE commentlikes
			SET liked = 1
			WHERE (userid = ? AND commentid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 1 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM commentlikes
			WHERE (userid = ? AND commentid = ?)
		`)
		if err != nil {
			return err
		}
	}

	_, err = sqlCommand.Exec(
		userid,
		commentid,
	)
	if err != nil {
		return err
	}

	log.Printf("Created a new like from userid %d for commentid %d\n", userid, commentid)

	return nil
}

// DislikeComment ...
func DislikeComment(commentid int, userid int) error {
	log.Printf("Creating new dislike from userid %d for commentid %d...\n", userid, commentid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, commentid, liked)
		VALUES (?, ?, 0);
	`)
	if err != nil {
		return err
	}

	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ?)
	`, userid, commentid)
	defer exists.Close()

	if err != nil {
		return err
	}

	liked := 2
	for exists.Next() {
		exists.Scan(&userid, &commentid, &liked)
	}

	if liked == 1 {
		sqlCommand, err = db.Prepare(`
			UPDATE commentlikes
			SET liked = 0
			WHERE (userid = ? AND commentid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 0 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM commentlikes
			WHERE (userid = ? AND commentid = ?)
		`)
		if err != nil {
			return err
		}
	}

	_, err = sqlCommand.Exec(
		userid,
		commentid,
	)
	if err != nil {
		return err
	}

	log.Printf("Created a new dislike from userid %d for commentid %d\n", userid, commentid)

	return nil
}

func getPostLikesAndDislikes(post *models.Post) error {
	likes, err := db.Query(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (postid = ? AND liked = 1)
	`, post.PostID)
	defer likes.Close()

	if err != nil {
		return err
	}

	for likes.Next() {
		likes.Scan(&post.Like)
	}

	dislikes, err := db.Query(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (postid = ? AND liked = 0)
	`, post.PostID)
	defer dislikes.Close()

	if err != nil {
		return err
	}

	for dislikes.Next() {
		dislikes.Scan(&post.Dislike)
	}

	return nil
}

func getCommentLikesAndDislikes(comment *models.Comment) error {
	likes, err := db.Query(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 1)
	`, comment.CommentID)
	defer likes.Close()

	if err != nil {
		return err
	}

	for likes.Next() {
		likes.Scan(&comment.Like)
	}

	dislikes, err := db.Query(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 0)
	`, comment.CommentID)
	defer dislikes.Close()

	if err != nil {
		return err
	}

	for dislikes.Next() {
		dislikes.Scan(&comment.Dislike)
	}

	return nil
}
