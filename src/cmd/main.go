package main

import (
	"mongoGolang/src/api"
	"mongoGolang/src/config"
	"mongoGolang/src/pkg/mongoClient"
	"net/http"
)

func main() {

	client := mongoClient.ConnectToMongo(config.GetConfiguration().MongoEndpoint, config.GetConfiguration().MongoPort)
	router := api.NewRouter(client)
	http.ListenAndServe(config.GetConfiguration().ClientPort, router)
}
