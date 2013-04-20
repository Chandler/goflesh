package controllers

import (
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Organizations struct {
	Application
}

func (c Organizations) List() revel.Result {
	query := `
    SELECT *
    FROM organization
    `

	org, _ := dbm.Select(models.Organization{}, query)

	out := make(map[string]interface{})
	out["Organization"] = org
	return c.RenderJson(out)
}
