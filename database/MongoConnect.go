package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgoConnect *mongo.Collection

var client *mongo.Client

func MongoDb() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://192.168.56.102:27017"))

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("GolangApi").Collection("TestGolangApi")
	MgoConnect = collection
}

func MongoDisconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
