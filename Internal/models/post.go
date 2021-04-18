// Package models provides the models.
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post model ...
type Post struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Body string             `json:"body,omitempty" bson:"body,omitempty"`
	Slug string             `json:"slug,omitempty" bson:"slug,omitempty"`
}
