package controllers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/tpl"
)

// CreatePost route
func CreatePost(w http.ResponseWriter, r *http.Request, data models.PageData) {

	data.PageTitle = "Create Post"

	if r.Method == http.MethodPost {
		var err error
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}

		categories := r.Form["categories"]
		categoryexist := make(map[string]bool)
		title := r.FormValue("title")
		content := r.FormValue("content")

		for _, category := range data.Categories {
			categoryexist[category] = true
		}

		for _, category := range categories {
			if !categoryexist[category] {
				data.Data = "Invalid category " + category
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
				return
			}
		}

		if len(categories) == 0 {
			data.Data = "Categories must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if title == "" {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if content == "" {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		}

		now := time.Now()

		newPost := models.Post{
			UserID:      data.User.UserID,
			Username:    data.User.Username,
			Categories:  categories,
			Title:       title,
			Content:     template.HTML(strings.ReplaceAll(content, "\n", "<br>")),
			TimeCreated: now,
			TimeString:  now.Format("2006-01-02 15:04"),
		}

		err = db.CreatePost(&newPost)
		if InternalError(w, r, err) {
			return
		}

		if fh != nil {
			f, err := os.Create(fmt.Sprintf("./static/images/%v", newPost.PostID))
			defer f.Close()
			_, err = io.Copy(f, mf)
			if InternalError(w, r, err) {
				return
			}
		}

		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
	}
}
