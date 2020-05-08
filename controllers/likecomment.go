package controllers

import (
	"net/http"
	"strconv"

	"../db"
	"../models"
)

// LikeComment route for liking and disliking comments
func LikeComment(w http.ResponseWriter, r *http.Request, user models.User) {
	var err error
	commentid, _ := strconv.Atoi(r.FormValue("commentid"))
	liked := r.FormValue("submit")
	link := r.FormValue("link")

	if commentid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
		errorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	if liked == "like" {
		err = db.LikeComment(commentid, user.UserID)
	} else {
		err = db.DislikeComment(commentid, user.UserID)
	}

	if internalError(w, r, err) {
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
