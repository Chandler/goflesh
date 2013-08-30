package controllers

import (
	"encoding/json"
	"errors"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
	"time"
)

type Players struct {
	AuthController
}

type PlayerRead struct {
	models.Player
	StatusString string `json:"status"`
	HumanCode    string `json:"human_code,omitempty"`
}

func (c *Players) ReadPlayer(where string, args ...interface{}) revel.Result {
	query := `
	    SELECT p.*
	    FROM player p
    ` + where
	type readObjectType PlayerRead

	c.Auth()

	results, err := Dbm.Select(&readObjectType{}, query, args...)
	if err != nil {
		return c.RenderError(err)
	}
	user_ids := make([]int, len(results))
	players := make([]*readObjectType, len(results))
	for i, result := range results {
		readObject := result.(*readObjectType)
		readObject.StatusString = readObject.Status()
		if c.User != nil && c.User.Id == readObject.Player.User_id {
			human_code := readObject.Player.HumanCode()
			readObject.HumanCode = human_code.Code
		}
		user_ids[i] = readObject.Player.User_id
		if err != nil {
			return c.RenderJson(err)
		}
		players[i] = readObject
	}

	templateStr := IntArrayToString(user_ids)
	users, err := FetchUsers(c.User, "WHERE u.id = ANY('{"+templateStr+"}')")
	if err != nil {
		return c.RenderJson(err)
	}

	out := make(map[string]interface{})
	out["players"] = players
	out["users"] = users

	return c.RenderJson(out)
}

func (c *Players) ReadList(game_id int, ids []int) revel.Result {
	if game_id != 0 {
		return c.ReadPlayer("INNER JOIN game g ON p.game_id = g.id WHERE g.id = $1", game_id)
	}
	if len(ids) == 0 {
		return c.ReadPlayer("")
	}
	templateStr := IntArrayToString(ids)
	return c.ReadPlayer("WHERE p.id = ANY('{" + templateStr + "}')")
}

/////////////////////////////////////////////////////////////////////

func (c *Players) Read(id int) revel.Result {
	return c.ReadPlayer("WHERE p.id = $1", id)
}

/////////////////////////////////////////////////////////////////////

func (c *Players) Create() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	const keyName string = "players"
	var typedJson map[string][]models.Player

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	modelObjects := typedJson[keyName]

	// Prepare for bulk insert (only way to do it, promise)
	interfaces := make([]interface{}, len(modelObjects))
	for i, player := range modelObjects {
		interfaces[i] = interface{}(&player)

		user_id := player.User_id
		game_id := player.Game_id

		result, err := MemberExists(user_id, game_id)
		if err != nil {
			return c.RenderError(err)
		}
		// if this user is not a member of an org, add them
		if result == nil {
			game, err := models.GameFromId(game_id)
			if err != nil {
				return c.RenderError(err)
			}
			member := models.Member{0, user_id, game.Organization_id, models.TimeTrackedModel{}}
			err = Dbm.Insert(&member)
			if err != nil {
				return c.RenderError(err)
			}
		}
	}

	// do the bulk insert
	err = Dbm.Insert(interfaces...)
	if err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	for _, playerInterface := range interfaces {
		player := playerInterface.(*models.Player)
		// add a human code for the player
		human_code := models.HumanCode{player.Id, "", models.TimeTrackedModel{}}
		human_code.GenerateCode()
		err = Dbm.Insert(&human_code)
		if err != nil {
			return c.RenderError(err)
		}
	}

	out := make(map[string]interface{})
	out[keyName] = interfaces
	return c.RenderJson(out)
}
func MemberExists(user_id int, game_id int) (*models.Member, error) {
	query := `
		SELECT member.Id, member.User_id, member.Organization_id, member.Created, member.Updated 
		FROM member JOIN organization ON member.organization_id = organization.id 
		INNER JOIN game ON game.organization_id = organization.id
		WHERE user_id = $1 AND game.id = $2
		`
	member := models.Member{}
	_, err := Dbm.Select(member, query, user_id, game_id)

	return &member, err
}

func (c *Players) Tag(player_id int, code string) revel.Result {
	query := `
		SELECT *
		FROM human_code
		WHERE code = $1
	`
	err := c.Auth()
	if err != nil {
		return c.RenderError(err)
	}

	if c.SentAuth() {
		c.Response.Status = 401
		return c.RenderText("")
	}

	if c.User == nil {
		c.Response.Status = 403
		return c.RenderText("User credentials bad")
	}

	tagger, err := models.PlayerFromId(player_id)
	if err != nil {
		return c.RenderError(err)
	}

	if !tagger.IsZombie() {
		c.Response.Status = 403
		return c.RenderText("You cannot tag because you are not a zombie")
	}

	if tagger.User_id != c.User.Id {
		c.Response.Status = 403
		return c.RenderText("Tags cannot be registered for other users")
	}

	var list []*models.HumanCode
	_, err = Dbm.Select(&list, query, code)
	if err != nil {
		return c.RenderError(err)
	}
	if len(list) != 1 {
		return c.RenderError(errors.New("Invalid human code"))
	}
	human_code := list[0]
	player, err := Dbm.Get(models.Player{}, human_code.Id)
	if err != nil {
		return c.RenderError(err)
	}

	human := player.(*models.Player)

	gameObj, err := Dbm.Get(models.Game{}, human.Game_id)
	if err != nil {
		return c.RenderError(err)
	}
	game := gameObj.(*models.Game)

	now := time.Now()
	_, err = models.NewTag(game, tagger, human, &now)
	if err != nil {
		return c.RenderError(err)
	}

	return c.Read(human.Id)
}
