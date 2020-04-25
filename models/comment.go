package models

import "time"

// Comment type for DB
type Comment struct {
	CommentID   int
	PostID      int
	UserID      int
	Text        string
	Like        int
	Dislike     int
	TimeCreated time.Time
}
