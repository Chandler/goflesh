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
	var dat map[string][]models.Game
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		return c.RenderError(err)
	}
	games := dat["games"]

	// Prepare for bulk insert (only way to do it, promise)
	gameInterfaces := make([]interface{}, len(games))
	for i := range games {
		gameInterfaces[i] = interface{}(&games[i])
	}
	// do the bulk insert
	err = Dbm.Insert(gameInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(gameInterfaces)
}
