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
	Title       string
	Content     template.HTML
	Categories  []string
	ImageExist  bool
	Like        int
	Dislike     int
	Liked       bool
	Disliked    bool
	TimeCreated time.Time
	TimeString  string
}
