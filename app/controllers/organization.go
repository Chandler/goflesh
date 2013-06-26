package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"fmt"
	"github.com/robfig/revel"
)

type Organizations struct {
	*revel.Controller
}

/////////////////////////////////////////////////////////////////////

func (c Organizations) ReadList() revel.Result {
	return GetList(models.Organization{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c Organizations) Read(id int) revel.Result {
	return GetById(models.Organization{}, nil, id)
}

/////////////////////////////////////////////////////////////////////

func createOrganizations(data []byte) ([]interface{}, error) {
	const keyName string = "organizations"
	var typedJson map[string][]models.Organization

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

func (c Organizations) Create() revel.Result {
	return CreateList(createOrganizations, c.Request.Body)
}

/////////////////////////////////////////////////////////////////////

func (c Organizations) ListGames(organization_id int) revel.Result {
	template := `
		SELECT *
		FROM game
		WHERE organization_id = %d
		`
	query := fmt.Sprintf(template, organization_id)

	result, err := Dbm.Select(models.Game{}, query)
	if err != nil {
		return c.RenderError(err)
	}
	out := make(map[string]interface{})
	out["games"] = result

	return c.RenderJson(out)
}
