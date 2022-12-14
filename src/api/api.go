package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var r *mux.Router

func NewRouter(mongoClient *mongo.Client) *mux.Router {

	var handler apiHandler
	handler.setClient(mongoClient)

	r = mux.NewRouter().PathPrefix("/v1").Subrouter()
	baseUrl := ""

	r.HandleFunc(baseUrl+"/cluster", handler.CreateClusterHandler).Methods("POST")
	r.HandleFunc(baseUrl+"/cluster", handler.GetClustersHandler).Methods("GET")

	return r

}
