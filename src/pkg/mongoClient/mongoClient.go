package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongo() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://10.254.185.109:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	return client

}
