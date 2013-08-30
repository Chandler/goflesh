package controllers

import (
	"flesh/app/models"
	"encoding/json"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Oz struct {
	GorpController
}

/////////////////////////////////////////////////////////////////////

func (c *Oz) SelectOzs(game_id int) revel.Result {
	query := `
	INSERT INTO oz  
	SELECT oz_pool.id, false, now(), now()
	FROM oz_pool
	INNER JOIN player
		ON oz_pool.id = player.id
	LEFT OUTER JOIN oz
		ON oz_pool.id = oz.id
	WHERE oz.id IS NULL
	AND player.game_id = $1
	ORDER BY random()
	LIMIT $2;
    `

	var typedJson map[string]int

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	num_ozs := typedJson["num_ozs"]

	response, err := Dbm.Exec(query, game_id, num_ozs)
	if err != nil {
		return c.RenderError(err)
	}

	rowsAffected, err := response.RowsAffected()

	return c.RenderJson(rowsAffected)
}

func (c *Oz) CreateTestOz(player_id int) revel.Result {
	oz := &models.Oz{player_id, false, models.TimeTrackedModel{}}

	oz.Confirm()
	return c.RenderJson(oz)
}

