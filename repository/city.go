package repository

import (
	"context"

	"github.com/TakeshiHA/test-middleware/database"
	"github.com/TakeshiHA/test-middleware/models"
)

type CityRepository struct {
	cityCollection *database.CityCollection
}

type CityRepositoryInterface interface {
	GetCities(ctx context.Context) ([]*models.City, error)
}

func NewCityRepository() *CityRepository {
	return &CityRepository{database.GetCityCollection()}
}

func (cityRepo *CityRepository) GetCities(ctx context.Context) ([]*models.City, error) {
	cities, err := cityRepo.cityCollection.GetCities(ctx)
	return cities, err
}
