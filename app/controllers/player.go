package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Players struct {
	*revel.Controller
}

/////////////////////////////////////////////////////////////////////

func (c Players) ReadList() revel.Result {
	return GetList(models.Player{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c Players) Read(id int) revel.Result {
	return GetById(models.Player{}, nil, id)
}

/////////////////////////////////////////////////////////////////////

func createPlayers(data []byte) ([]interface{}, error) {
	const keyName string = "players"
	var typedJson map[string][]models.Player

	err := json.Unmarshal(data, &typedJson)
	if err != nil {
		return nil, err
	}

	modelObjects := typedJson[keyName]

	// Prepare for bulk insert (only way to do it, promise)
	interfaces := make([]interface{}, len(modelObjects))
	for i := range modelObjects {
		interfaces[i] = interface{}(&modelObjects[i])
	}
	return interfaces, nil
}

func (c Players) Create() revel.Result {
	return CreateList(createPlayers, c.Request.Body)
}
