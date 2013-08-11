package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Players struct {
	AuthController
}

/////////////////////////////////////////////////////////////////////

func (c *Players) ReadList() revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	return GetList(models.Player{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c *Players) Read(id int) revel.Result {
	if result := c.DevOnly(); result != nil {
		return *result
	}
	return GetById(models.Player{}, nil, id)
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
			Dbm.Insert(member)
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
			revel.WARN.Print(err, human_code)
			return c.RenderError(err)
		}
	}
	return c.RenderJson(interfaces)
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
