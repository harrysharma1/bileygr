package handler

import (
	"bileygr/db"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(ctx echo.Context) error {
	username := ctx.Param("id")
	if username == "" {
		return ctx.JSON(http.StatusBadGateway, map[string]string{
			"error": "username is required",
		})
	}

	var userID string
	err := db.DevDB.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "user not found",
		})
	}

	tokenUserID := ctx.Get("user_id").(string)
	if userID != tokenUserID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "unauthorized access",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"user_id":  userID,
		"username": username,
	})
}

func UpdateUser(ctx echo.Context) error {
	return errors.New("error not implemented update users function")
}

func DeleteUser(ctx echo.Context) error {
	return errors.New("error not implemented delete users function")
}
