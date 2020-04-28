package main

import (
	"log"
	"net/http"

	"./controllers"
	"./db"
)

func init() {
	db.SetUp()
}

func main() {
	defer db.Close()
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/createpost", controllers.CreatePost)
	http.HandleFunc("/posts", controllers.GetPosts)
	http.HandleFunc("posts/", controllers.GetPostsByCategory)
	http.HandleFunc("/comment/", controllers.CreateComment)
	http.HandleFunc("/posts/id/", controllers.GetPostByID)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
