package db

import (
	"log"

	"../errorhandle"
	"../models"
)

// LikePost ...
func LikePost(postid int, userid int) {
	log.Printf("Creating new like from userid %d for postid %d...\n", userid, postid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 1);
	`)
	errorhandle.Check(err)

	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, userid, postid)
	defer exists.Close()

	errorhandle.Check(err)

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
		errorhandle.Check(err)
	} else if liked == 1 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		errorhandle.Check(err)
	}

	_, err = sqlCommand.Exec(
		userid,
		postid,
	)
	errorhandle.Check(err)
	log.Printf("Created a new like from userid %d for postid %d\n", userid, postid)
}

// DislikePost ...
func DislikePost(postid int, userid int) {
	log.Printf("Creating new dislike from userid %d for postid %d...\n", userid, postid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 0);
	`)
	errorhandle.Check(err)

	exists, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, userid, postid)
	defer exists.Close()

	errorhandle.Check(err)

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
		errorhandle.Check(err)
	} else if liked == 0 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		errorhandle.Check(err)
	}

	_, err = sqlCommand.Exec(
		userid,
		postid,
	)
	errorhandle.Check(err)
	log.Printf("Created a new dislike from userid %d for postid %d\n", userid, postid)
}

// LikeComment ...
func LikeComment(commentid int, userid int) {
	log.Printf("Creating new like from userid %d for commentid %d...\n", userid, commentid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, commentid, liked)
		VALUES (?, ?, 1);
	`)
	errorhandle.Check(err)

	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ?)
	`, userid, commentid)
	defer exists.Close()

	errorhandle.Check(err)

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
		errorhandle.Check(err)
	} else if liked == 1 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM commentlikes
			WHERE (userid = ? AND commentid = ?)
		`)
		errorhandle.Check(err)
	}

	_, err = sqlCommand.Exec(
		userid,
		commentid,
	)
	errorhandle.Check(err)
	log.Printf("Created a new like from userid %d for commentid %d\n", userid, commentid)
}

// DislikeComment ...
func DislikeComment(commentid int, userid int) {
	log.Printf("Creating new dislike from userid %d for commentid %d...\n", userid, commentid)
	sqlCommand, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, commentid, liked)
		VALUES (?, ?, 0);
	`)
	errorhandle.Check(err)

	exists, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ?)
	`, userid, commentid)
	defer exists.Close()

	errorhandle.Check(err)

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
		errorhandle.Check(err)
	} else if liked == 0 {
		sqlCommand, err = db.Prepare(`
			DELETE FROM commentlikes
			WHERE (userid = ? AND commentid = ?)
		`)
		errorhandle.Check(err)
	}

	_, err = sqlCommand.Exec(
		userid,
		commentid,
	)
	errorhandle.Check(err)
	log.Printf("Created a new dislike from userid %d for commentid %d\n", userid, commentid)
}

func getPostLikesAndDislikes(post *models.Post) {
	likes, err := db.Query(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (postid = ? AND liked = 1)
	`, post.PostID)
	defer likes.Close()

	errorhandle.Check(err)

	for likes.Next() {
		likes.Scan(&post.Like)
	}

	dislikes, err := db.Query(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (postid = ? AND liked = 0)
	`, post.PostID)
	defer dislikes.Close()

	errorhandle.Check(err)

	for dislikes.Next() {
		dislikes.Scan(&post.Dislike)
	}
}

func getCommentLikesAndDislikes(comment *models.Comment) {
	likes, err := db.Query(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 1)
	`, comment.CommentID)
	defer likes.Close()

	errorhandle.Check(err)

	for likes.Next() {
		likes.Scan(&comment.Like)
	}

	dislikes, err := db.Query(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 0)
	`, comment.CommentID)
	defer dislikes.Close()

	errorhandle.Check(err)

	for dislikes.Next() {
		dislikes.Scan(&comment.Dislike)
	}
}
