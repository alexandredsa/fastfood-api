package repositories

import (
	"context"
	"fastfood-api/config"
	"fastfood-api/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodReviewRepository struct{}

func (frr *FoodReviewRepository) getCollection() *mongo.Collection {
	mongoConfig := config.MongoConfig{}
	return mongoConfig.GetCollection("foodReviews")
}

func (frr *FoodReviewRepository) Save(foodReview *models.FoodReview) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := frr.getCollection()
	foodReview.ID = primitive.NewObjectID()
	collection.InsertOne(ctx, foodReview)
}
func (frr *FoodReviewRepository) FindAllByTypeId(foodTypeId primitive.ObjectID) (foodReviews []*models.FoodReview) {
	collection := frr.getCollection()
	cur, err := collection.Find(context.TODO(), bson.M{"foodSummaryId": foodTypeId})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.FoodReview
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		foodReviews = append(foodReviews, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return foodReviews
}
