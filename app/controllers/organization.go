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

func (c Organizations) ReadList() revel.Result {
	return GetList(models.Organization{}, nil)
}

func (c Organizations) Create(data string) revel.Result {
	// read JSON into models or error out

	var dat map[string][]models.Organization
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		return c.RenderError(err)
	}
	orgs := dat["organizations"]

	// Prepare for bulk insert (only way to do it, promise)
	orgInterfaces := make([]interface{}, len(orgs))
	for i := range orgs {
		orgInterfaces[i] = interface{}(&orgs[i])
	}

	// do the bulk insert
	err = dbm.Insert(orgInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(orgInterfaces)
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
