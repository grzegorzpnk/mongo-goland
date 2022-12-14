package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mongoGolang/src/pkg/structure"
	"net/http"
	"time"
)

type apiHandler struct {
	clientMongo *mongo.Client
}

func (h *apiHandler) setClient(client *mongo.Client) {
	h.clientMongo = client
}

func (h *apiHandler) CreateClusterHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("New request")
	response.Header().Add("content-type", "application/json")
	var cluster structure.Cluster
	err := json.NewDecoder(request.Body).Decode(&cluster)
	if err != nil {
		fmt.Println(fmt.Errorf("error: " + err.Error()))
	}
	fmt.Println(cluster)
	collection := h.clientMongo.Database("topology").Collection("clusters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, cluster)
	if err != nil {
		log.Fatal("BLAD", err)
	}
	err = json.NewEncoder(response).Encode(result)
	if err != nil {
		fmt.Println(fmt.Errorf("error: " + err.Error()))
	}

}

func (h *apiHandler) GetClustersHandler(response http.ResponseWriter, _ *http.Request) {
	fmt.Println("GET Clusters method")
	response.Header().Add("content-type", "application/json")

	var cluster []structure.Cluster

	collection := h.clientMongo.Database("topology").Collection("clusters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err2 := response.Write([]byte(`{ "message 1": "` + err.Error() + `" }`))
		fmt.Println(err2.Error())
		return
	}

	for cursor.Next(ctx) {
		var clusterTmp structure.Cluster
		err := cursor.Decode(&clusterTmp)
		if err != nil {
			fmt.Printf("error: " + err.Error())
		}
		cluster = append(cluster, clusterTmp)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err2 := response.Write([]byte(`{ "message 2": "` + err.Error() + `" }`))
		fmt.Println(err2.Error())
		return
	}

	err2 := json.NewEncoder(response).Encode(cluster)
	if err2 != nil {
		fmt.Printf("error: " + err2.Error())
	}

}
