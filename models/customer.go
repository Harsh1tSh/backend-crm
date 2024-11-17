package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	Company string             `json:"company" bson:"company"`
	Status  string             `json:"status" bson:"status"`
	Notes   string             `json:"notes" bson:"notes"`
}
