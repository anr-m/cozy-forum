package models

// PageData ...
type PageData struct {
	PageTitle  string
	Categories []string
	User       User
	Data       interface{}
}
