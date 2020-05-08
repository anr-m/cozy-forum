package db

import (
	"../errorhandle"
	"../models"
)

// GetPosts ...
func GetPosts() []models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
	`)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}

// GetPostByID ...
func GetPostByID(postid int) models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE postid = ?
	`, postid)
	defer row.Close()

	errorhandle.Check(err)

	var post models.Post
	for row.Next() {
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
	}

	getPostLikesAndDislikes(&post)

	return post
}

// GetPostsByCategory ...
func GetPostsByCategory(category string) []models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE category = ?
	`, category)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}

// GetPostsByUserID ...
func GetPostsByUserID(userid int) []models.Post {
	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE userid = ?
	`, userid)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}

// GetLikedPostsByUserID ...
func GetLikedPostsByUserID(userid int) []models.Post {
	row, err := db.Query(`
		SELECT posts.postid, posts.userid, posts.category, posts.title, posts.content, posts.image, posts.timecreated
		FROM posts
		INNER JOIN postlikes
		ON posts.postid = postlikes.postid
		WHERE postlikes.userid = ? AND postlikes.liked = 1
	`, userid)
	defer row.Close()

	errorhandle.Check(err)

	var posts []models.Post

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated)
		getPostLikesAndDislikes(&post)
		posts = append(posts, post)
	}

	return posts
}
