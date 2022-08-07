package model

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"booking_system/database"
	"booking_system/model/schema"
	"booking_system/utils"
)

func MgoCreateUser(user schema.MongoUser) (schema.MongoUser, bool) {
	isExisted := false
	userId := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(user.Email))
	user.UserID = userId

	isExistedUser := MgoFindUserById(userId.String())

	if isExistedUser.UserID == userId {
		isExisted = true
		return schema.MongoUser{}, isExisted
	}

	// TODO: The timestamp here can be changed to receive parameters later.
	p_id := primitive.NewObjectIDFromTimestamp(time.Now())

	insertUser := schema.MongoUser{
		Id:     p_id,
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
	}
	result, _ := database.MgoConnect.InsertOne(context.TODO(), insertUser)
	log.Println(result)
	return insertUser, isExisted
}

func MgoFindUsers() []schema.MongoUser {
	var users []schema.MongoUser
	cursor, err := database.MgoConnect.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var user schema.MongoUser
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func MgoFindUserById(id string) schema.MongoUser {
	user := schema.MongoUser{}
	filter := utils.BsonIdUtil(id)
	err := database.MgoConnect.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Println(err)
	}
	return user
}

func MgoUpdateUser(id string, user schema.User) schema.User {
	filter := utils.BsonIdUtil(id)

	updateUserData := bson.M{"$set": user}
	result, err := database.MgoConnect.UpdateOne(context.TODO(), filter, updateUserData)

	if err != nil {
		log.Println(err)
		return schema.User{}
	}
	log.Println(result.ModifiedCount)

	return user
}

func MgoDeleteUserById(id string) bool {
	filter := utils.BsonIdUtil(id)

	result, err := database.MgoConnect.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(result)

	return true
}
