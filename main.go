package main

import (
	"log"
	"net/http"

	"./controllers"
	"./database"
	"./models"
)

var db database.DB

func init() {
	database.SetUp()
	// newPost := models.Post{
	// 	UserID:      1,
	// 	Category:    "General",
	// 	Title:       "New post",
	// 	Content:     "New Content",
	// 	Image:       nil,
	// 	TimeCreated: time.Now(),
	// }
	// db.CreatePost(&newPost)
	// newComment := models.Comment{
	// 	PostID:      1,
	// 	UserID:      1,
	// 	Text:        "Newcomment",
	// 	TimeCreated: time.Now(),
	// }
	// db.CreateComment(&newComment)
	// newSession := models.Session{
	// 	SessionID:   "123",
	// 	UserID:      1,
	// 	TimeCreated: time.Now(),
	// }
	// db.CreateSession(newSession)
	// newPost.Like = 3
	// db.UpdatePost(newPost)
	// newComment.Dislike = 4
	// db.UpdateComment(newComment)
	// fmt.Println(db.EmailExists("anr"))
	// fmt.Println(db.EmailExists("aanr"))
	// fmt.Println(db.UsernameExists("anr"))
	// fmt.Println(db.UsernameExists("aanr"))
}

func main() {
	newUser := models.User{
		Hash:      "123",
		FirstName: "Yelnur",
		LastName:  "Nurzhanuly",
		Email:     "yelnur@gmail.com",
		Username:  "Yelnur",
	}
	database.DataBase.CreateUser(&newUser)
	// http.HandleFunc("/", index)
	http.HandleFunc("/register", controllers.Register)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		w.Header().Set("Content-Type", "text/html")
// 		w.Write([]byte(`<form method="POST"><input type="text" name="username"><input type="email" name="email"><button type="submit">Submit</button></form>`))
// 	} else if r.Method == http.MethodPost {
// 		r.ParseForm()
// 		if db.EmailExists(r.FormValue("email")) {
// 			w.Write([]byte("Email exists"))
// 		} else if db.UsernameExists(r.FormValue("username")) {
// 			w.Write([]byte("Username exists"))
// 		} else {
// 			w.Write([]byte(r.FormValue("email")))
// 		}
// 	}
// }
