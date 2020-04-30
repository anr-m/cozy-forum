package models

import (
	"html/template"
	"time"
)

// Post type for DB
type Post struct {
	PostID      int
	UserID      int
	Category    string
	Title       string
	Content     string
	HTMLImage   template.HTML
	Like        int
	Dislike     int
	TimeCreated time.Time
}
