package main

import (
	"log"
	"net/http"

	"./controllers"
	"./database"
)

func init() {
	database.SetUp()
}

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/createpost", controllers.CreatePost)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
