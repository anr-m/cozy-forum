package controllers

import (
	"net/http"
	"path"
	"strconv"

	"cozy-forum/db"
	"cozy-forum/models"
	"cozy-forum/tpl"
)

type postData struct {
	Link  string
	Posts []models.Post
}

// GetPosts route for browsing all posts
func GetPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		posts, err := db.GetPosts(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = "All Posts"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

// GetPostByID route for getting a post by id
func GetPostByID(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		dir, endpoint := path.Split(r.URL.Path)
		postid, _ := strconv.Atoi(endpoint)

		if dir != "/posts/id/" || postid == 0 {
			NotFoundHandler(w, r)
			return
		}

		post, err := db.GetPostByID(postid, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		if post.PostID != postid {
			NotFoundHandler(w, r)
			return
		}

		comments, err := db.GetCommentsByPostID(postid, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = post.Title
		data.Data = struct {
			Post     models.Post
			Comments []models.Comment
		}{
			Post:     post,
			Comments: comments,
		}

		InternalError(w, r, tpl.ExecuteTemplate(w, "post.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

// GetPostsByCategory route for browsing posts by category
func GetPostsByCategory(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		dir, category := path.Split(r.URL.Path)

		if dir != "/posts/" {
			NotFoundHandler(w, r)
			return
		}

		for i := range data.Categories {
			if category == data.Categories[i] {
				break
			}
			if i == len(data.Categories)-1 {
				NotFoundHandler(w, r)
				return
			}
		}

		posts, err := db.GetPostsByCategory(category, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = category
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

// GetMyPosts ...
func GetMyPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		posts, err := db.GetPostsByUserID(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = "My Posts"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

// GetMyLikedPosts ...
func GetMyLikedPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		posts, err := db.GetLikedPostsByUserID(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = "My Likes"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
