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
	Phone       string     `json:"phone"`
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

func NewUser(
	email string,
	first_name string,
	last_name string,
	screen_name string,
	phone string,
	password string,
) (user *User, status_code int, err error) {
	// Naive password checks
	if len(password) < 8 {
		return nil, 422, errors.New("Password must be at least 8 characters")
	}
	if strings.Contains(email, password) {
		return nil, 422, errors.New("Password cannot be part of email")
	}
	if strings.Contains(password, first_name) ||
		strings.Contains(password, last_name) ||
		strings.Contains(password, screen_name) {
		return nil, 422, errors.New("Password cannot contain your name or screen name")
	}
	now := time.Now()
	user = &User{0, email, first_name, last_name, screen_name, "208 555-5555", "", "", &now, TimeTrackedModel{}}
	user.ChangePassword(password)
	err = Dbm.Insert(user)
	if err != nil {
		// insert failed. Perform some diagnostic queries to find out why
		count, diagnostic_err := Dbm.SelectInt(`SELECT count(*) FROM "user" WHERE email = $1` email)
		if diagnostic_err != nil {
			return nil, 500, diagnostic_err
		}
		if count > 1 {
			return nil, 409, errors.New("An account with this email already exists")
		}

		count, diagnostic_err = Dbm.SelectInt(`SELECT count(*) FROM "user" WHERE screen_name = $1`, email)
		if diagnostic_err != nil {
			return nil, 500, diagnostic_err
		}
		if count > 1 {
			return nil, 409, errors.New("An account with this screen name already exists")
		}
		// couldn't determine the source of the error. let's just return it
		return nil, 500, err
	}
	return user, 209, nil
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
