package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"-" json:"password,omitempty"` // Plain text password, not stored in DB
	PasswordHash string             `bson:"password_hash" json:"-"`
}
