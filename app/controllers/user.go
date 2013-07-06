package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"errors"
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

func (c Users) Update(id int) revel.Result {
	var typedJson map[string]models.User
	keyname := "user"
	query := `
		UPDATE "user"
		SET
			email = $1,
			first_name = $2,
			last_name = $3,
			screen_name = $4
		WHERE id = $5
	`

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		return c.RenderError(err)
	}

	model := typedJson[keyname]
	model.Id = id

	result, err := Dbm.Exec(query, model.Email, model.First_name, model.Last_name, model.Screen_name, id)
	if err != nil {
		return c.RenderError(err)
	}
	val, err := result.RowsAffected()
	if err != nil {
		return c.RenderError(err)
	}
	if val != 1 {
		c.Response.Status = 500
		return c.RenderError(errors.New("Did not update exactly one record"))
	}
	return c.RenderJson(val)
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
	query := `
		SELECT *
		FROM "user"
		WHERE
		email = $1
		OR screen_name = $2
		`

	list, err := Dbm.Select(&models.User{}, query, userInfo.Email, userInfo.Screen_name)
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
