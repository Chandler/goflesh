package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Games struct {
	GorpController
}

/////////////////////////////////////////////////////////////////////

func (c Games) ReadList() revel.Result {
	return GetList(models.Game{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c Games) Read(id int) revel.Result {
	return GetById(models.Game{}, nil, id)
}

/////////////////////////////////////////////////////////////////////

func createModels(data []byte) ([]interface{}, error) {
	const keyName string = "games"
	var typedJson map[string][]models.Game

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

func (c Games) Create() revel.Result {
	return CreateList(createModels, c.Request.Body)
}
