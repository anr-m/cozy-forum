package database

import (
	"database/sql"
	"log"

	"../errorhandle"
	"../models"

	// Import the driver only
	_ "github.com/mattn/go-sqlite3"
)

// DB type for DB
type DB struct {
	db *sql.DB
}

// DataBase for
var DataBase DB

// SetUp for setting up DB
func SetUp() {
	log.Println("Opening DB...")
	db, err := sql.Open("sqlite3", "./database.db")
	errorhandle.Check(err)
	log.Println("DB opened")

	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"userid"    integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username"  TEXT NOT NULL UNIQUE,
		"hash"      TEXT NOT NULL,
		"email"     TEXT NOT NULL UNIQUE,
		"firstname" TEXT NOT NULL,
		"lastname"  TEXT NOT NULL		
	);`

	log.Println("Creating users table...")
	createUserTable, err := db.Prepare(createUserTableSQL)
	errorhandle.Check(err)
	_, err = createUserTable.Exec()
	errorhandle.Check(err)
	log.Println("User table created")

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS posts (
		"postid" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"userid"      integer NOT NULL,
		"category"    TEXT NOT NULL,
		"title"       TEXT NOT NULL,
		"content"     TEXT NOT NULL,
		"image"       BLOB,
		"like"        integer NOT NULL DEFAULT 0,
		"dislike"     integer NOT NULL DEFAULT 0,
		"timecreated" timestamp NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	log.Println("Creating posts table...")
	createPostTable, err := db.Prepare(createPostTableSQL)
	errorhandle.Check(err)
	_, err = createPostTable.Exec()
	errorhandle.Check(err)
	log.Println("Posts table created")

	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS comments (
		"commentid"   integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"postid"      integer NOT NULL,
		"userid"      integer NOT NULL,
		"text"  	  TEXT NOT NULL,
		"like"        integer NOT NULL DEFAULT 0,
		"dislike"     integer NOT NULL DEFAULT 0,
		"timecreated" timestamp NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating comments table...")
	createCommentTable, err := db.Prepare(createCommentTableSQL)
	errorhandle.Check(err)
	_, err = createCommentTable.Exec()
	errorhandle.Check(err)
	log.Println("Comments table created")

	createSessionTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		"sessionid"   string NOT NULL PRIMARY KEY,
		"userid"      integer NOT NULL,
		"timecreated" timestamp NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	log.Println("Creating sessions table...")
	createSessionTable, err := db.Prepare(createSessionTableSQL)
	errorhandle.Check(err)
	_, err = createSessionTable.Exec()
	errorhandle.Check(err)
	log.Println("Sessions table created")

	DataBase = DB{db}

	// return DB{db}
}

// CreateUser to create a new user
func (DB *DB) CreateUser(newUser *models.User) {
	log.Println("Creating new user...")
	createUser, err := DB.db.Prepare(`
		INSERT INTO users
		(hash, firstname, lastname, username, email)
		VALUES (?, ?, ?, ?, ?);
	`)
	errorhandle.Check(err)
	res, err := createUser.Exec(
		newUser.Hash,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
	)
	errorhandle.Check(err)
	userid, _ := res.LastInsertId()
	newUser.UserID = int(userid)
	log.Printf("Created a new user with id %d\n", userid)
}

// CreatePost to create a new post
func (DB *DB) CreatePost(newPost *models.Post) {
	log.Printf("Creating new post for userid %d...\n", newPost.UserID)
	createPost, err := DB.db.Prepare(`
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
		newPost.Image,
		newPost.TimeCreated,
	)
	errorhandle.Check(err)
	postid, _ := res.LastInsertId()
	newPost.PostID = int(postid)
	log.Printf("Created a new post with id %d\n", postid)
}

// CreateComment to create a new comment
func (DB *DB) CreateComment(newComment *models.Comment) {
	log.Printf("Creating new comment from userid %d for post %d...\n", newComment.UserID, newComment.PostID)
	createComment, err := DB.db.Prepare(`
		INSERT INTO comments
		(userid, postid, text, timecreated)
		VALUES (?, ?, ?, ?);
	`)

	errorhandle.Check(err)
	res, err := createComment.Exec(
		newComment.UserID,
		newComment.PostID,
		newComment.Text,
		newComment.TimeCreated,
	)
	errorhandle.Check(err)
	commentid, _ := res.LastInsertId()
	newComment.CommentID = int(commentid)
	log.Printf("Created a new comment with id %d\n", commentid)
}

// CreateSession to create a new session
func (DB *DB) CreateSession(newSession models.Session) {
	log.Printf("Creating new session for userid %d...\n", newSession.UserID)
	createSession, err := DB.db.Prepare(`
		INSERT INTO sessions
		(sessionid, userid, timecreated)
		VALUES (?, ?, ?);
	`)

	errorhandle.Check(err)
	_, err = createSession.Exec(
		newSession.SessionID,
		newSession.UserID,
		newSession.TimeCreated,
	)
	errorhandle.Check(err)
	log.Printf("Created a new session for userid %d with id %v\n", newSession.UserID, newSession.SessionID)
}

// UpdatePost to update a post
func (DB *DB) UpdatePost(updatedPost models.Post) {
	// TODO update all values
	log.Printf("Updating post with id %d...\n", updatedPost.PostID)
	updatePost, err := DB.db.Prepare(`
		UPDATE posts
		SET like = ?, dislike = ?
		WHERE postid = ?;
	`)

	errorhandle.Check(err)
	_, err = updatePost.Exec(
		updatedPost.Like,
		updatedPost.Dislike,
		updatedPost.PostID,
	)
	errorhandle.Check(err)
	log.Printf("Updated post with id %d\n", updatedPost.PostID)
}

// UpdateComment to update a comment
func (DB *DB) UpdateComment(updatedComment models.Comment) {
	// TODO update all values
	log.Printf("Updating comment with id %d...\n", updatedComment.CommentID)
	updateComment, err := DB.db.Prepare(`
		UPDATE comments
		SET like = ?, dislike = ?
		WHERE commentid = ?;
	`)

	errorhandle.Check(err)
	_, err = updateComment.Exec(
		updatedComment.Like,
		updatedComment.Dislike,
		updatedComment.CommentID,
	)
	errorhandle.Check(err)
	log.Printf("Updated comment with id %d\n", updatedComment.CommentID)
}

// EmailExists checks if email is already in database
func (DB *DB) EmailExists(email string) bool {
	emails, err := DB.db.Query(`
		SELECT 1
		FROM users
		WHERE email = ?
	`, email)
	defer emails.Close()

	errorhandle.Check(err)
	return emails.Next()
}

// UsernameExists checks if email is already in database
func (DB *DB) UsernameExists(username string) bool {
	usernames, err := DB.db.Query(`
		SELECT 1
		FROM users
		WHERE username = ?
	`, username)
	defer usernames.Close()

	errorhandle.Check(err)
	return usernames.Next()
}
