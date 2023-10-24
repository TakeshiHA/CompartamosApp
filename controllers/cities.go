package controllers

import (
	"github.com/TakeshiHA/test-middleware/usecase"
	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase *usecase.CityUsecase
}

func NewCityController(group *echo.Group, cityUCase *usecase.CityUsecase) {
	handler := &CityController{
		cityUsecase: cityUCase,
	}

	group.GET("", handler.GetCities)
}

// @Summary	CITY01 GetCities
// @Tags Cities
// @Description GetAttentionsNoEmo
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.City
// @Failure 400 {object} models.ResponseError
// @Router /api/cities [get]
func (m *CityController) GetCities(ctx echo.Context) error {
	c := ctx.Request().Context()

	emos, err := m.cityUsecase.GetCities(c)
	if err != nil {
		return err
	}
	return ctx.JSON(200, emos)
}
