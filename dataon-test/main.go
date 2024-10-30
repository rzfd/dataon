package main

import (
	"github.com/rzfd/dataon-test/internal/config"
	"github.com/rzfd/dataon-test/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitDb()
	config.SeedDB()
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
