package controllers

import (
	"net/http"
	"strconv"

	"../db"
	"../sessions"
)

// LikeComment route for liking and disliking comments
func LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	commentid, _ := strconv.Atoi(r.FormValue("commentid"))
	liked := r.FormValue("submit")
	link := r.FormValue("link")

	if commentid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	if liked == "like" {
		db.LikeComment(commentid, sessions.GetUser(w, r).UserID)
	} else {
		db.DislikeComment(commentid, sessions.GetUser(w, r).UserID)
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
