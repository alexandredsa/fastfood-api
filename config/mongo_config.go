package config

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

type MongoConfig struct {
	mongoClient *mongo.Client
}

func (mc *MongoConfig) GetClient() *mongo.Client {
	var err error

	if mc.mongoClient != nil {
		return mc.mongoClient
	}

	mc.mongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mc.mongoClient.Connect(ctx)
	return mc.mongoClient
}
func (mc *MongoConfig) GetDatabase(databaseName string) *mongo.Database {
	client := mc.GetClient()
	return client.Database(databaseName)
}
func (mc *MongoConfig) GetCollection(collectionName string) *mongo.Collection {
	database := mc.GetDatabase("foodranks")
	return database.Collection(collectionName)
}
