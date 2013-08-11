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

type GameRead struct {
	models.Game
	Players    string `json:"-"`
	Player_ids []int  `json:"player_ids"`
}

func (c *Games) ReadGame(where string, args ...interface{}) revel.Result {
	query := `
	    SELECT *, array(
			SELECT DISTINCT p.id
			FROM player p
			INNER JOIN "user"
				ON g.id = p.user_id
			) players
	    FROM "game" g
    ` + where
	name := "games"
	type readObjectType GameRead

	results, err := Dbm.Select(&readObjectType{}, query, args...)
	if err != nil {
		return c.RenderError(err)
	}
	readObjects := make([]*readObjectType, len(results))
	for i, result := range results {
		readObject := result.(*readObjectType)
		readObject.Player_ids, err = PostgresArrayStringToIntArray(readObject.Players)
		if err != nil {
			return c.RenderJson(err)
		}
		readObjects[i] = readObject
	}

	out := make(map[string]interface{})
	out[name] = readObjects

	return c.RenderJson(out)
}

/////////////////////////////////////////////////////////////////////

func (c *Games) ReadList(ids []int) revel.Result {
	if len(ids) == 0 {
		return c.ReadGame("")
	}
	templateStr := IntArrayToString(ids)
	return c.ReadGame("WHERE g.id = ANY('{" + templateStr + "}')")
}

/////////////////////////////////////////////////////////////////////

func (c *Games) Read(id int) revel.Result {
	return c.ReadGame("WHERE g.id = $1", id)
}

/////////////////////////////////////////////////////////////////////

func (c *Games) Update(id int) revel.Result {
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

/////////////////////////////////////////////////////////////////////

func (c *Games) AllEmailList(game_id int) revel.Result {
	return c.emailList("", "", game_id)
}

func (c *Games) HumanEmailList(game_id int) revel.Result {
	return c.emailList(
		"LEFT JOIN tag ON u.id = taggee_id LEFT JOIN oz on u.id = oz.id",
		"AND taggee_id IS NULL AND (oz.id IS NULL OR oz.confirmed = FALSE)",
		game_id,
	)
}

func (c *Games) ZombieEmailList(game_id int) revel.Result {
	return c.emailList(
		"LEFT JOIN tag ON u.id = taggee_id LEFT JOIN oz on u.id = oz.id",
		"AND (taggee_id IS NOT NULL OR (oz.id IS NOT NULL AND oz.confirmed = TRUE))",
		game_id,
	)
}

func (c *Games) emailList(join string, where_and string, args ...interface{}) revel.Result {
	query := `
		SELECT u.email email
		FROM "game" g
		INNER JOIN player p
			ON p.game_id = g.id
		INNER JOIN "user" u
			on p.user_id = u.id
	` + join + `
		WHERE g.id = $1
	` + where_and

	revel.WARN.Print(query)

	rows, err := Dbm.Db.Query(query, args...)
	if err != nil {
		return c.RenderError(err)
	}

	var emails []string

	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails = append(emails, email)
	}
	jsn := make(map[string][]string)
	jsn["emails"] = emails
	return c.RenderJson(jsn)
}
