package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Price       float64            `bson:"price,omitempty" json:"price"`
	Quantity    int                `bson:"quantity,omitempty" json:"quantity"`
}
