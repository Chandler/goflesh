package controllers

import (
	"flesh/app/models"
	"flesh/app/routes"
	"flesh/app/utils"
	"github.com/robfig/revel"
	"time"
	"os"
	"github.com/sfreiberg/gotwilio"
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

func (c *Tags) TagByPhone(Body string, From string, AccountSid string) revel.Result {

  accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
  authToken  := os.Getenv("TWILIO_AUTH_TOKEN")
  from_phone := os.Getenv("TWILIO_FROM_NUMBER")
  twilio     := gotwilio.NewTwilioClient(accountSid, authToken)

	errJson := make(map[string]string)

  if AccountSid != accountSid {
		c.Response.Status = 400
		errJson["error"] = "Not Authorized"
		return c.RenderJson(errJson)
	}



	phone, err := utils.NormalizePhoneToE164(From)
	if err != nil {
		c.Response.Status = 400
		errJson["error"] = "Invalid phone number. Phone number must be passed as a string in E.164 format"
		return c.RenderJson(errJson)
	}
  
  revel.WARN.Print(phone)
	
	taggerUser, err := models.UserFromPhone(phone)
	if err != nil {
		c.Response.Status = 422
		errJson["error"] = "No user is registered with this phone number"
    twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	human, status, err := models.PlayerFromHumanCode(Body)
	if err != nil {
		c.Response.Status = status
		errJson["error"] = "Human code invalid"
		twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	if !human.IsHuman() {
		c.Response.Status = 403
		errJson["error"] = "You cannot tag a non-human"
		twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	game := human.Game()

	if !game.IsRunning() {
		c.Response.Status = 422
		errJson["error"] = "Tags cannot be registered when the game is closed"
		twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	tagger, err := models.PlayerFromUserIdGameId(taggerUser.Id, game.Id)
	revel.WARN.Print(tagger,taggerUser.Id, game.Id)
	if err != nil {
		c.Response.Status = 400
		errJson["error"] = "You cannot tag players when you aren't in the same game!"
		twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	if !tagger.IsZombie() {
		c.Response.Status = 403
		errJson["error"] = "You cannot tag because you are not a zombie. You are " + tagger.Status()
		twilio.SendSMS(from_phone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	now := time.Now()
	_, status, err = models.NewTag(game, tagger, human, &now)
	if err != nil {
		c.Response.Status = status
  	twilio.SendSMS(from_phone, phone, err.Error(), "", "")
		return c.RenderError(err)
	}

	success := human.User().Screen_name + " successfully tagged!"
  twilio.SendSMS(from_phone, phone, success, "", "")
	return c.RenderText(success)
}
