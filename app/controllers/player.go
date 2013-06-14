package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Players struct {
	*revel.Controller
}

func (c Players) ReadList() revel.Result {
	return GetList(models.Player{}, nil)
}

func (c Players) Create() revel.Result {
	tableName := "players"
	var typedJson map[string][]models.Player

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	modelObjects := typedJson[tableName]

	// Prepare for bulk insert (only way to do it, promise)
	interfaces := make([]interface{}, len(modelObjects))
	for i := range modelObjects {
		interfaces[i] = interface{}(&modelObjects[i])
	}

	// do the bulk insert
	err = Dbm.Insert(interfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(interfaces)
}
