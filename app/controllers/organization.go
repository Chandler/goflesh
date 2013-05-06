package controllers

import (
	"encoding/json"
	"flesh/app/models"
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
	var orgs []models.Organization
	err := json.Unmarshal([]byte(data), &orgs)
	if err != nil {
		return c.RenderError(err)
	}

	// Prepare for bulk insert (only way to do it, promise)
	orgInterfaces := make([]interface{}, len(orgs))
	for i, org := range orgs {
		orgInterfaces[i] = interface{}(&org)
	}
	// do the bulk insert
	err = dbm.Insert(orgInterfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(orgInterfaces)
}
