package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Berhasil terhubung ke MongoDB!")
	TaskCollection = client.Database("belajar_go").Collection("tasks")
	UserCollection = client.Database("belajar_go").Collection("users")

	return client
}
