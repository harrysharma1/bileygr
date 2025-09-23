package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func SaveReading(ctx echo.Context) error {
	return errors.New("error not implemented save users function")
}

func GetReading(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func UpdateReading(ctx echo.Context) error {
	return errors.New("error not implemented update users function")
}

func DeleteReading(ctx echo.Context) error {
	return errors.New("error not implemented delete users function")
}
