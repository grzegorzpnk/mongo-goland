package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

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

func GetClustersHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("GET Clusters method")
	response.Header().Add("content-type", "application/json")

	var cluster []Cluster

	collection := client.Database("topology").Collection("clusters")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message 1": "` + err.Error() + `" }`))
		return
	}

	for cursor.Next(ctx) {
		var cluster_tmp Cluster
		cursor.Decode(&cluster_tmp)
		cluster = append(cluster, cluster_tmp)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message 2": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(cluster)

}
