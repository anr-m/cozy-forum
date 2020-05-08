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
	http.HandleFunc("/posts", controllers.GetPosts)
	http.HandleFunc("/posts/", controllers.GetPostsByCategory)
	http.HandleFunc("/posts/id/", controllers.GetPostByID)
	http.HandleFunc("/createpost", controllers.AuthorizationMW(controllers.CreatePost))
	http.HandleFunc("/comment", controllers.AuthorizationMW(controllers.CreateComment))
	http.HandleFunc("/likepost", controllers.AuthorizationMW(controllers.LikePost))
	http.HandleFunc("/likecomment", controllers.AuthorizationMW(controllers.LikeComment))
	http.HandleFunc("/myposts", controllers.AuthorizationMW(controllers.GetMyPosts))
	http.HandleFunc("/mylikes", controllers.AuthorizationMW(controllers.GetMyLikedPosts))
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
