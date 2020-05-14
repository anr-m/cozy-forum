package models

import (
	"html/template"
	"time"
)

// Comment type for DB
type Comment struct {
	CommentID   int
	PostID      int
	Username    string
	Text        template.HTML
	Like        int
	Dislike     int
	Liked       bool
	Disliked    bool
	TimeCreated time.Time
	TimeString  string
}
