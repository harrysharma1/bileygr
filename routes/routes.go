package routes

import (
	"bileygr/handler"
	"bileygr/middleware"

	"github.com/labstack/echo"
)

func Run(app *echo.Echo) {
	app.Static("/static", "static")
	app.GET("/", handler.Home)

	users(app)
	auth(app)
	reading(app)

	app.Logger.Fatal(app.Start(":6969"))
}

func auth(app *echo.Echo) {
	app.GET("/registration", handler.HandleRegistation)
	app.GET("/login", handler.HandleLogin)
	app.POST("/auth/login", handler.HandleLoginAuth)
	app.POST("/auth/register", handler.SaveUser)
}

func users(app *echo.Echo) {
	protected := app.Group("")
	protected.Use(middleware.JWT)
	protected.GET("/users/:id", handler.GetUser)
	protected.PUT("/users/:id", handler.UpdateUser)
	protected.DELETE("/users/:id", handler.DeleteUser)
}

func reading(app *echo.Echo) {
	app.GET("/:readingType/:id", handler.GetReading)
	app.POST("/reading", handler.SaveReading)
	app.PUT("/:readingType/:id", handler.UpdateReading)
	app.DELETE("/:readingType/:id", handler.DeleteReading)
}
