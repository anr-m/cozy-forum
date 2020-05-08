package controllers

import (
	"net/http"
	"strconv"

	"../db"
	"../sessions"
)

// LikePost route for liking and disliking posts
func LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	postid, _ := strconv.Atoi(r.FormValue("postid"))
	liked := r.FormValue("submit")
	link := r.FormValue("link")

	if postid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	if liked == "like" {
		db.LikePost(postid, sessions.GetUser(w, r).UserID)
	} else {
		db.DislikePost(postid, sessions.GetUser(w, r).UserID)
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
