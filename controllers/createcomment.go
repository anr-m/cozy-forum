package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cozy-forum/db"
	"cozy-forum/models"
)

// CreateComment route for creating comments
func CreateComment(w http.ResponseWriter, r *http.Request, user models.User) {
	text := r.FormValue("text")
	postid, _ := strconv.Atoi(r.FormValue("postid"))

	if text == "" || postid == 0 {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	now := time.Now()

	newComment := models.Comment{
		PostID:      postid,
		Username:    user.Username,
		Text:        text,
		TimeCreated: now,
		TimeString:  now.Format("2006-01-02 15:04"),
	}

	err := db.CreateComment(&newComment)

	if internalError(w, r, err) {
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/posts/id/%d#%d", newComment.PostID, newComment.CommentID), http.StatusSeeOther)
}
