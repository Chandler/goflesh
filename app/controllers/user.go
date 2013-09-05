package controllers

import (
	"bytes"
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"errors"
	"flesh/app/models"
	"flesh/app/types"
	"flesh/app/utils"
	"fmt"
	"github.com/robfig/revel"
	"html/template"
	"io/ioutil"
)

type Users struct {
	AuthController
}

var reset_password_email_template *template.Template

/////////////////////////////////////////////////////////////////////

func init() {
	var err error
	reset_password_email_template, err = template.ParseFiles("app/views/Users/SendPasswordReset.html")
	if err != nil {
		panic(err)
	}
}

/////////////////////////////////////////////////////////////////////

func FetchUsers(current_user *models.User, where string, args ...interface{}) ([]*models.UserRead, error) {
	query := `
	    SELECT *, array(
			SELECT DISTINCT p.id
			FROM player p
			INNER JOIN "user"
				ON u.id = p.user_id
			) players, array(
			SELECT DISTINCT m.organization_id
			FROM member m
			INNER JOIN "user"
				ON u.id = m.user_id
			) organizations
	    FROM "user" u
    ` + where

	results, err := Dbm.Select(&models.UserRead{}, query, args...)
	if err != nil {
		return nil, err
	}
	readObjects := make([]*models.UserRead, len(results))
	for i, result := range results {
		readObject := result.(*models.UserRead)
		readObject.Player_ids, err = PostgresArrayStringToIntArray(readObject.Players)
		readObject.User.CleanSensitiveFields(current_user == nil || current_user.Id != readObject.Id)
		readObject.AddAvatars()
		if current_user == nil || current_user.Id != readObject.Id {
		}
		if err != nil {
			return nil, err
		}
		readObject.Organization_ids, err = PostgresArrayStringToIntArray(readObject.Organizations)
		if err != nil {
			return nil, err
		}
		readObjects[i] = readObject
	}
	return readObjects, nil
}

func (c *Users) ReadUser(where string, args ...interface{}) revel.Result {
	c.Auth()
	readObjects, err := FetchUsers(c.User, where, args...)
	if err != nil {
		return c.RenderError(err)
	}

	out := make(map[string]interface{})
	out["users"] = readObjects

	return c.RenderJson(out)
}

func (c *Users) ReadList(ids []int) revel.Result {
	if len(ids) == 0 {
		return c.ReadUser("")
	}
	templateStr := IntArrayToString(ids)
	return c.ReadUser("WHERE u.id = ANY('{" + templateStr + "}')")
}

/////////////////////////////////////////////////////////////////////

func (c *Users) Read(id int) revel.Result {
	return c.ReadUser("WHERE u.id = $1", id)
}

/////////////////////////////////////////////////////////////////////

func (c *Users) Update(id int) revel.Result {
	c.Auth()
	if c.User == nil || c.User.Id != id {
		return c.PermissionDenied()
	}

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

func (c *Users) Create() revel.Result {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderError(err)
	}

	const keyName string = "users"
	var typedJson map[string][]*models.User

	err = json.Unmarshal(data, &typedJson)
	if err != nil {
		c.Response.Status = 400
		return c.RenderError(err)
	}

	users := typedJson[keyName]

	for i, user := range users {
		insertedUser, statusCode, err := models.NewUser(
			user.Email,
			user.First_name,
			user.Last_name,
			user.Screen_name,
			user.Password,
		)
		if err != nil {
			c.Response.Status = statusCode
			return c.RenderError(err)
		}
		users[i] = insertedUser
	}

	return c.RenderJson(users)
}

/////////////////////////////////////////////////////////////////////

// Upload email (or screen name) + password
type UserAuthenticateInput struct {
	Email       string `json:"email"`
	Screen_name string `json:"screen_name"`
	Password    string `json:"password"`
	Api_key     string `json:"api_key"` // TODO: fix client-side auth so we don't have this hack
}

type UserAuthenticateOutput struct {
	Id      int    `json:"id"`
	Api_key string `json:"api_key"`
}

func (userInfo *UserAuthenticateInput) Model() (*models.User, error) {
	query := `
		SELECT *
		FROM "user"
		WHERE email = $1
		OR screen_name = $2
		OR api_key = $3` // TODO: fix client-side auth so we don't have this hack

	list, err := Dbm.Select(&models.User{}, query, userInfo.Email, userInfo.Screen_name,
		userInfo.Api_key) // TODO: fix client-side auth so we don't have this hack
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
func (c *Users) Authenticate() revel.Result {
	var authInfo UserAuthenticateInput
	data, err := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal([]byte(data), &authInfo); err != nil {
		return c.RenderError(err)
	}

	user, err := authInfo.Model()
	if err != nil {
		return c.RenderError(err)
	}

	out := UserAuthenticateOutput{user.Id, user.Api_key}

	if authInfo.Api_key == user.Api_key { // TODO: fix client-side auth so we don't have this hack
		return c.RenderJson(out)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInfo.Password))

	if err != nil {
		c.Response.Status = 401
		return c.RenderText("")
	}

	return c.RenderJson(out)
}

/////////////////////////////////////////////////////////////////////

func (c *Users) SendPasswordReset() revel.Result {
	var authInfo UserAuthenticateInput
	data, err := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal([]byte(data), &authInfo); err != nil {
		return c.RenderError(err)
	}

	user, err := authInfo.Model()
	if err != nil {
		return c.RenderError(err)
	}

	// only have one password reset link active at a time
	_, err = Dbm.Exec("DELETE FROM password_reset WHERE id = $1", user.Id)
	if err != nil {
		return c.RenderError(err)
	}
	reset := models.PasswordReset{user.Id, nil, "", models.TimeTrackedModel{}}
	err = reset.GenerateCode()
	if err != nil {
		return c.RenderError(err)
	}
	err = Dbm.Insert(&reset)
	if err != nil {
		return c.RenderError(err)
	}

	b := new(bytes.Buffer)
	reset_password_email_template.Execute(b, reset)

	utils.SendEmail(
		"Flesh Password Reset",
		b.String(),
		"Flesh Server",
		"flesh@example.com",
		user.First_name+" "+user.Last_name,
		user.Email,
	)

	return c.RenderText("")
}

func (c *Users) PasswordReset(code string) revel.Result {
	query := `
		SELECT id
		FROM password_reset
		WHERE code = $1
		AND expires > now()
	`
	user_id, err := Dbm.SelectInt(query, code)
	if err != nil {
		return c.RenderError(err)
	}
	userInterface, err := Dbm.Get(models.User{}, user_id)
	if err != nil {
		return c.RenderError(err)
	}
	user := userInterface.(*models.User)

	// return as if we authenticated
	out := UserAuthenticateOutput{user.Id, user.Api_key}

	return c.RenderJson(out)
}
