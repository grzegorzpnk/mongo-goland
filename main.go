package mongoGolang

import (
	"context"
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
	http.ListenAndServe(":80", router)
}

func CreateClusterHandler(response http.ResponseWriter, request *http.Request) {

}

type Cluster struct {
	name      string
	resources Resources
}

type Resources struct {
	cpu    int
	memory int
}
