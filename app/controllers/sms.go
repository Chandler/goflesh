package controllers

import (
    "github.com/sfreiberg/gotwilio"
    "github.com/robfig/revel"
)

type Sms struct {
  GorpController
}

func (c *Sms) Index() revel.Result {
  accountSid := "ACabd6be0f388b473592246ed204b78586"
  authToken := "408f0ca9fc3ea08d63c6fd4b155db629"
  twilio := gotwilio.NewTwilioClient(accountSid, authToken)

  from := "+12084024500"
  to := "+12089912446"
  message := "Welcome to gotwilio!"
  twilio.SendSMS(from, to, message, "", "")
  return c.RenderJson("sadf")
}


