package database

import (
	"context"

	"github.com/TakeshiHA/test-middleware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CityCollection struct {
	Collection *mongo.Collection
}

func GetCityCollection() *CityCollection {
	db := DB
	city := db.Collection("cities")
	return &CityCollection{Collection: city}
}

func (city *CityCollection) GetCities(ctx context.Context) ([]*models.City, error) {
	filter := bson.M{}
	cur, err := city.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if cur.Err() != nil {
		return nil, cur.Err()
	}
	cities := []*models.City{}
	erro := cur.All(ctx, &cities)
	if erro != nil {
		return nil, erro
	}
	return cities, nil
}
