package handler

import (
	"bileygr/components"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func HandleRegistation(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, components.Registration())
}

func HandleLogin(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, components.Login())
}

func HandleLoginAuth(ctx echo.Context) error {
	// _username := ctx.FormValue("username")
	// password := ctx.FormValue("password")
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, map[string]string{
	// 		"error": err.Error(),
	// 	})
	// }

	return errors.New("error not implemented")
}
