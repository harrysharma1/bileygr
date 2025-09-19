package main

import (
	"bileygr/handler"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	app.GET("/", handler.HomeHandler)

	app.Static("/static", "static")
	app.Logger.Fatal(app.Start(":6969"))
}
