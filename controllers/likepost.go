package controllers

import (
	"net/http"
	"strconv"

	"cozy-forum/db"
	"cozy-forum/models"
)

// LikePost route for liking and disliking posts
func LikePost(w http.ResponseWriter, r *http.Request, user models.User) {
	var err error
	postid, _ := strconv.Atoi(r.FormValue("postid"))
	liked := r.FormValue("submit")
	link := r.FormValue("link")

	if postid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	if liked == "like" {
		err = db.LikePost(postid, user.UserID)
	} else {
		err = db.DislikePost(postid, user.UserID)
	}

	if internalError(w, r, err) {
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
