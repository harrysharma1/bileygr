package main

import (
	"bileygr/routes"

	"github.com/labstack/echo"
)

func main() {
	routes.Run(echo.New())
}
