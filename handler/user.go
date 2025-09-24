package handler

import (
	"bileygr/db"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	id := uuid.New()
	_, err = db.DevDB.Exec("INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, NOW())",
		id.String(), username, string(hashedPassword))
	if err != nil {
		log.Printf("Database error: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
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
