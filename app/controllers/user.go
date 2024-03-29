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
	"strings"
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
		revel.ERROR.Print(err)
		return nil, err
	}
	users := make([]*models.UserRead, len(results))
	for i, result := range results {
		user := result.(*models.UserRead)
		user.Player_ids, err = PostgresArrayStringToIntArray(user.Players)
		user.AddAvatars()
		user.User.CleanSensitiveFields(current_user == nil || current_user.Id != user.Id)
		if current_user == nil || current_user.Id != user.Id {
		}
		if err != nil {
			revel.ERROR.Print(err)
			return nil, err
		}
		user.Organization_ids, err = PostgresArrayStringToIntArray(user.Organizations)
		if err != nil {
			revel.ERROR.Print(err)
			return nil, err
		}
		users[i] = user
	}
	// // TODO: think this through better
	// if len(results) > 0 { // only bother if other results were returned
	// 	for i := 0; i < len(user_ids); i++ {
	// 		if user_ids[i] == models.OZ_USER_ID {
	// 			players[len(results)] = GetOzUserRead()
	// 		}
	// 	}
	// }
	return users, nil
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
	if id == models.OZ_USER_ID {
		out := make(map[string]interface{})
		out["users"] = []*models.UserRead{GetOzUserRead()}
		return c.RenderJson(out)
	}
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
			screen_name = $4,
			phone = $5,
			updated = now()
		WHERE id = $6
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

	// changing passwords has to be handled specially
	changingPassword := true
	if model.Password == "" {
		// TODO: redesign so I don't need to validate an arbitrary password
		model.Password = "^?8`8468`86`L^866229~"
		changingPassword = false
	}

	statusCode, err := model.ValidateAndNormalizeUserFields()
	if err != nil {
		c.Response.Status = statusCode
		return c.RenderError(err)
	}

	result, err := Dbm.Exec(query, model.Email, model.First_name, model.Last_name, model.Screen_name, model.Phone, id)
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
	if changingPassword {
		err = model.ChangePassword(model.Password)
		if err != nil {
			return c.RenderError(err)
		}
		query = `
		UPDATE "user"
		SET
			password = $1,
			api_key = $2,
			updated = now()
		WHERE id = $3
		`
		result, err = Dbm.Exec(query, model.Password, model.Api_key, id)
		if err != nil {
			return c.RenderError(err)
		}
		val, err = result.RowsAffected()
		if val != 1 {
			return c.RenderError(errors.New("Did not update exactly one record when changing password"))
		}
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
			*user.Phone,
			user.Password,
		)
		if err != nil {
			c.Response.Status = statusCode
			return c.RenderError(err)
		}
		insertedUser.CleanSensitiveFields(false)
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
		SELECT DISTINCT *
		FROM "user"
		WHERE (email = $1 AND email != '')
		OR (screen_name = $2 AND screen_name != '')
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

	authInfo.Email = strings.ToLower(authInfo.Email)

	user, err := authInfo.Model()
	if err != nil {
		c.Response.Status = 401
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

func GetOzUserRead() *models.UserRead {
	return &models.UserRead{models.User{models.OZ_USER_ID, "", "Original", "Zombie", "original zombie", nil, "", "", nil, models.TimeTrackedModel{}}, map[string]string{"hash": "fe4568abcf47619251dce5119dd820f4"}, "", []int{models.OZ_PLAYER_ID}, "", []int{}}
}
