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
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/register", middleware.PageDataMW(false, controllers.Register))
	mux.HandleFunc("/login", middleware.PageDataMW(false, controllers.Login))
	mux.HandleFunc("/logout", controllers.Logout)
	mux.HandleFunc("/posts", middleware.PageDataMW(false, controllers.GetPosts))
	mux.HandleFunc("/posts/", middleware.PageDataMW(false, controllers.GetPostsByCategory))
	mux.HandleFunc("/posts/id/", middleware.PageDataMW(false, controllers.GetPostByID))
	mux.HandleFunc("/createpost", middleware.PageDataMW(true, controllers.CreatePost))
	mux.HandleFunc("/comment", middleware.PageDataMW(true, controllers.CreateComment))
	mux.HandleFunc("/likepost", middleware.PageDataMW(true, controllers.LikePost))
	mux.HandleFunc("/likecomment", middleware.PageDataMW(true, controllers.LikeComment))
	mux.HandleFunc("/myposts", middleware.PageDataMW(true, controllers.GetMyPosts))
	mux.HandleFunc("/mylikes", middleware.PageDataMW(true, controllers.GetMyLikedPosts))
	return mux
}
