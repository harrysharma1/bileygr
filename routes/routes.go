package routes

import (
	"bileygr/handler"

	"github.com/labstack/echo"
)

func Run(app *echo.Echo) {
	app.Static("/static", "static")
	app.GET("/", handler.Home)

	users(app)
	reading(app)

	app.Logger.Fatal(app.Start(":6969"))
}

func users(app *echo.Echo) {
	app.GET("/users/:id", handler.GetUser)
	app.POST("/users", handler.SaveUser)
	app.PUT("/users/:id", handler.UpdateUser)
	app.DELETE("/users/:id", handler.DeleteUser)
}

func reading(app *echo.Echo) {
	app.GET("/:readingType/:id", handler.GetReading)
	app.POST("/reading", handler.SaveReading)
	app.PUT("/:readingType/:id", handler.UpdateReading)
	app.DELETE("/:readingType/:id", handler.DeleteReading)
}
