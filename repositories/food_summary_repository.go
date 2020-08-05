package repositories

import (
	"context"
	"errors"
	"fastfood-api/config"
	"fastfood-api/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodSummaryRepository struct{}

func (fsr *FoodSummaryRepository) getCollection() *mongo.Collection {
	mongoConfig := config.MongoConfig{}
	return mongoConfig.GetCollection("foodSummaries")
}

func (fsr *FoodSummaryRepository) Save(foodSummary models.FoodSummary) {
	collection := fsr.getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, foodSummary)
}

func (fsr *FoodSummaryRepository) FindById(foodSummaryId primitive.ObjectID) *models.FoodSummary {
	collection := fsr.getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var foodSummary models.FoodSummary
	collection.FindOne(ctx, bson.M{"_id": foodSummaryId}).Decode(&foodSummary)
	return &foodSummary
}

func (fsr *FoodSummaryRepository) FindAll() (foodSummaries []*models.FoodSummary) {
	collection := fsr.getCollection()
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.FoodSummary
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		foodSummaries = append(foodSummaries, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return foodSummaries
}

func (fsr *FoodSummaryRepository) IncreaseSummary(foodSummary *models.FoodSummary) (err error) {
	currentSummary := fsr.FindById(foodSummary.ID)
	if currentSummary == nil {
		return errors.New("Summary not exists")
	}
	currentSummary.Count++
	fsr.UpdateSummary(currentSummary)
	return nil
}

func (fsr *FoodSummaryRepository) DecreaseSummary(foodSummary *models.FoodSummary) (err error) {
	currentSummary := fsr.FindById(foodSummary.ID)
	if currentSummary == nil {
		return errors.New("Summary not exists")
	}
	currentSummary.Count--
	fsr.UpdateSummary(currentSummary)
	return nil
}

func (fsr *FoodSummaryRepository) UpdateSummary(foodSummary *models.FoodSummary) (err error) {
	collection := fsr.getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection.UpdateOne(ctx, bson.M{"_id": foodSummary.ID}, bson.D{
		{"$set", bson.D{{"count", foodSummary.Count}}},
	})
	return nil
}
