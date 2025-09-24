package handler

import (
	"bileygr/components"
	"bileygr/config"
	"bileygr/db"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegistation(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, components.Registration())
}

func Register(ctx echo.Context) error {
	return errors.New("error not implemented")
}

func HandleLogin(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, components.Login())
}

func HandleLoginAuth(ctx echo.Context) error {
	cfg, errC := config.Load()
	if errC != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to load config",
		})
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	var hashedPassword, userID string
	err := db.DevDB.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&userID, &hashedPassword)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid credentials",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(cfg.JWT.TokenExpiry).Unix()

	t, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not generate token",
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(cfg.JWT.TokenExpiry)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
	})
}
