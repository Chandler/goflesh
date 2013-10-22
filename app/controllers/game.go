package controllers

import (
	"encoding/json"
	"errors"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
	"time"
)

type Games struct {
	AuthController
}

type GameRead struct {
	models.Game
	Players    string `json:"-"`
	Player_ids []int  `json:"player_ids"`
}

type BasicGameStatsRead struct {
	NumHumans  int   `json:"num_humans"`
	NumZombies int   `json:"num_zombies"`
	TagTimes   []int `json:"tag_times"`
}

type date_wrapper struct {
	Value time.Time
}

func (c *Games) ReadGame(where string, args ...interface{}) revel.Result {
	query := `
	    SELECT *, array(
			SELECT DISTINCT p.id
			FROM player p
			INNER JOIN "user"
				ON g.id = p.game_id
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
		if result := c.DevOnly(); result != nil {
			return *result
		}
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
	if result := c.DevOnly(); result != nil {
		return *result
	}

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
	// TODO: add moderator authentication
	return c.emailList("", game_id)
}

func (c *Games) HumanEmailList(game_id int) revel.Result {
	// TODO: add moderator authentication
	return c.emailList("AND p.last_fed IS NULL", game_id)
}

func (c *Games) ZombieEmailList(game_id int) revel.Result {
	// TODO: add moderator authentication
	return c.emailList("AND p.last_fed IS NOT NULL", game_id)
}

func (c *Games) emailList(where_and string, args ...interface{}) revel.Result {
	query := `
		SELECT u.email email
		FROM "game" g
		INNER JOIN player p
			ON p.game_id = g.id
		INNER JOIN "user" u
			on p.user_id = u.id
		WHERE g.id = $1
	` + where_and

	revel.WARN.Print(query)

	rows, err := Dbm.Db.Query(query, args...)
	if err != nil {
		return c.RenderError(err)
	}

	var emails string

	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails = emails + "," + email
	}

	return c.RenderText(emails)
}

/////////////////////////////////////////////////////////////////////

func (c *Games) ReadBasicStats(game_id int) revel.Result {
	game, err := models.GameFromId(game_id)
	if err != nil {
		return c.RenderError(err)
	}

	player_count := `
		SELECT count(1)
		FROM player
		WHERE player.game_id = $1
	`

	nPlayer, err := Dbm.SelectInt(player_count, game_id)
	if err != nil {
		return c.RenderError(err)
	}

	zombies := `
		SELECT count(1)
		FROM player
		WHERE player.game_id = $1
		AND player.last_fed < $2 
	`

	nZombies, err := Dbm.SelectInt(zombies, game_id, time.Now().Add(-game.TimeToStarve()))
	if err != nil {
		return c.RenderError(err)
	}

	tag_list := `
		SELECT claimed "value"
		FROM tag
		INNER JOIN player
			ON tag.tagger_id = player.id
		WHERE player.game_id = $1
	`
	var wrapped_tag_timestamps []*date_wrapper
	_, err = Dbm.Select(&wrapped_tag_timestamps, tag_list, game_id)
	if err != nil {
		return c.RenderError(err)
	}
	tag_timestamps := make([]string, len(wrapped_tag_timestamps))
	for i, timestamp := range wrapped_tag_timestamps {
		tag_timestamps[i] = timestamp.Value.String()
	}

	out := make(map[string]interface{})
	out["num_humans"] = nPlayer - nZombies
	out["num_zombies"] = nZombies
	out["tag_list"] = tag_timestamps
	return c.RenderJson(out)
}
