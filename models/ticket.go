package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CustomerID  primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"` // Open, In Progress, Resolved
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
}
