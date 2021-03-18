package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model ...
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	LastName string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
}
