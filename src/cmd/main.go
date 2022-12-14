package main

import (
	"mongoGolang/src/api"
	"mongoGolang/src/pkg/mongoClient"
	"net/http"
)

func main() {

	client := mongoClient.ConnectToMongo()
	router := api.NewRouter(client)
	http.ListenAndServe(":12345", router)

}
