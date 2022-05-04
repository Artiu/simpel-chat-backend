package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init() {
	var err error

	uri := os.Getenv("MONGODB_URI")
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB")
}

func Disconnect() {
	Client.Disconnect(context.TODO())
	log.Println("Disconnected from MongoDB")
}
