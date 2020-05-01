package models

import "time"

// Comment type for DB
type Comment struct {
	CommentID   int
	PostID      int
	Username    string
	Text        string
	Like        int
	Dislike     int
	TimeCreated time.Time
}
