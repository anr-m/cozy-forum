package db

import (
	"fmt"

	"../models"
)

// GetPosts ...
func GetPosts(userid int) ([]models.Post, error) {

	var posts []models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
	`)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated, &post.TimeString)
		err = getPostLikesAndDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostByID ...
func GetPostByID(postid int, userid int) (models.Post, error) {

	var post models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE postid = ?
	`, postid)
	defer row.Close()

	if err != nil {
		return post, err
	}

	for row.Next() {
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated, &post.TimeString)
	}

	err = getPostLikesAndDislikes(&post)
	if err != nil {
		return post, err
	}
	if userid != 0 {
		err = postLikedByUser(&post, userid)
		if err != nil {
			return post, err
		}
		err = postDislikedByUser(&post, userid)
		if err != nil {
			return post, err
		}
	}

	return post, nil
}

// GetPostsByCategory ...
func GetPostsByCategory(category string, userid int) ([]models.Post, error) {

	var posts []models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE category = ?
	`, category)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated, &post.TimeString)
		err = getPostLikesAndDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostsByUserID ...
func GetPostsByUserID(userid int) ([]models.Post, error) {

	var posts []models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE userid = ?
	`, userid)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated, &post.TimeString)
		err = getPostLikesAndDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetLikedPostsByUserID ...
func GetLikedPostsByUserID(userid int) ([]models.Post, error) {

	var posts []models.Post

	row, err := db.Query(`
		SELECT posts.postid, posts.userid, posts.username, posts.category, posts.title, posts.content, posts.image, posts.timecreated, posts.timestring
		FROM posts
		INNER JOIN postlikes
		ON posts.postid = postlikes.postid
		WHERE postlikes.userid = ? AND postlikes.liked = 1
	`, userid)
	defer row.Close()

	if err != nil {
		fmt.Println("error!")
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Category, &post.Title, &post.Content, &post.HTMLImage, &post.TimeCreated, &post.TimeString)
		err = getPostLikesAndDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		posts = append(posts, post)
	}

	return posts, nil
}
