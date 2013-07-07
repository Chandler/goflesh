package controllers

import (
	"encoding/json"
	"flesh/app/models"
	"github.com/robfig/revel"
	"database/sql"
)

type Players struct {
	GorpController
}

/////////////////////////////////////////////////////////////////////

func (c Players) ReadList() revel.Result {
	return GetList(models.Player{}, nil)
}

/////////////////////////////////////////////////////////////////////

func (c Players) Read(id int) revel.Result {
	return GetById(models.Player{}, nil, id)
}

/////////////////////////////////////////////////////////////////////

func createPlayers(data []byte) ([]interface{}, error) {
	const keyName string = "players"
	var typedJson map[string][]models.Player

	err := json.Unmarshal(data, &typedJson)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		// if this user is not a member of an org, add them
		if result.RowsAffected() == 0 {
			// game := Mike's function 
			models.Member{0, user_id, game.organization_id}
		}
	}
	return interfaces, nil
}

func (c Players) Create() revel.Result {
	return CreateList(createPlayers, c.Request.Body)
}
func MemberExists(user_id int, game_id int) (models.Member{}, error) {
	query := `
		SELECT *
		FROM member JOIN organization ON member.organization_id = organization.id 
		INNER JOIN game ON game.organization_id = organization.id
		WHERE user_id = $1 AND game.id = $2
		`
	result, err := Dbm.Select(models.Member{}, query, user_id, organization_id)

	return result, err
} 