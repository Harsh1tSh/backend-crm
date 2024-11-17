package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Message   string             `json:"message" bson:"message"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}
