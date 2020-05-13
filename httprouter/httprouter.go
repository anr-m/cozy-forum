package httprouter

import (
	"cozy-forum/controllers"
	"cozy-forum/middleware"
	"net/http"
)

// GetServeMux returns servemux for forum
func GetServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	mux.HandleFunc("/", middleware.AllowedMethodsMW([]string{"GET"}, controllers.Index))
	mux.HandleFunc("/register", middleware.AllowedMethodsMW([]string{"GET", "POST"}, middleware.PageDataMW(false, controllers.Register)))
	mux.HandleFunc("/login", middleware.AllowedMethodsMW([]string{"GET", "POST"}, middleware.PageDataMW(false, controllers.Login)))
	mux.HandleFunc("/logout", middleware.AllowedMethodsMW([]string{"GET"}, controllers.Logout))
	mux.HandleFunc("/posts", middleware.AllowedMethodsMW([]string{"GET"}, middleware.PageDataMW(false, controllers.GetPosts)))
	mux.HandleFunc("/posts/", middleware.AllowedMethodsMW([]string{"GET"}, middleware.PageDataMW(false, controllers.GetPostsByCategory)))
	mux.HandleFunc("/posts/id/", middleware.AllowedMethodsMW([]string{"GET"}, middleware.PageDataMW(false, controllers.GetPostByID)))
	mux.HandleFunc("/createpost", middleware.AllowedMethodsMW([]string{"GET", "POST"}, middleware.PageDataMW(true, controllers.CreatePost)))
	mux.HandleFunc("/comment", middleware.AllowedMethodsMW([]string{"POST"}, middleware.PageDataMW(true, controllers.CreateComment)))
	mux.HandleFunc("/likepost", middleware.AllowedMethodsMW([]string{"POST"}, middleware.PageDataMW(true, controllers.LikePost)))
	mux.HandleFunc("/likecomment", middleware.AllowedMethodsMW([]string{"POST"}, middleware.PageDataMW(true, controllers.LikeComment)))
	mux.HandleFunc("/myposts", middleware.AllowedMethodsMW([]string{"GET"}, middleware.PageDataMW(true, controllers.GetMyPosts)))
	mux.HandleFunc("/mylikes", middleware.AllowedMethodsMW([]string{"GET"}, middleware.PageDataMW(true, controllers.GetMyLikedPosts)))
	return mux
}
