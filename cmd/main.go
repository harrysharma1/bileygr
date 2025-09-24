package main

import (
	"bileygr/db"
	"bileygr/routes"
	"log"

	"github.com/labstack/echo"
)

func main() {
	_, err := db.NewDatabase("postgresql://harrysharma@localhost/bileygr")
	if err != nil {
		log.Fatalf("error connecting to database: %e", err)
	}

	routes.Run(echo.New())
}
