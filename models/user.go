package models

// User type for DB
type User struct {
	UserID    int
	Hash      []byte
	Salt      string
	FirstName string
	LastName  string
	Email     string
	Username  string
}
