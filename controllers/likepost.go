package controllers

import (
	"net/http"
	"strconv"

	"cozy-forum/db"
	"cozy-forum/models"
)

// LikePost route for liking and disliking posts
func LikePost(w http.ResponseWriter, r *http.Request, data models.PageData) {

	if r.Method == http.MethodPost {
		var err error
		postid, _ := strconv.Atoi(r.FormValue("postid"))
		liked := r.FormValue("submit")
		link := r.FormValue("link")

		if postid == 0 || !(liked == "like" || liked == "dislike") || link == "" {
			ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
			return
		}

		if liked == "like" {
			err = db.LikePost(postid, data.User.UserID)
		} else {
			err = db.DislikePost(postid, data.User.UserID)
		}

		if InternalError(w, r, err) {
			return
		}

		http.Redirect(w, r, link, http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
