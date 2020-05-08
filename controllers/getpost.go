package controllers

import (
	"net/http"
	"path"
	"regexp"
	"strconv"

	"../db"
	"../models"
	"../tpl"
)

type postData struct {
	Link  string
	Posts []models.Post
}

// GetPosts route for browsing all posts
func GetPosts(w http.ResponseWriter, r *http.Request, user models.User) {
	posts, err := db.GetPosts(user.UserID)
	if internalError(w, r, err) {
		return
	}

	data := pageData{"All Posts", user, postData{r.URL.Path, posts}}
	internalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))
}

// GetPostByID route for getting a post by id
func GetPostByID(w http.ResponseWriter, r *http.Request, user models.User) {
	dir, endpoint := path.Split(r.URL.Path)
	postid, _ := strconv.Atoi(endpoint)
	data := pageData{"", user, nil}

	if dir != "/posts/id/" || postid == 0 {
		NotFoundHandler(w, r)
		return
	}

	post, err := db.GetPostByID(postid, user.UserID)
	if internalError(w, r, err) {
		return
	}

	if post.PostID != postid {
		NotFoundHandler(w, r)
		return
	}

	comments, err := db.GetCommentsByPostID(postid, user.UserID)
	if internalError(w, r, err) {
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

	internalError(w, r, tpl.ExecuteTemplate(w, "post.html", data))
}

// GetPostsByCategory route for browsing posts by category
func GetPostsByCategory(w http.ResponseWriter, r *http.Request, user models.User) {
	dir, category := path.Split(r.URL.Path)
	catregex := regexp.MustCompile(`^(Gaming|Technology|Programming|Books|Music)$`)

	if dir != "/posts/" || !catregex.MatchString(category) {
		NotFoundHandler(w, r)
		return
	}

	posts, err := db.GetPostsByCategory(category, user.UserID)
	if internalError(w, r, err) {
		return
	}

	data := pageData{category, user, postData{r.URL.Path, posts}}
	internalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))
}

// GetMyPosts ...
func GetMyPosts(w http.ResponseWriter, r *http.Request, user models.User) {
	posts, err := db.GetPostsByUserID(user.UserID)
	if internalError(w, r, err) {
		return
	}

	data := pageData{"My Posts", user, postData{r.URL.Path, posts}}
	internalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))
}

// GetMyLikedPosts ...
func GetMyLikedPosts(w http.ResponseWriter, r *http.Request, user models.User) {
	posts, err := db.GetLikedPostsByUserID(user.UserID)
	if internalError(w, r, err) {
		return
	}

	data := pageData{"My Posts", user, postData{r.URL.Path, posts}}
	internalError(w, r, tpl.ExecuteTemplate(w, "posts.html", data))
}
