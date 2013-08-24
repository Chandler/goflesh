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

func (c *Members) ReadMember(where string, args ...interface{}) revel.Result {
	query := `
	    SELECT *
	    FROM "member" m
    ` + where
	name := "members"
	type readObjectType models.Member

	results, err := Dbm.Select(&readObjectType{}, query, args...)
	if err != nil {
		return c.RenderError(err)
	}
	readObjects := make([]*readObjectType, len(results))
	for i, result := range results {
		readObject := result.(*readObjectType)
		if err != nil {
			return c.RenderJson(err)
		}
		readObjects[i] = readObject
	}

	out := make(map[string]interface{})
	out[name] = readObjects

	return c.RenderJson(out)
}
func (c *Members) ReadList(ids []int) revel.Result {
	// if result := c.DevOnly(); result != nil {
	// 	return *result
	// }
	if len(ids) == 0 {
		return c.ReadMember("")
	}
	templateStr := IntArrayToString(ids)
	return c.ReadMember("WHERE m.id = ANY('{" + templateStr + "}')")
}

/////////////////////////////////////////////////////////////////////

func (c *Members) Read(id int) revel.Result {
	return c.ReadMember("WHERE m.id = $1", id)
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
