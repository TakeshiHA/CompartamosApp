package database

import (
	"context"

	"github.com/TakeshiHA/test-middleware/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	Collection *mongo.Collection
}

func GetUserCollection() *UserCollection {
	db := DB
	user := db.Collection("users")
	return &UserCollection{Collection: user}
}

func (usr *UserCollection) CreateUser(ctx context.Context, body *models.User) (*models.User, error) {
	if len(body.ID) == 0 {
		body.ID = primitive.NewObjectID().Hex()
	}
	_, err := usr.Collection.InsertOne(ctx, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
