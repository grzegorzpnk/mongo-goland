package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client

func main() {

	connectToMongo()

	router := NewRouter()
	http.ListenAndServe(":12345", router)

}

func connectToMongo() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://10.254.185.109:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

}
