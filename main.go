package main

import (
	"github.com/TakeshiHA/test-middleware/controllers"
	"github.com/TakeshiHA/test-middleware/database"
	"github.com/TakeshiHA/test-middleware/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	// Init Databases
	db := database.InitDatabases()
	defer db.DisconnectDatabases()

	// Echo instance
	e := echo.New()

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	docs.SwaggerInfo.Title = "Swagger Compartamos Test"

	// Init Controllers
	controllers.InitController(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
