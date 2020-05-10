package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"../db"
	"../models"
	"../tpl"
)

// CreatePost route
func CreatePost(w http.ResponseWriter, r *http.Request, user models.User) {

	data := pageData{"Create Post", user, nil}

	if r.Method == http.MethodPost {
		var err error
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG)$`)
		catregex := regexp.MustCompile(`^(Gaming|Technology|Programming|Books|Music)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}

		category := r.FormValue("category")
		title := r.FormValue("title")
		content := r.FormValue("content")

		if category == "" {
			data.Data = "Category must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if !catregex.MatchString(category) {
			data.Data = "Invalid category"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if title == "" {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if content == "" {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif"
			w.WriteHeader(http.StatusUnprocessableEntity)
			internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		}

		now := time.Now()

		newPost := models.Post{
			UserID:      user.UserID,
			Username:    user.Username,
			Category:    category,
			Title:       title,
			Content:     content,
			TimeCreated: now,
			TimeString:  now.Format("2006-01-02 15:04"),
		}

		err = db.CreatePost(&newPost)
		if internalError(w, r, err) {
			return
		}

		if fh != nil {
			f, err := os.Create(fmt.Sprintf("./static/images/%v", newPost.PostID))
			defer f.Close()
			_, err = io.Copy(f, mf)
			if internalError(w, r, err) {
				return
			}
		}

		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		internalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
	}
}
