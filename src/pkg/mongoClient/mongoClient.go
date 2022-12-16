package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongo(ip, port string) *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(ip + port)
	client, _ := mongo.Connect(ctx, clientOptions)

	return client

}
