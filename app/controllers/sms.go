package controllers

import (
	"flesh/app/models"
	"flesh/app/utils"
	"github.com/robfig/revel"
	"github.com/sfreiberg/gotwilio"
	"os"
	"strings"
)

type SmsController struct {
	AuthController
}

func (c *SmsController) SmsRouter(Body string, From string, AccountSid string) revel.Result {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	responderPhone := os.Getenv("TWILIO_FROM_NUMBER")

	errJson := make(jsn)

	if accountSid == "" || authToken == "" || responderPhone == "" {
		if AccountSid != accountSid {
			message := "Twilio cannot be used before setting up the environment!"
			revel.ERROR.Print(message)
			errJson.SetError(message)
			c.Response.Status = 500
			return c.RenderJson(errJson)
		}
	}

	if AccountSid != accountSid {
		revel.ERROR.Print("Got a text from Twilio with incorrect account")
		c.Response.Status = 400
		errJson.SetError("Not Authorized")
		return c.RenderJson(errJson)
	}

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	if AccountSid != accountSid {
		message := "Twilio not authorized"
		revel.ERROR.Print(message)
		errJson.SetError(message)
		c.Response.Status = 400
		return c.RenderJson(errJson)
	}

	phone, err := utils.NormalizePhoneToE164(From)
	if err != nil {
		message := "Invalid phone number. Phone number must be passed as a string in E.164 format"
		revel.ERROR.Print(message)
		errJson.SetError(message)
		c.Response.Status = 400
		return c.RenderJson(errJson)
	}

	textingUser, err := models.UserFromPhone(phone)
	if err != nil {
		message := "No user is registered with this phone number. Please sign up with your phone number online and try again."
		errJson.SetError(message)
		twilio.SendSMS(responderPhone, phone, message, "", "")
		c.Response.Status = 422
		return c.RenderJson(errJson)
	}

	// Do the actual dispatching to different handlers
	lowerTrimmedBody := strings.ToLower(strings.TrimSpace(Body))
	var (
		smsText    string
		statusCode int
		jsnObject  map[string]string
	)
	switch {
	case strings.Contains(lowerTrimmedBody, "stats"):
		// call stats
	default: // attempting a tag via text
		smsText, statusCode, jsnObject = TagByPhone(lowerTrimmedBody, textingUser)
	}

	twilio.SendSMS(responderPhone, phone, smsText, "", "")
	c.Response.Status = statusCode
	return c.RenderJson(jsnObject)
}
