package db

import (
	"database/sql"
	"log"

	"../errorhandle"
	"../models"

	// Import the driver only
	_ "github.com/mattn/go-sqlite3"
)

// Database file
var db *sql.DB

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
		postid   INTEGER NOT NULL,
		liked    INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating commentlikes table...")
	createCommentLikesTable, err := db.Prepare(createCommentLikesTableSQL)
	errorhandle.Check(err)
	_, err = createCommentLikesTable.Exec()
	errorhandle.Check(err)
	log.Println("Commentlikes table created")
}

// CreateUser to create a new user
func CreateUser(newUser *models.User) {
	log.Println("Creating new user...")
	createUser, err := db.Prepare(`
		INSERT INTO users
		(hash, salt, firstname, lastname, username, email)
		VALUES (?, ?, ?, ?, ?, ?);
	`)
	errorhandle.Check(err)
	res, err := createUser.Exec(
		newUser.Hash,
		newUser.Salt,
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

// CreateSession to create a new session
func CreateSession(newSession models.Session) {
	log.Printf("Creating new session for userid %d...\n", newSession.UserID)
	createSession, err := db.Prepare(`
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
func UpdatePost(updatedPost models.Post) {
	// TODO update all values
	// log.Printf("Updating post with id %d...\n", updatedPost.PostID)
	// updatePost, err := db.Prepare(`
	// 	UPDATE posts
	// 	SET like = ?, dislike = ?
	// 	WHERE postid = ?;
	// `)

	// errorhandle.Check(err)
	// _, err = updatePost.Exec(
	// 	updatedPost.Like,
	// 	updatedPost.Dislike,
	// 	updatedPost.PostID,
	// )
	// errorhandle.Check(err)
	// log.Printf("Updated post with id %d\n", updatedPost.PostID)
}

// UpdateComment to update a comment
func UpdateComment(updatedComment models.Comment) {
	// TODO update all values
	// log.Printf("Updating comment with id %d...\n", updatedComment.CommentID)
	// updateComment, err := db.Prepare(`
	// 	UPDATE comments
	// 	SET like = ?, dislike = ?
	// 	WHERE commentid = ?;
	// `)

	// errorhandle.Check(err)
	// _, err = updateComment.Exec(
	// 	updatedComment.Like,
	// 	updatedComment.Dislike,
	// 	updatedComment.CommentID,
	// )
	// errorhandle.Check(err)
	// log.Printf("Updated comment with id %d\n", updatedComment.CommentID)
}

// EmailExists checks if email is already in database
func EmailExists(email string) bool {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	errorhandle.Check(err)
	return row.Next()
}

// UsernameExists checks if email is already in database
func UsernameExists(username string) bool {
	row, err := db.Query(`
		SELECT 1
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	errorhandle.Check(err)
	return row.Next()
}

// GetSession looks up session by sessionID
func GetSession(sessionID string) models.Session {
	row, err := db.Query(`
		SELECT *
		FROM sessions
		WHERE sessionid = ?
	`, sessionID)
	defer row.Close()

	errorhandle.Check(err)

	var session models.Session

	for row.Next() {
		row.Scan(&session.SessionID, &session.UserID, &session.TimeCreated)
	}

	return session
}

// UpdateSession updates session's created time
func UpdateSession(updatedSession models.Session) {
	log.Printf("Updating session with id %v...\n", updatedSession.SessionID)
	updateSession, err := db.Prepare(`
		UPDATE sessions
		SET timecreated = ?
		WHERE sessionid = ?;
	`)

	errorhandle.Check(err)
	_, err = updateSession.Exec(
		updatedSession.TimeCreated,
		updatedSession.SessionID,
	)
	errorhandle.Check(err)
	log.Printf("Updated session with id %v\n", updatedSession.SessionID)
}

// GetUserByID gets the user from userid
func GetUserByID(userID int) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE userid = ?
	`, userID)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User

	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}

// GetUserByEmail gets the user from userid
func GetUserByEmail(email string) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE email = ?
	`, email)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}

// GetUserByUsername gets the user from userid
func GetUserByUsername(username string) models.User {
	row, err := db.Query(`
		SELECT *
		FROM users
		WHERE username = ?
	`, username)
	defer row.Close()

	errorhandle.Check(err)

	var user models.User
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email, &user.FirstName, &user.LastName)
	}

	return user
}

func GetPostByID(postid int) models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE postid = ?
	`, postid)
	defer row.Close()

	errorhandle.Check(err)

	var post models.Post
	for row.Next() {
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
	}

	getPostLikesAndDislikes(&post)

	return post
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

func GetPosts() []models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
	`)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}

func GetPostsByCategory(category string) []models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE category = ?
	`, category)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}

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

func getCommentLikesAndDislikes(comment *models.Comment) {
	likes, err := db.Query(`
	SELECT COUNT(*)
	FROM commentlikes
	WHERE (postid = ? AND liked = 1)
`, comment.CommentID)
	defer likes.Close()

	errorhandle.Check(err)

	for likes.Next() {
		likes.Scan(&comment.Like)
	}

	dislikes, err := db.Query(`
	SELECT COUNT(*)
	FROM commentlikes
	WHERE (postid = ? AND liked = 0)
`, comment.CommentID)
	defer dislikes.Close()

	errorhandle.Check(err)

	for dislikes.Next() {
		dislikes.Scan(&comment.Dislike)
	}
}

func Close() {
	db.Close()
}
