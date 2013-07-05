package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type Organizations struct {
	GorpController
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
	query := `
		SELECT *
		FROM game
		WHERE organization_id = ?
		`
	result, err := Dbm.Select(models.Game{}, query, organization_id)
	if err != nil {
		return c.RenderError(err)
	}
	out := make(map[string]interface{})
	out["games"] = result

	return c.RenderJson(out)
}

/////////////////////////////////////////////////////////////////////

type OrganizationInformation struct {
	models.Organization
	NumMembers int    `json:"numMembers"`
	Image      string `json:"image"`
}

// Endpoint for discovery page organizations list
func (c Organizations) DiscoveryInformationList() revel.Result {
	query := `
		SELECT *
		FROM organization
		`

	result, err := Dbm.Select(OrganizationInformation{}, query)
	if err != nil {
		return c.RenderError(err)
	}
	out := make(map[string]interface{})
	out["organizations"] = result
	return c.RenderJson(out)
}
