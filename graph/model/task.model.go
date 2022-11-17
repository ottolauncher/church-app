package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID `json:"id" bson:"id"`
	Title     string             `json:"title" bson:"title"`
	Slug      string             `json:"slug" bson:"slug,omitempty"`
	Note      string             `json:"note" bson:"note"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}
