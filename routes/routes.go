package routes

import (
	"bileygr/handler"

	"github.com/labstack/echo"
)

func Run(app *echo.Echo) {
	app.Static("/static", "static")
	app.GET("/", handler.Home)
	users(app)
	app.Logger.Fatal(app.Start(":6969"))
}

func users(app *echo.Echo) {
	app.GET("/users/:id", handler.GetUser)
	app.POST("/users", handler.SaveUser)
	app.PUT("/users/:id", handler.UpdateUser)
	app.DELETE("/users/:id", handler.DeleteUser)
}
