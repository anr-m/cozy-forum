package main

import (
	"time"

	"./database"
	"./models"
)

func main() {
	db := database.SetUp()
	newUser := models.User{
		Hash:      "123",
		FirstName: "Anuar",
		LastName:  "Mukhanov",
		Email:     "anr",
		Username:  "anr",
	}
	db.CreateUser(&newUser)
	newPost := models.Post{
		UserID:      1,
		Category:    "General",
		Title:       "New post",
		Content:     "New Content",
		Image:       nil,
		TimeCreated: time.Now(),
	}
	db.CreatePost(&newPost)
	newComment := models.Comment{
		PostID:      1,
		UserID:      1,
		Text:        "Newcomment",
		TimeCreated: time.Now(),
	}
	db.CreateComment(&newComment)
	newSession := models.Session{
		SessionID:   "123",
		UserID:      1,
		TimeCreated: time.Now(),
	}
	db.CreateSession(newSession)
	newPost.Like = 3
	db.UpdatePost(newPost)
	newComment.Dislike = 4
	db.UpdateComment(newComment)
}
