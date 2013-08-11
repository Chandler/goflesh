package models

import (
	"flesh/app/routes"
	"flesh/app/utils"
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
	TimeTrackedModel
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

/////////////////////////////////////////////////////////////////////

type PasswordReset struct {
	Id      int        `json:"id"`
	Expires *time.Time `json:"expires"`
	Code    string     `json:"code"`
	TimeTrackedModel
}

func (m *PasswordReset) GenerateCode() error {
	keyObj, err := uuid.NewV4()
	if err != nil {
		return err
	}
	m.Code = keyObj.String()
	week, err := time.ParseDuration("168h")
	if err != nil {
		return err
	}
	expires := time.Now().Add(week).UTC()
	m.Expires = &expires
	return nil
}

func (m *PasswordReset) ResetLink() string {
	return routes.Users.PasswordReset(m.Code)
}
