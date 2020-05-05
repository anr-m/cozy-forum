package main

import (
	"log"
	"net/http"

	"./controllers"
	"./db"
	"./tpl"
)

func init() {
	db.SetUp()
	tpl.SetUp()
}

func main() {
	defer db.Close()
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	http.HandleFunc("/", controllers.NotFoundHandler)
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/createpost", controllers.CreatePost)
	http.HandleFunc("/posts", controllers.GetPosts)
	http.HandleFunc("/posts/", controllers.GetPostsByCategory)
	http.HandleFunc("/comment", controllers.CreateComment)
	http.HandleFunc("/likepost", controllers.LikePost)
	http.HandleFunc("/likecomment", controllers.LikeComment)
	http.HandleFunc("/posts/id/", controllers.GetPostByID)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
