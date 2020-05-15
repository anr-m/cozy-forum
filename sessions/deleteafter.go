package sessions

import (
	"net/http"
)

// var timer = time.NewTimer(24*time.Hour)

// DeleteAfter ...
func DeleteAfter(c *http.Cookie) {
	// timer.Stop()
	// timer = time.AfterFunc(time.Duration(c.MaxAge)*time.Second/(24*60*60)*5, func() {
	// 	db.DeleteSession(c.Value)
	// })
}
