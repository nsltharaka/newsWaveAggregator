package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(oldPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(oldPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
