package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"flesh/app/models"
	"flesh/app/types"
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
)

type Users struct {
	GorpController
}

/////////////////////////////////////////////////////////////////////

func (c Users) ReadList() revel.Result {
	return GetList(models.User{}, []string{"Password", "Api_key"})
}

/////////////////////////////////////////////////////////////////////

func (c Users) Read(id int) revel.Result {
	return GetById(models.User{}, []string{"Password", "Api_key"}, id)
}

/////////////////////////////////////////////////////////////////////

func createUsers(data []byte) ([]interface{}, error) {
	const keyName string = "users"
	var typedJson map[string][]models.User

	err := json.Unmarshal(data, &typedJson)
	if err != nil {
		return nil, err
	}

	modelObjects := typedJson[keyName]

	// Prepare for bulk insert (only way to do it, promise)
	interfaces := make([]interface{}, len(modelObjects))
	for i := range modelObjects {
		modelObject := modelObjects[i]
		modelObject.ChangePassword(modelObject.Password)
		interfaces[i] = interface{}(&modelObject)
	}
	return interfaces, nil
}

func (c Users) Create() revel.Result {
	return CreateList(createUsers, c.Request.Body)
}

/////////////////////////////////////////////////////////////////////

// Upload email (or screen name) + password
type UserAuthenticateInput struct {
	Email       string `json:"email"`
	Screen_name string `json:"screen_name"`
	Password    string `json:"password"`
}

type UserAuthenticateOutput struct {
	Id      int    `json:"id"`
	Api_key string `json:"api_key"`
}

func (userInfo *UserAuthenticateInput) Model() (*models.User, error) {
	template := `
    SELECT *
    FROM "user"
    WHERE
    	email = '%s' 
    	OR screen_name = '%s'
    `
	query := fmt.Sprintf(template, userInfo.Email, userInfo.Screen_name)

	list, err := Dbm.Select(&models.User{}, query)
	if llen := len(list); llen != 1 {
		return nil, &types.DatabaseError{fmt.Sprintf("Got %d users instead of 1", llen)}
	}
	user := list[0].(*models.User)
	return user, err
}

/*
Endpoint: given email (or screen_name) + password,
return user_id and api_key
*/
func (c Users) Authenticate() revel.Result {
	var authInfo UserAuthenticateInput
	data, err := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal([]byte(data), &authInfo); err != nil {
		return c.RenderError(err)
	}

	user, err := authInfo.Model()
	if err != nil {
		return c.RenderError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInfo.Password))

	if err != nil {
		c.Response.Status = 401
		return c.RenderText("")
	}

	out := UserAuthenticateOutput{user.Id, user.Api_key}

	return c.RenderJson(out)
}
