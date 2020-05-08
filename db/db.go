package db

import (
	"database/sql"

	// Import the driver only
	_ "github.com/mattn/go-sqlite3"
)

// Database file
var db *sql.DB
