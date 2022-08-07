package utils

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonIdUtil(id string) primitive.M {
	filter := bson.M{}

	isObjectId := primitive.IsValidObjectID(id)

	if isObjectId {
		objectId, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objectId}
	} else {
		id := uuid.MustParse(id)
		filter = bson.M{"user_id": id}
	}

	return filter

}
