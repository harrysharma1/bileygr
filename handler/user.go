package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

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
