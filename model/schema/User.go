package schema

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    uuid.UUID `from:"Id"`
	Name  string    `json:"Name" binding:"required"`
	Email string    `json:"Email" binding:"required"`
	Phone string    `json:"Phone" binding:"required"`
}

type MongoUser struct {
	Id     primitive.ObjectID `from:"Id" bson:"_id"`
	UserID uuid.UUID          `from:"UserID" bson:"user_id"`
	Name   string             `json:"Name" binding:"required"`
	Email  string             `json:"Email" binding:"required"`
	Phone  string             `json:"Phone" binding:"required"`
}
