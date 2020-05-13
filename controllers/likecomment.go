package controllers

import (
	"net/http"
	"strconv"

	"cozy-forum/db"
	"cozy-forum/models"
)

// LikeComment route for liking and disliking comments
func LikeComment(w http.ResponseWriter, r *http.Request, data models.PageData) {
	var err error
	commentid, _ := strconv.Atoi(r.FormValue("commentid"))
	liked := r.FormValue("submit")
	link := r.FormValue("link")

	if commentid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
		ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
		return
	}

	if liked == "like" {
		err = db.LikeComment(commentid, data.User.UserID)
	} else {
		err = db.DislikeComment(commentid, data.User.UserID)
	}

	if InternalError(w, r, err) {
		return
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}
