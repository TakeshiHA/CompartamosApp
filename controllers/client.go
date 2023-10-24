package controllers

import (
	"encoding/json"

	"github.com/TakeshiHA/test-middleware/models"
	"github.com/TakeshiHA/test-middleware/usecase"
	"github.com/labstack/echo/v4"
)

type ClientController struct {
	clientUsecase *usecase.ClientUsecase
}

func NewClientController(group *echo.Group, clientUCase *usecase.ClientUsecase) {
	handler := &ClientController{
		clientUsecase: clientUCase,
	}

	group.POST("", handler.CreateClient)
	group.GET("", handler.GetClients)
	group.PUT("/:id", handler.UpdateClient)
	group.DELETE("/:id", handler.DeleteClientById)
	group.GET("/:id", handler.GetClientById)
}

// @Summary	CLIENT01 CreateClient
// @Tags Clients
// @Description CreateClient
// @Param   client     body    models.Client     true        "New Client"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Client
// @Failure 400 {object} models.ResponseError
// @Router /api/clients [post]
func (m *ClientController) CreateClient(ctx echo.Context) error {
	c := ctx.Request().Context()

	var client models.Client
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&client)
	if errDecode != nil {
		return errDecode
	}

	emos, err := m.clientUsecase.CreateClient(c, &client)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}

// @Summary	CLIENT02 GetClients
// @Tags Clients
// @Description GetClients
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Client
// @Failure 400 {object} models.ResponseError
// @Router /api/clients [get]
func (m *ClientController) GetClients(ctx echo.Context) error {
	c := ctx.Request().Context()

	emos, err := m.clientUsecase.GetClients(c)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}

// @Summary	CLIENT03 UpdateClient
// @Tags Clients
// @Description UpdateClient
// @Param   id     path    string     true        "ID"
// @Param   client     body    models.Client     true        "New Client"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Client
// @Failure 400 {object} models.ResponseError
// @Router /api/clients/{id} [put]
func (m *ClientController) UpdateClient(ctx echo.Context) error {
	c := ctx.Request().Context()

	var client models.Client
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&client)
	if errDecode != nil {
		return errDecode
	}
	id := ctx.Param("id")

	emos, err := m.clientUsecase.UpdateClient(c, id, &client)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}

// @Summary	CLIENT04 DeleteClientById
// @Tags Clients
// @Description DeleteClientById
// @Param   id     path    string     true        "ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Client
// @Failure 400 {object} models.ResponseError
// @Router /api/clients/{id} [delete]
func (m *ClientController) DeleteClientById(ctx echo.Context) error {
	c := ctx.Request().Context()

	id := ctx.Param("id")

	emos, err := m.clientUsecase.DeleteClientById(c, id)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}

// @Summary	CLIENT05 GetClientById
// @Tags Clients
// @Description GetClientById
// @Param   id     path    string     true        "ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Client
// @Failure 400 {object} models.ResponseError
// @Router /api/clients/{id} [get]
func (m *ClientController) GetClientById(ctx echo.Context) error {
	c := ctx.Request().Context()

	id := ctx.Param("id")

	emos, err := m.clientUsecase.GetClientById(c, id)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}
