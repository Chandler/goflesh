package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Members struct {
	AuthController
}

/////////////////////////////////////////////////////////////////////

func (c *Members) ReadList() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	return GetList(models.Member{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c *Members) Read(id int) revel.Result {
	return GetById(models.Member{}, nil, id)
}

/////////////////////////////////////////////////////////////////////

func createMembers(data []byte) ([]interface{}, error) {
	const keyName string = "members"
	var typedJson map[string][]models.Member

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

func (c *Members) Create() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	return CreateList(createMembers, c.Request.Body)
}
