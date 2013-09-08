package controllers

import (
    "github.com/sfreiberg/gotwilio"
    "github.com/robfig/revel"
)

type Sms struct {
  GorpController
}

func (c *Sms) Index() revel.Result {
  accountSid := ""
  authToken := ""
  twilio := gotwilio.NewTwilioClient(accountSid, authToken)

  from := "+12084024500"
  to := "+12089912446"
  message := "Welcome to gotwilio!"
  twilio.SendSMS(from, to, message, "", "")
  return c.RenderJson("sadf")
}
