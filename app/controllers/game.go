package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Games struct {
	*revel.Controller
}

func (c Games) ReadList() revel.Result {
	return GetList(models.Game{}, nil)
}

func (c Games) Create() revel.Result {
	tableName := "games"
	var typedJson map[string][]models.Game

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
