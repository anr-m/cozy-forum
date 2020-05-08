package controllers

import "../models"

type pageData struct {
	PageTitle string
	User      models.User
	Data      interface{}
}
