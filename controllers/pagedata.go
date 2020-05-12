package controllers

import "cozy-forum/models"

type pageData struct {
	PageTitle string
	User      models.User
	Data      interface{}
}
