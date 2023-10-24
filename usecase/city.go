package usecase

import (
	"context"
	"time"

	"github.com/TakeshiHA/test-middleware/models"
	"github.com/TakeshiHA/test-middleware/repository"
	"github.com/labstack/echo/v4"
)

type CityUsecase struct {
	cityRepository *repository.CityRepository
	contextTimeout time.Duration
}

type CityUsecaseInterface interface {
	GetCities(ctx context.Context) ([]*models.City, *echo.HTTPError)
}

func NewCityUsecase(
	cityRepository *repository.CityRepository,
	timeout time.Duration,
) *CityUsecase {
	return &CityUsecase{
		cityRepository: cityRepository,
		contextTimeout: timeout,
	}
}

func (cityUCase *CityUsecase) GetCities(ctx context.Context) ([]*models.City, *echo.HTTPError) {
	ctx, cancel := context.WithTimeout(ctx, cityUCase.contextTimeout)
	defer cancel()

	cities, err := cityUCase.cityRepository.GetCities(ctx)

	if err != nil {
		return nil, echo.NewHTTPError(400, models.ResponseError{Message: "Error! Contáctese con el administrador."})
	}

	return cities, nil
}
