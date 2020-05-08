package db

import (
	"database/sql"
	"log"

	"../errorhandle"
)

// SetUp for setting up DB
func SetUp() {
	var err error
	log.Println("Opening DB...")
	db, err = sql.Open("sqlite3", "./database.db")
	errorhandle.Check(err)
	log.Println("DB opened")

	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		userid    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username  TEXT NOT NULL UNIQUE,
		hash      BLOB NOT NULL,
		salt      TEXT NOT NULL,
		email     TEXT NOT NULL UNIQUE,
		firstname TEXT NOT NULL,
		lastname  TEXT NOT NULL		
	);`

	log.Println("Creating users table...")
	createUserTable, err := db.Prepare(createUserTableSQL)
	errorhandle.Check(err)
	_, err = createUserTable.Exec()
	errorhandle.Check(err)
	log.Println("Users table created")

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS posts (
		postid      INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userid      INTEGER NOT NULL,
		category    TEXT NOT NULL,
		title       TEXT NOT NULL,
		content     TEXT NOT NULL,
		image       TEXT,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	log.Println("Creating posts table...")
	createPostTable, err := db.Prepare(createPostTableSQL)
	errorhandle.Check(err)
	_, err = createPostTable.Exec()
	errorhandle.Check(err)
	log.Println("Posts table created")

	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS comments (
		commentid   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postid      INTEGER NOT NULL,
		username    TEXT NOT NULL,
		text  	    TEXT NOT NULL,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (username) REFERENCES users(username),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating comments table...")
	createCommentTable, err := db.Prepare(createCommentTableSQL)
	errorhandle.Check(err)
	_, err = createCommentTable.Exec()
	errorhandle.Check(err)
	log.Println("Comments table created")

	createSessionTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		sessionid   STRING NOT NULL PRIMARY KEY,
		userid      INTEGER NOT NULL,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	log.Println("Creating sessions table...")
	createSessionTable, err := db.Prepare(createSessionTableSQL)
	errorhandle.Check(err)
	_, err = createSessionTable.Exec()
	errorhandle.Check(err)
	log.Println("Sessions table created")

	createPostLikesTableSQL := `CREATE TABLE IF NOT EXISTS postlikes (
		userid   INTEGER NOT NULL,
		postid   INTEGER NOT NULL,
		liked    INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating postlikes table...")
	createPostLikesTable, err := db.Prepare(createPostLikesTableSQL)
	errorhandle.Check(err)
	_, err = createPostLikesTable.Exec()
	errorhandle.Check(err)
	log.Println("Postlikes table created")

	createCommentLikesTableSQL := `CREATE TABLE IF NOT EXISTS commentlikes (
		userid   INTEGER NOT NULL,
		commentid   INTEGER NOT NULL,
		liked    INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (commentid) REFERENCES comments(commentid)
	);`

	log.Println("Creating commentlikes table...")
	createCommentLikesTable, err := db.Prepare(createCommentLikesTableSQL)
	errorhandle.Check(err)
	_, err = createCommentLikesTable.Exec()
	errorhandle.Check(err)
	log.Println("Commentlikes table created")
}
