package models

import (
	"flesh/app/utils"
	"github.com/coopernurse/gorp"
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

type User struct {
	Id          int        `json:"id"`
	Email       string     `json:"email"`
	First_name  string     `json:"first_name"`
	Last_name   string     `json:"last_name"`
	Screen_name string     `json:"screen_name"`
	Password    string     `json:"password,omitempty"`
	Api_key     string     `json:"api_key,omitempty"`
	Last_login  *time.Time `json:"last_login"`
	Created     *time.Time `json:"created"`
	Updated     *time.Time `json:"updated"`
}

func (model *User) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *User) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}

type UserGetAuthenticate struct {
	Id      int    `json:"id"`
	Api_key string `json:"api_key,omitempty"`
}

/*
Hash the password that has been set in the User model,
and also generate a random ApiKey
*/
func (u *User) ChangePassword(plaintext string) error {

	keyObj, err := uuid.NewV4()
	if err != nil {
		return err
	}
	u.Api_key = keyObj.String()

	u.Password, err = utils.HashPassword(plaintext)
	if err != nil {
		return err
	}

	return nil
}
