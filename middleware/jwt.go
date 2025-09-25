package middleware

import (
	"bileygr/config"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("authToken")
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing authentication",
			})
		}
		cfg, err := config.Load()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "loading config",
			})
		}
		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}
		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("username", claims["username"])
		return next(ctx)
	}
}
