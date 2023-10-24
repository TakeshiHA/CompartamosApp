package database

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

type MyDatabase struct {
	MongoDB *mongo.Database
}

func InitDatabases() *MyDatabase {
	mongoDB := initMongoDB()
	DB = mongoDB

	return &MyDatabase{
		MongoDB: mongoDB,
	}
}

func initMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb+srv://testAdmin:ROJXETCx40n2fbYb@golang-middleware.ty0iwrb.mongodb.net/") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client.Database("test-golang")
}

func (dbs *MyDatabase) DisconnectDatabases() {
	err := dbs.MongoDB.Client().Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func StructToInterface(obj interface{}) interface{} {
	vp := reflect.New(reflect.TypeOf(obj))
	vp.Elem().Set(reflect.ValueOf(obj))
	return vp.Interface()
}
