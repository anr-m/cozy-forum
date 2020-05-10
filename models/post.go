package models

import (
	"time"
)

// Post type for DB
type Post struct {
	PostID      int
	UserID      int
	Username    string
	Category    string
	Title       string
	Content     string
	Like        int
	Dislike     int
	Liked       bool
	Disliked    bool
	TimeCreated time.Time
	TimeString  string
}
