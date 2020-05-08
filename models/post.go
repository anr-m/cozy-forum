package models

import (
	"html/template"
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
	HTMLImage   template.HTML
	Like        int
	Dislike     int
	Liked       bool
	Disliked    bool
	TimeCreated time.Time
	TimeString  string
}
