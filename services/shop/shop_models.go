package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Shop struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	OwnerID     primitive.ObjectID `json:"ownerId" bson:"ownerId"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	CreatedAt   primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
