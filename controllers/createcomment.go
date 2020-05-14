package controllers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/tpl"
)

// CreateComment route for creating comments
func CreateComment(w http.ResponseWriter, r *http.Request, data models.PageData) {

	if r.Method == http.MethodPost {
		var err error
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}

		text := r.FormValue("text")
		postid, _ := strconv.Atoi(r.FormValue("postid"))

		if text == "" || postid == 0 {
			ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif, svg"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		}

		now := time.Now()

		newComment := models.Comment{
			PostID:      postid,
			Username:    data.User.Username,
			Text:        template.HTML(strings.ReplaceAll(text, "\n", "<br>")),
			TimeCreated: now,
			TimeString:  now.Format("2006-01-02 15:04"),
		}

		err = db.CreateComment(&newComment)

		if InternalError(w, r, err) {
			return
		}

		if fh != nil {

			f, err := os.Create(fmt.Sprintf("./static/images/%v", newComment.CommentID))
			defer f.Close()
			_, err = io.Copy(f, mf)
			if InternalError(w, r, err) {

				return
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d#%d", newComment.PostID, newComment.CommentID), http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
	}
}
