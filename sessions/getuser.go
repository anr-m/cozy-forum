package sessions

import (
	"net/http"

	"cozy-forum/db"
	"cozy-forum/models"
)

// GetUser gets user from the database from the session
func GetUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User

	c, err := r.Cookie("session")
	if err != nil {
		return user, nil
	}

	session, err := db.GetSession(c.Value)
	if session.UserID == 0 || err != nil {
		return user, err
	}

	user, err = db.GetUserByID(session.UserID)
	if user.UserID == 0 || err != nil {
		return user, err
	}

	// session.TimeCreated = time.Now()
	// db.UpdateSession(session)

	c.Path = "/"
	c.MaxAge = 24 * 60 * 60
	http.SetCookie(w, c)

	DeleteAfter(c)

	return user, nil
}
