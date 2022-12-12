package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var client *mongo.Client

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://10.254.185.109:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/cluster", CreateClusterHandler).Methods("POST")
	http.ListenAndServe(":12345", router)

}

func CreateClusterHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("New request")
	response.Header().Add("content-type", "application/json")
	var cluster Cluster
	json.NewDecoder(request.Body).Decode(&cluster)
	fmt.Println(cluster)
	collection := client.Database("topology").Collection("clusters")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, cluster)
	if err != nil {
		log.Fatal("BLAD", err)
	}
	json.NewEncoder(response).Encode(result)

}
