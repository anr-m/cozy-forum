package models

import "time"

// Post type for DB
type Post struct {
	PostID      int
	UserID      int
	Category    string
	Title       string
	Content     string
	Image       []byte
	Like        int
	Dislike     int
	TimeCreated time.Time
}
