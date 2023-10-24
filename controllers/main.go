package controllers

import (
	"time"

	"github.com/TakeshiHA/test-middleware/repository"
	"github.com/TakeshiHA/test-middleware/usecase"
	"github.com/labstack/echo/v4"
)

func InitController(e *echo.Echo) {
	timeoutContext := time.Duration(60) * time.Second
	api := e.Group("/api")

	// Groups
	cityGroup := api.Group("/cities")
	clientGroup := api.Group("/clients")

	// Repository
	cityRepo := repository.NewCityRepository()
	clientRepo := repository.NewClientRepository()

	// Usecase
	cityUCase := usecase.NewCityUsecase(cityRepo, timeoutContext)
	clientUCase := usecase.NewClientUsecase(clientRepo, timeoutContext)

	// Controller
	NewCityController(cityGroup, cityUCase)
	NewClientController(clientGroup, clientUCase)
}
