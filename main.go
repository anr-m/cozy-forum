package main

import (
	"log"
	"net/http"

	"./controllers"
	"./db"
	"./errorhandle"
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
	http.HandleFunc("/index", controllers.AllowedMethodsMW([]string{"GET"}, controllers.Index))
	http.HandleFunc("/register", controllers.AllowedMethodsMW([]string{"GET", "POST"}, controllers.Register))
	http.HandleFunc("/login", controllers.AllowedMethodsMW([]string{"GET", "POST"}, controllers.Login))
	http.HandleFunc("/logout", controllers.AllowedMethodsMW([]string{"GET"}, controllers.Logout))
	http.HandleFunc("/posts", controllers.AllowedMethodsMW([]string{"GET"}, controllers.AuthorizationMW(false, controllers.GetPosts)))
	http.HandleFunc("/posts/", controllers.AllowedMethodsMW([]string{"GET"}, controllers.AuthorizationMW(false, controllers.GetPostsByCategory)))
	http.HandleFunc("/posts/id/", controllers.AllowedMethodsMW([]string{"GET"}, controllers.AuthorizationMW(false, controllers.GetPostByID)))
	http.HandleFunc("/createpost", controllers.AllowedMethodsMW([]string{"GET", "POST"}, controllers.AuthorizationMW(true, controllers.CreatePost)))
	http.HandleFunc("/comment", controllers.AllowedMethodsMW([]string{"POST"}, controllers.AuthorizationMW(true, controllers.CreateComment)))
	http.HandleFunc("/likepost", controllers.AllowedMethodsMW([]string{"POST"}, controllers.AuthorizationMW(true, controllers.LikePost)))
	http.HandleFunc("/likecomment", controllers.AllowedMethodsMW([]string{"POST"}, controllers.AuthorizationMW(true, controllers.LikeComment)))
	http.HandleFunc("/myposts", controllers.AllowedMethodsMW([]string{"GET"}, controllers.AuthorizationMW(true, controllers.GetMyPosts)))
	http.HandleFunc("/mylikes", controllers.AllowedMethodsMW([]string{"GET"}, controllers.AuthorizationMW(true, controllers.GetMyLikedPosts)))
	log.Println("Server started on port 8080")
	errorhandle.Fatal(http.ListenAndServe(":8080", nil))
}
