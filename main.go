package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://10.254.185.109:27017")
	_, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
		//fmt.Errorf("error", err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/cluster", CreateClusterHandler).Methods("POST")
	http.ListenAndServe(":12345", router)

}

func CreateClusterHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("New request")
	response.Header().Add("content-type", "application/json")
	var cluster Cluster
	json.NewDecoder(request.Body).Decode(&cluster)
	collection := client.Database("topology").Collection("clusters")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, cluster)
	json.NewEncoder(response).Encode(result)
}
