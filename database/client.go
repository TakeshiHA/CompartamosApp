package database

import (
	"context"

	"github.com/TakeshiHA/test-middleware/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientCollection struct {
	Collection *mongo.Collection
}

func GetClientCollection() *ClientCollection {
	db := DB
	client := db.Collection("clients")
	return &ClientCollection{Collection: client}
}

func (cli *ClientCollection) CreateClient(ctx context.Context, body *models.Client) (*models.Client, error) {
	if len(body.ID) == 0 {
		body.ID = primitive.NewObjectID().Hex()
	}
	_, err := cli.Collection.InsertOne(ctx, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (wp *ClientCollection) GetClients(ctx context.Context) ([]*models.Client, error) {
	filter := bson.M{}
	res, err := wp.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	arrClients := []*models.Client{}
	erro := res.All(ctx, &arrClients)
	if erro != nil {
		return nil, erro
	}
	return arrClients, nil
}

func (wp *ClientCollection) GetClientById(ctx context.Context, id string) (*models.Client, error) {
	filter := bson.M{"_id": id}
	res := wp.Collection.FindOne(ctx, filter)
	var attention models.Client
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Decode(&attention)
	if err != nil {
		return nil, err
	}
	return &attention, nil
}

func (wp *ClientCollection) GetClientByDNI(ctx context.Context, dni string) (*models.Client, error) {
	filter := bson.M{"dni": dni}
	res := wp.Collection.FindOne(ctx, filter)
	var attention models.Client
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Decode(&attention)
	if err != nil {
		return nil, err
	}
	return &attention, nil
}

func (wp *ClientCollection) UpdateClient(ctx context.Context, id string, client *models.Client) (bool, error) {
	filter := bson.M{"_id": id}
	v := StructToInterface(client)
	update := bson.M{"$set": v}
	_, err := wp.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (wp *ClientCollection) DeleteClientById(ctx context.Context, id string) (bool, error) {
	filter := bson.M{"_id": id}
	res, err := wp.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	if res.DeletedCount > 0 {
		return true, nil
	}
	return false, nil
}
