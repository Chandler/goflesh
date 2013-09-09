package controllers

import (
	"encoding/base64"
	"errors"
	"flesh/app/models"
	"github.com/robfig/revel"
	"strings"
)

type AuthController struct {
	GorpController
	User  *models.User // the logged-in user
	Cache map[string]map[int]interface{}
}

func (c *AuthController) Auth() error {
	if c.User != nil {
		return nil // already authorized
	}

	encodedAuth := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Basic ")
	decodedAuth, err := base64.StdEncoding.DecodeString(encodedAuth)
	if err != nil {
		return err
	}
	components := strings.SplitN(string(decodedAuth), ":", 2)
	if len(components) != 2 {
		return errors.New("Did not get username and API key")
	}
	userId := components[0]
	apiKey := components[1]

	// get the user by id
	obj, err := Dbm.Get(models.User{}, userId)
	if err != nil {
		return err
	}
	user := obj.(*models.User)

	if user.Api_key != apiKey {
		return errors.New("Invalid API key")
	}

	c.User = user

	return nil
}

func (c *AuthController) DevOnly() *revel.Result {
	if !c.isDevMode() {
		response := c.NotFound("")
		return &response
	}
	return nil
}

func (c *AuthController) PermissionDenied() revel.Result {
	c.Response.Status = 403
	return c.RenderError(errors.New("Permission denied"))
}

func (c *AuthController) SentAuth() bool {
	return c.Request.Header.Get("Authorization") != ""
}
