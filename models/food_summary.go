package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodSummary struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Count int8               `json:"count,omitempty"`
}
