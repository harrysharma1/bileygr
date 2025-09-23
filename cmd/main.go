package main

import (
	"bileygr/db"
	"bileygr/routes"

	"github.com/labstack/echo"
)

func main() {
	db.InitDevDB()
	routes.Run(echo.New())
}
