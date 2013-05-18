package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Games struct {
	*revel.Controller
}

func (c Games) ReadList() revel.Result {
	return GetList(models.Game{}, nil)
}

func (c Games) Create(data string) revel.Result {
	// read JSON into models or error out
	var games []models.Game
	err := json.Unmarshal([]byte(data), &games)
	if err != nil {
		return c.RenderError(err)
	}

	// Prepare for bulk insert (only way to do it, promise)
	gameInterfaces := make([]interface{}, len(games))
	for i, org := range games {
		gameInterfaces[i] = interface{}(&org)
	}
	// do the bulk insert
	err = dbm.Insert(gameInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(gameInterfaces)
}
