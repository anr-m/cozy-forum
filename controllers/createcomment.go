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

		if postid == 0 {
			ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
			return
		} else if isEmpty(text) {
			ErrorHandler(w, r, http.StatusUnprocessableEntity, "Text must not be empty")
			return
		} else if fh != nil {
			if fh.Size > 20000000 {
				ErrorHandler(w, r, http.StatusUnprocessableEntity, "File too large, please limit size to 20MB")
				return
			} else if !regex.MatchString(fh.Filename) {
				ErrorHandler(w, r, http.StatusUnprocessableEntity, "Invalid file type, please upload jpg, jpeg, png, gif, svg")
				return
			}
		}

		now := time.Now()
		loc, err := time.LoadLocation("Asia/Almaty")
		if InternalError(w, r, err) {
			return
		}
		now = now.In(loc)

		newComment := models.Comment{
			PostID:      postid,
			Username:    data.User.Username,
			Text:        template.HTML(strings.ReplaceAll(text, "\n", "<br>")),
			ImageExist:  fh != nil,
			TimeCreated: now,
			TimeString:  now.Format("2006-01-02 15:04"),
		}

		err = db.CreateComment(&newComment)

		if InternalError(w, r, err) {
			return
		}

		if fh != nil {
			f, err := os.Create(fmt.Sprintf("./static/images/c%v", newComment.CommentID))
			defer f.Close()
			_, err = io.Copy(f, mf)
			if InternalError(w, r, err) {
				return
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d#%d", newComment.PostID, newComment.CommentID), http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
