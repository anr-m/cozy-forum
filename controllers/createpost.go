package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"regexp"
	"time"

	"../db"
	"../errorhandle"
	"../models"
	"../sessions"
	"../tpl"
)

// CreatePost route
func CreatePost(w http.ResponseWriter, r *http.Request) {

	data := pageData{"Create Post", sessions.GetUser(w, r), nil}

	if r.Method == http.MethodPost {
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
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if !catregex.MatchString(category) {
			data.Data = "Invalid category"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if title == "" {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if content == "" {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif"
			w.WriteHeader(http.StatusUnprocessableEntity)
			tpl.ExecuteTemplate(w, "createpost.html", data)
			return
		}

		var encodedImage string
		if fh != nil {
			var buff bytes.Buffer
			img, ext, err := image.Decode(mf)
			errorhandle.Check(err)
			switch ext {
			case "png":
				png.Encode(&buff, img)
			case "jpg", "jpeg":
				jpeg.Encode(&buff, img, nil)
			case "gif":
				mf.Seek(0, io.SeekStart)
				gifimg, err := gif.DecodeAll(mf)
				errorhandle.Check(err)
				gif.EncodeAll(&buff, gifimg)
			}
			png.Encode(&buff, img)
			encodedImage = fmt.Sprintf(`<img src="data:image/%v;base64, %v"/>`, ext, base64.StdEncoding.EncodeToString(buff.Bytes()))
		}

		newPost := models.Post{
			UserID:      sessions.GetUser(w, r).UserID,
			Category:    category,
			Title:       title,
			Content:     content,
			HTMLImage:   template.HTML(encodedImage),
			TimeCreated: time.Now(),
		}

		db.CreatePost(&newPost)

		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "createpost.html", data)
	}
}
