package models

import (
	"time"
)

// Post type for DB
type Post struct {
	PostID      int
	UserID      int
	Username    string
	Title       string
	Content     string
	Categories  []string
	Like        int
	Dislike     int
	Liked       bool
	Disliked    bool
	TimeCreated time.Time
	TimeString  string
}
