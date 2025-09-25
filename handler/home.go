package handler

import (
	"bileygr/components"
	"bileygr/config"
	"bileygr/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func Home(ctx echo.Context) error {
	var authenticated bool = false
	var username string = ""
	cookie, err := ctx.Cookie("authToken")
	if err == nil && cookie.Value != "" {
		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			cfg, err := config.Load()
			if err != nil {
				return nil, err
			}
			return []byte(cfg.JWT.Secret), nil
		})
		if err == nil && token.Valid {
			authenticated = true
			username, err = utils.GetUserFromToken(token)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{
					"error": "failed to get user information",
				})
			}
		}
	}

	return Render(ctx, http.StatusOK, components.Home(authenticated, username))
}
