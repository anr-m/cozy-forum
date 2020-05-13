package main

import (
	"log"
	"net/http"
	"os"

	"cozy-forum/db"
	"cozy-forum/errorhandle"
	"cozy-forum/httprouter"
	"cozy-forum/tpl"
)

func init() {
	db.SetUp()
	tpl.SetUp()
}

func main() {
	defer db.Close()
	mux := httprouter.GetServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port", port)
	errorhandle.Fatal(http.ListenAndServe(":"+port, mux))
}
