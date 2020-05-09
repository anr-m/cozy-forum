package sessions

import "net/http"

// Logout to log the user out
func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		return
	}

	c.MaxAge = -1
	http.SetCookie(w, c)
}
