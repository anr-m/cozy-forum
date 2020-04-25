package models

import "time"

// Session type for DB
type Session struct {
	SessionID   string
	UserID      int
	TimeCreated time.Time
}
