package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Organizations struct {
	*revel.Controller
}

func (c Organizations) ReadList() revel.Result {
	return GetList(models.Organization{}, nil)
}

func (c Organizations) Create() revel.Result {
	tableName := "organizations"
	var typedJson map[string][]models.Organization

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	modelObjects := typedJson[tableName]

	err = json.Unmarshal([]byte(data), &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	// Prepare for bulk insert (only way to do it, promise)
	interfaces := make([]interface{}, len(modelObjects))
	for i := range modelObjects {
		interfaces[i] = interface{}(&modelObjects[i])
	}

	// do the bulk insert
	err = dbm.Insert(interfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(interfaces)
}

func (c Organizations) ListGames(organization_id int) revel.Result {
	template := `
		SELECT *
		FROM game
		WHERE organization_id = %d
		`
	query := fmt.Sprintf(template, organization_id)

	result, err := dbm.Select(models.Game{}, query)
	if err != nil {
		return c.RenderError(err)
	}
	out := make(map[string]interface{})
	out["games"] = result

	return c.RenderJson(out)
}
