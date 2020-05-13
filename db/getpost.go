package db

import (
	"fmt"

	"cozy-forum/models"
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
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.TimeCreated, &post.TimeString)
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
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.TimeCreated, &post.TimeString)
	}

	err = getPostCategories(&post)
	if err != nil {
		return post, err
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
		SELECT posts.*
		FROM posts
		INNER JOIN postcategories
		ON postcategories.postid = posts.postid
		INNER JOIN categories
		ON categories.categoryid = postcategories.categoryid
		WHERE categories.categoryname = ?
	`, category)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.TimeCreated, &post.TimeString)
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
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.TimeCreated, &post.TimeString)
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
		err = getPostCategories(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetLikedPostsByUserID ...
func GetLikedPostsByUserID(userid int) ([]models.Post, error) {

	var posts []models.Post

	row, err := db.Query(`
		SELECT posts.*
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
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.TimeCreated, &post.TimeString)
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
		err = getPostCategories(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
