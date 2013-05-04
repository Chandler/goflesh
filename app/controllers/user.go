package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Users struct {
	*revel.Controller
}

func (c Users) ReadList() revel.Result {
	return GetList(models.User{})
}

func (c Users) Create(data string) revel.Result {
	// read JSON into models or error out
	var users []models.User
	err := json.Unmarshal([]byte(data), &users)
	if err != nil {
		return c.RenderError(err)
	}

	// Prepare for bulk insert (only way to do it, promise)
	userInterfaces := make([]interface{}, len(users))
	for i, user := range users {
		user.ChangePassword(user.Password)
		userInterfaces[i] = interface{}(&user)
	}
	// do the bulk insert
	err = dbm.Insert(userInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(userInterfaces)
}
