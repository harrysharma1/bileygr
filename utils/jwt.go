package utils

import (
	"bileygr/db"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"username":        username,
			"expiration_time": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("error invalid token")
	}

	return nil
}

func GetUserFromToken(token *jwt.Token) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	var username string
	err := db.DevDB.QueryRow("SELECT username FROM users WHERE id = $1", userId).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
