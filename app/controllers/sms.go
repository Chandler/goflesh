package controllers

import (
	"flesh/app/models"
	"flesh/app/routes"
	"flesh/app/utils"
	"github.com/robfig/revel"
	"github.com/sfreiberg/gotwilio"
	"github.com/sfreiberg/gotwilio"
	"os"
	"time"
)

type SmsController struct {
	AuthController
}

type TwilioInfo struct {
	Client      *gotwilio.Twilio
	TextingUser *models.User
}

func (c *SmsController) SmsRouter(body string, from_number string, account_sid string) revel.Result {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	responderPhone := os.Getenv("TWILIO_FROM_NUMBER")

	errJson := make(map[string]string)

	if accountSid == "" || authToken == "" || responderPhone == "" {
		if account_sid != accountSid {
			message := "Twilio cannot be used before setting up the environment!"
			revel.ERROR.Print(message)
			c.Response.Status = 500
			errJson["error"] = message
			return c.RenderJson(errJson)
		}
	}

	if account_sid != accountSid {
		revel.ERROR.Print("Got a text from Twilio with incorrect account")
		c.Response.Status = 400
		errJson["error"] = "Not Authorized"
		return c.RenderJson(errJson)
	}

	revel.WARN.Print(accountSid, authToken, responderPhone)
	revel.WARN.Print(body, from_number, account_sid)

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	if account_sid != accountSid {
		c.Response.Status = 400
		errJson["error"] = "Not Authorized"
		return c.RenderJson(errJson)
	}

	phone, err := utils.NormalizePhoneToE164(from_number)
	if err != nil {
		c.Response.Status = 400
		errJson["error"] = "Invalid phone number. Phone number must be passed as a string in E.164 format"
		return c.RenderJson(errJson)
	}

	revel.WARN.Print(phone)

	taggerUser, err := models.UserFromPhone(phone)
	if err != nil {
		c.Response.Status = 422
		errJson["error"] = "No user is registered with this phone number. Please sign up with your phone number online and try again."
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	human, status, err := models.PlayerFromHumanCode(body)
	if err != nil {
		c.Response.Status = status
		errJson["error"] = "Human code invalid"
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	if !human.IsHuman() {
		c.Response.Status = 403
		errJson["error"] = "You cannot tag a non-human"
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	game := human.Game()

	if !game.IsRunning() {
		c.Response.Status = 422
		errJson["error"] = "Tags cannot be registered when the game is closed"
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	tagger, err := models.PlayerFromUserIdGameId(taggerUser.Id, game.Id)
	revel.WARN.Print(tagger, taggerUser.Id, game.Id)
	if err != nil {
		c.Response.Status = 400
		errJson["error"] = "You cannot tag players when you aren't in the same game!"
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	if !tagger.IsZombie() {
		c.Response.Status = 403
		errJson["error"] = "You cannot tag because you are not a zombie. You are " + tagger.Status()
		twilio.SendSMS(responderPhone, phone, errJson["error"], "", "")
		return c.RenderJson(errJson)
	}

	now := time.Now()
	_, status, err = models.NewTag(game, tagger, human, &now)
	if err != nil {
		c.Response.Status = status
		twilio.SendSMS(responderPhone, phone, err.Error(), "", "")
		return c.RenderError(err)
	}

	success := human.User().Screen_name + " successfully tagged!"
	twilio.SendSMS(responderPhone, phone, success, "", "")
	return c.RenderText(success)
}
