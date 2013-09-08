package controllers

import (
	"flesh/app/models"
	"flesh/app/routes"
	"github.com/robfig/revel"
	"time"
)

type Tags struct {
	AuthController
}

func (c *Tags) Tag(player_id int, code string) revel.Result {
	query := `
        SELECT *
        FROM human_code
        WHERE code = $1
    `
	err := c.Auth()
	if err != nil {
		return c.RenderError(err)
	}

	if !c.SentAuth() {
		c.Response.Status = 401
		errJson["error"] = "No authentication information was sent"
		return c.RenderJson(errJson)
	}

	if c.User == nil {
		c.Response.Status = 403
		errJson["error"] = "User credentials bad"
		return c.RenderJson(errJson)
	}

	tagger, err := models.PlayerFromId(player_id)
	revel.WARN.Print(player_id)
	if err != nil {
		return c.RenderError(err)
	}

	if !tagger.IsZombie() {
		c.Response.Status = 403
		errJson["error"] = "You cannot tag because you are not a zombie"
		return c.RenderJson(errJson)
	}

	if tagger.User_id != c.User.Id {
		c.Response.Status = 403
		errJson["error"] = "Tags cannot be registered for other users"
		return c.RenderJson(errJson)
	}

	var list []*models.HumanCode
	_, err = Dbm.Select(&list, query, code)
	if err != nil {
		return c.RenderError(err)
	}
	if len(list) != 1 {
		c.Response.Status = 403
		errJson["error"] = "Invalid human code"
		return c.RenderJson(errJson)
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
	return c.Redirect(routes.Players.Read(human.Id))
}
