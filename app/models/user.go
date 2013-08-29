package models

import (
	"crypto/md5"
	"errors"
	"flesh/app/routes"
	"flesh/app/utils"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"strings"
	"time"
)

type User struct {
	Id          int        `json:"id"`
	Email       string     `json:"email,omitempty"`
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

type UserRead struct {
	User
	Avatar           map[string]string `json:"avatar"`
	Players          string            `json:"-"`
	Player_ids       []int             `json:"player_ids"`
	Organizations    string            `json:"-"`
	Organization_ids []int             `json:"organization_ids"`
}

func UserFromId(id int) (*User, error) {
	user, err := Dbm.Get(User{}, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User could not be found")
	}
	return user.(*User), err
}

func (u *User) CleanSensitiveFields(clearEmail bool) {
	// omit passsword and api key
	u.Password = ""
	u.Api_key = ""
	// blank out email
	if clearEmail {
		u.Email = ""
	}
}

func (u *User) UserRead() *UserRead {
	userRead := new(UserRead)
	userRead.User = *u
	userRead.AddAvatars()
	return userRead
}

func (ur *UserRead) AddAvatars() {
	// make a Gravatar-compatible email hash
	emailHash := md5.New()
	emailHash.Write([]byte(strings.ToLower(strings.TrimSpace(ur.Email))))
	ur.Avatar = make(map[string]string)
	ur.Avatar["hash"] = fmt.Sprintf("%x", emailHash.Sum(nil))
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
