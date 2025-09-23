package handler

import (
	"bileygr/db"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type Creds struct {
	Username string `json:"username", db:"username"`
	Password string `json:"password", db:"password"`
}

func SaveUser(ctx echo.Context) error {
	creds := &Creds{}
	err := json.NewDecoder(ctx.Request().Body).Decode(creds)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request format",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "processing password",
		})
	}
	id := uuid.New()
	log.Printf("Username: %s\n Password: %s", creds.Username, string(hashedPassword))

	_, err = db.DevDB.Exec("INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, NOW())",
		id.String(), creds.Username, string(hashedPassword))
	if err != nil {
		log.Printf("Database error: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "creating user",
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "user created successfully",
	})
}

func GetUser(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func UpdateUser(ctx echo.Context) error {
	return errors.New("error not implemented update users function")
}

func DeleteUser(ctx echo.Context) error {
	return errors.New("error not implemented delete users function")
}
