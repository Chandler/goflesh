package controllers

import (
	"encoding/json"
	"errors"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
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

func (c Games) Update(id int) revel.Result {
	var typedJson map[string]models.Game
	keyname := "game"
	query := `
		UPDATE "game"
		SET
			name = $1,
			timezone = $2,
			registration_start_time = $3,
			registration_end_time = $4,
			running_start_time = $5,
			running_end_time = $6
		WHERE id = $7
	`

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	model := typedJson[keyname]
	model.Id = id

	result, err := Dbm.Exec(
		query,
		model.Name,
		model.Timezone,
		model.Registration_start_time,
		model.Registration_end_time,
		model.Running_start_time,
		model.Running_end_time,
		id,
	)
	if err != nil {
		return c.RenderError(err)
	}
	val, err := result.RowsAffected()
	if err != nil {
		return c.RenderError(err)
	}
	if val != 1 {
		c.Response.Status = 500
		return c.RenderError(errors.New("Did not update exactly one record"))
	}
	return c.RenderJson(val)
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
