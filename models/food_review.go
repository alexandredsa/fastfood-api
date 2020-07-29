package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodReview struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name          string             `json:"name,omitempty"`
	Flavor        float64            `json:"flavor,omitempty"`
	Sauce         float64            `json:"sauce,omitempty"`
	Texture       float64            `json:"texture,omitempty"`
	FoodSummaryId primitive.ObjectID `bson:"foodSummaryId" json:"foodSummaryId,omitempty"`
}
