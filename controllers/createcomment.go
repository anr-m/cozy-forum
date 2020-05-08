package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../db"
	"../models"
	"../sessions"
)

// CreateComment route for creating comments
func CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	text := r.FormValue("text")
	postid, _ := strconv.Atoi(r.FormValue("postid"))
	user := sessions.GetUser(w, r)

	if text == "" || postid == 0 {
		return
	}

	newComment := models.Comment{
		PostID:      postid,
		Username:    user.Username,
		Text:        text,
		TimeCreated: time.Now(),
	}

	db.CreateComment(&newComment)

	http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newComment.PostID), http.StatusSeeOther)
}
