package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person is model for People
type Person struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Age  int    `json:"age,omitempty" bson:"age,omitempty"`
}

// People is slice of People
var People []Person