package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"flesh/app/models"
	"flesh/app/types"
	"fmt"
	"github.com/robfig/revel"
)

type Users struct {
	*revel.Controller
}

/////////////////////////////////////////////////////////////////////

func (c Users) ReadList() revel.Result {
	return GetList(models.User{}, []string{"Password", "Api_key"})
}

/////////////////////////////////////////////////////////////////////

func (c Users) Create(data string) revel.Result {
	// read JSON into models or error out
	var dat map[string][]models.User
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		return c.RenderError(err)
	}
	users := dat["users"]

	// Prepare for bulk insert (only way to do it, promise)
	userInterfaces := make([]interface{}, len(users))
	for i := range users {
		user := users[i]
		user.ChangePassword(user.Password)
		userInterfaces[i] = interface{}(&user)
	}
	// do the bulk insert
	if err := dbm.Insert(userInterfaces...); err != nil {
		return c.RenderError(err)
	}

	// Return a copy of the data with id's set
	return c.RenderJson(userInterfaces)
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

	list, err := dbm.Select(&models.User{}, query)
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
func (c Users) Authenticate(data string) revel.Result {
	var authInfo UserAuthenticateInput
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
