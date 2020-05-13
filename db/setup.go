package db

import (
	"database/sql"
	"log"

	"cozy-forum/errorhandle"

	// Import the driver only
	_ "github.com/mattn/go-sqlite3"
)

// SetUp for setting up DB
func SetUp() {
	var err error
	log.Println("Opening DB...")
	db, err = sql.Open("sqlite3", "./database.db")
	errorhandle.Fatal(err)
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
	errorhandle.Fatal(err)
	_, err = createUserTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Users table created")

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS posts (
		postid      INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userid      INTEGER NOT NULL,
		username    TEXT NOT NULL,
		title       TEXT NOT NULL,
		content     TEXT NOT NULL,
		timecreated TIMESTAMP NOT NULL,
		timestring  TEXT NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (username) REFERENCES users(username)
	);`

	log.Println("Creating posts table...")
	createPostTable, err := db.Prepare(createPostTableSQL)
	errorhandle.Fatal(err)
	_, err = createPostTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Posts table created")

	createCategoryTableSQL := `CREATE TABLE IF NOT EXISTS categories (
		categoryid    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		categoryname  TEXT NOT NULL UNIQUE
	);`

	log.Println("Creating categories table...")
	createCategoryTable, err := db.Prepare(createCategoryTableSQL)
	errorhandle.Fatal(err)
	_, err = createCategoryTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Categories table created")

	log.Println("Creating default categories...")

	defaultCategories := []string{
		"Technology",
		"Programming",
		"Gaming",
		"Music",
		"Books",
		"Movies",
	}

	for _, category := range defaultCategories {
		err = CreateCategory(category)
		if err != nil && err != ErrAlreadyExists {
			errorhandle.Fatal(err)
		}
	}

	log.Println("Default categories created...")

	createPostCategoryTableSQL := `CREATE TABLE IF NOT EXISTS postcategories (
		categoryid  INTEGER NOT NULL,
		postid      INTEGER NOT NULL,
		FOREIGN KEY (categoryid) REFERENCES categories(categoryid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating postcategories table...")
	createPostCategoryTable, err := db.Prepare(createPostCategoryTableSQL)
	errorhandle.Fatal(err)
	_, err = createPostCategoryTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Postcategories table created")

	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS comments (
		commentid   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postid      INTEGER NOT NULL,
		username    TEXT NOT NULL,
		text  	    TEXT NOT NULL,
		timecreated TIMESTAMP NOT NULL,
		timestring  TEXT NOT NULL,
		FOREIGN KEY (username) REFERENCES users(username),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`

	log.Println("Creating comments table...")
	createCommentTable, err := db.Prepare(createCommentTableSQL)
	errorhandle.Fatal(err)
	_, err = createCommentTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Comments table created")

	createSessionTableSQL := `CREATE TABLE IF NOT EXISTS sessions (
		sessionid   STRING NOT NULL PRIMARY KEY,
		userid      INTEGER NOT NULL UNIQUE,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	log.Println("Creating sessions table...")
	createSessionTable, err := db.Prepare(createSessionTableSQL)
	errorhandle.Fatal(err)
	_, err = createSessionTable.Exec()
	errorhandle.Fatal(err)
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
	errorhandle.Fatal(err)
	_, err = createPostLikesTable.Exec()
	errorhandle.Fatal(err)
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
	errorhandle.Fatal(err)
	_, err = createCommentLikesTable.Exec()
	errorhandle.Fatal(err)
	log.Println("Commentlikes table created")
}
