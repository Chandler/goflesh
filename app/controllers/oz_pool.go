package controllers

import (
	"encoding/json"
	"errors"
	"flesh/app/models"
	"github.com/robfig/revel"
)

type OzPools struct {
	AuthController
}

/////////////////////////////////////////////////////////////////////

func (c *OzPools) ReadList() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	name := "oz_pools"
	query := `
    SELECT *
    FROM "oz_pool"
    `
	result, err := Dbm.Select(models.OzPool{}, query)
	if err != nil {
		return c.RenderError(err)
	}
	out := make(map[string]interface{})
	out[name] = result

	return c.RenderJson(out)
}

/////////////////////////////////////////////////////////////////////

func (c *OzPools) Read(id int) revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	c.Auth()
	if c.User == nil || c.User.Id != id {
		return c.PermissionDenied()
	}
	name := "oz_pools"
	result, err := Dbm.Get(models.OzPool{}, id)
	if err != nil {
		return c.RenderError(err)
	}
	if result == nil {
		return FResponse404(FResponse404{interface{}(map[string]string{"error": "not found"})})
	}

	out := make(map[string][]interface{})
	out[name] = []interface{}{result}

	return c.RenderJson(out)
}

/////////////////////////////////////////////////////////////////////

func (c *OzPools) Delete(id int) revel.Result {
	c.Auth()
	if c.User == nil || c.User.Id != id {
		return c.PermissionDenied()
	}
	query := `
        DELETE
        FROM "oz_pool"
        WHERE id = $1
    `

	result, err := Dbm.Exec(query, id)
	if err != nil {
		return c.RenderError(err)
	}
	val, err := result.RowsAffected()
	if err != nil {
		return c.RenderError(err)
	}
	if val != 1 {
		c.Response.Status = 500
		return c.RenderError(errors.New("Did not delete exactly one record"))
	}
	return c.RenderJson(val)
}

/////////////////////////////////////////////////////////////////////

func (c *OzPools) Create() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	return CreateList(createOzPools, c.Request.Body)
}

/////////////////////////////////////////////////////////////////////

func createOzPools(data []byte) ([]interface{}, error) {
	const keyName string = "oz_pools"
	var typedJson map[string][]models.OzPool

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
