package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Players struct {
	*revel.Controller
}

func (c Players) ReadList() revel.Result {
	return GetList(models.Player{}, nil)
}

func (c Players) Create(data string) revel.Result {
	// read JSON into models or error out
	var dat map[string][]models.Player
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		return c.RenderError(err)
	}
	players := dat["players"]

	// Prepare for bulk insert (only way to do it, promise)
	playerInterfaces := make([]interface{}, len(players))
	for i := range players {
		playerInterfaces[i] = interface{}(&players[i])
	}
	// do the bulk insert
	err = dbm.Insert(playerInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(playerInterfaces)
}
