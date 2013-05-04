package utils

import (
	"code.google.com/p/go.crypto/bcrypt"
	"flesh/app/types"
	"github.com/robfig/revel"
)

/*
Return a bcrypt hash for a pre-salted password string
*/
func HashPassword(plaintext string) (string, error) {
	bcryptCost, ok := revel.Config.Int("user.bcrypt.cost")
	if !ok {
		return plaintext, &types.ValueNotSetError{"user.bcrypt.cost"}
	}

	bytesPassword := []byte(plaintext)
	bytesPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcryptCost)
	return string(bytesPassword), err
}
