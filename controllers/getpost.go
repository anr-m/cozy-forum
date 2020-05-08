package controllers

import (
	"net/http"
	"path"
	"regexp"
	"strconv"

	"../db"
	"../models"
	"../sessions"
	"../tpl"
)

type postData struct {
	Link  string
	Posts []models.Post
}

// GetPosts route for browsing all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := db.GetPosts()
	data := pageData{"All Posts", sessions.GetUser(w, r), postData{r.URL.Path, posts}}
	tpl.ExecuteTemplate(w, "posts.html", data)
}

// GetPostByID route for getting a post by id
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	dir, endpoint := path.Split(r.URL.Path)
	postid, _ := strconv.Atoi(endpoint)

	data := pageData{"", sessions.GetUser(w, r), nil}

	if dir != "/posts/id/" || postid == 0 {
		NotFoundHandler(w, r)
		return
	}

	post := db.GetPostByID(postid)

	if post.PostID != postid {
		NotFoundHandler(w, r)
		return
	}

	comments := db.GetCommentsByPostID(postid)

	data.PageTitle = post.Title
	data.Data = struct {
		Post     models.Post
		Comments []models.Comment
	}{
		Post:     post,
		Comments: comments,
	}

	tpl.ExecuteTemplate(w, "post.html", data)
}

// GetPostsByCategory route for browsing posts by category
func GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
	dir, category := path.Split(r.URL.Path)
	catregex := regexp.MustCompile(`^(Gaming|Technology|Programming|Books|Music)$`)

	if dir != "/posts/" || !catregex.MatchString(category) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid path/category"))
		return
	}

	posts := db.GetPostsByCategory(category)
	data := pageData{category, sessions.GetUser(w, r), postData{r.URL.Path, posts}}
	tpl.ExecuteTemplate(w, "posts.html", data)
}

// GetMyPosts ...
func GetMyPosts(w http.ResponseWriter, r *http.Request) {
	user := sessions.GetUser(w, r)
	posts := db.GetPostsByUserID(user.UserID)
	data := pageData{"My Posts", user, postData{r.URL.Path, posts}}
	tpl.ExecuteTemplate(w, "posts.html", data)
}

// GetMyLikedPosts ...
func GetMyLikedPosts(w http.ResponseWriter, r *http.Request) {
	user := sessions.GetUser(w, r)
	posts := db.GetLikedPostsByUserID(user.UserID)
	data := pageData{"My Posts", user, postData{r.URL.Path, posts}}
	tpl.ExecuteTemplate(w, "posts.html", data)
}
