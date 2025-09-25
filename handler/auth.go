package handler

import (
	"bileygr/components"
	"bileygr/config"
	"bileygr/db"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegistation(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, components.Registration())
}

func HandleRegistationAuth(ctx echo.Context) error {
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
	cookie.Name = "authToken"
	cookie.Value = t
	cookie.Expires = time.Now().Add(cfg.JWT.TokenExpiry)
	cookie.Path = "/"
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)

	return ctx.Redirect(http.StatusSeeOther, "/")
}

func HandleLogoutAuth(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "authToken"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour) // Set expiry in the past
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = -1 // Immediately expire the cookie

	ctx.SetCookie(cookie)
	return ctx.Redirect(http.StatusSeeOther, "/")
}
