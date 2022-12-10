package mongoGolang

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var client *mongo.Client

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")

	router := mux.NewRouter()
	router.HandleFunc("/cluster", CreateClusterHandler).Methods("POST")
	http.ListenAndServe(":80", router)
}

func CreateClusterHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var cluster Cluster
	json.NewDecoder(request.Body).Decode(&cluster)
	collection := client.Database("topology").Collection("clusters")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, cluster)
	json.NewEncoder(response).Encode(result)
}

type Cluster struct {
	name      string
	resources Resources
}

type Resources struct {
	cpu    int
	memory int
}
