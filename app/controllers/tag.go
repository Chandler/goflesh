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
	errJson := make(map[string]string)
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

	human, status, err := models.PlayerFromHumanCode(code)
	if err != nil {
		c.Response.Status = status
		c.RenderError(err)
	}

	game := human.Game()

	if !game.IsRunning() {
		c.Response.Status = 422
		errJson["error"] = "Tags cannot be registered when the game is closed"
		return c.RenderJson(errJson)
	}

	now := time.Now()
	_, status, err = models.NewTag(game, tagger, human, &now)
	if err != nil {
		c.Response.Status = status
		return c.RenderError(err)
	}
	return c.Redirect(routes.Players.Read(human.Id))
}

type jsn map[string]string

func (j *jsn) SetError(message string) {
	(*j)["error"] = message
}

func (j *jsn) SetSuccess() {
	(*j)["success"] = "true"
}

func TagByPhone(body string, taggerUser *models.User) (string, int, map[string]string) {
	msgJson := make(jsn)

	human, status, err := models.PlayerFromHumanCode(body)
	if err != nil {
		errorMessage := "Human code invalid"
		msgJson.SetError(errorMessage)
		return errorMessage, 422, msgJson
	}

	if !human.IsHuman() {
		errorMessage := "You cannot tag a non-human"
		msgJson.SetError(errorMessage)
		return errorMessage, 403, msgJson
	}

	game := human.Game()

	if !game.IsRunning() {
		errorMessage := "Tags cannot be registered when the game is closed"
		msgJson.SetError(errorMessage)
		return errorMessage, 422, msgJson
	}

	tagger, err := models.PlayerFromUserIdGameId(taggerUser.Id, game.Id)
	revel.WARN.Print(tagger, taggerUser.Id, game.Id)
	if err != nil {
		errorMessage := "You cannot tag players when you aren't in the same game!"
		msgJson.SetError(errorMessage)
		return errorMessage, 400, msgJson
	}

	if !tagger.IsZombie() {
		errorMessage := "You cannot tag because you are not a zombie. You are " + tagger.Status()
		msgJson.SetError(errorMessage)
		return errorMessage, 403, msgJson
	}

	now := time.Now()
	_, status, err = models.NewTag(game, tagger, human, &now)
	if err != nil {
		msgJson.SetError(err.Error())
		return err.Error(), status, msgJson
	}

	success := human.User().Screen_name + " successfully tagged!"
	msgJson.SetSuccess()
	return success, 200, msgJson
}
