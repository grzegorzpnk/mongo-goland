package main

import "github.com/gorilla/mux"

//var r *mux.Router

func NewRouter() *mux.Router {

	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	baseUrl := ""
	//refactored:
	r.HandleFunc(baseUrl+"/cluster", CreateClusterHandler).Methods("POST")
	r.HandleFunc(baseUrl+"/cluster", GetClustersHandler).Methods("GET")

	return r

}
