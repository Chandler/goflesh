package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"flesh/app/types"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/robfig/revel"
)

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Screen_name string `json:"screen_name"`
	Password    string // TODO: don't send back
	Api_key     string // TODO: don't send back
}

/*
Hash the password that has been set in the User model,
and also generate a random ApiKey
*/
func (u *User) HashPassword() error {
	bcryptCost, existed := revel.Config.Int("user.bcrypt.cost")
	if !existed {
		return &types.ValueNotSetError{"user.bcrypt.cost"}
	}

	var err error
	bytesPassword := []byte(u.Password)
	bytesPassword, err = bcrypt.GenerateFromPassword(bytesPassword, bcryptCost)
	u.Password = string(bytesPassword)
	if err != nil {
		return err
	}

	keyObj, err := uuid.NewV4()
	if err != nil {
		return err
	}
	u.Api_key = keyObj.String()

	return nil
}
