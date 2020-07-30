package repositories

import (
	"context"
	"fastfood-api/config"
	"fastfood-api/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

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

func (frr *FoodReviewRepository) Delete(foodReviewId primitive.ObjectID) {
	collection := frr.getCollection()
	collection.DeleteOne(context.TODO(), bson.M{"_id": foodReviewId})
}

func (frr *FoodReviewRepository) FindAllByTypeId(foodTypeId primitive.ObjectID) (foodReviews []*models.FoodReview) {
	collection := frr.getCollection()
	findOptions := options.Find()
	// Sort by `price` field descending
	match := bson.D{{"$match", bson.D{{"foodSummaryId", foodTypeId}}}}
	addFields := bson.D{{"$addFields", bson.D{{"position", bson.D{{"$add", []string{"$sauce", "$flavor", "$texture"}}}}}}}
	sort := bson.D{{"$sort", bson.D{{"position", -1}}}}
	findOptions.SetSort(bson.D{{"price", -1}})
	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match, addFields, sort})
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
