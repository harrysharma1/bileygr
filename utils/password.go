package utils

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckHashedPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	return err == nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("error password shorter than 8 charaacters")
	}

	return nil
}
