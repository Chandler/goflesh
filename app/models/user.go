package models

import (
	"flesh/app/utils"
	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Screen_name string `json:"screen_name"`
	Password    string // TODO: don't send back
	Salt        string // TODO: don't send back
	Api_key     string // TODO: don't send back
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

	saltObj, err := uuid.NewV4()
	if err != nil {
		return err
	}
	u.Salt = saltObj.String()

	u.Password, err = utils.HashPassword(plaintext + u.Salt)
	if err != nil {
		return err
	}

	return nil
}
