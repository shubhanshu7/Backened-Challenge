package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/orders", JwtVerify(GetOrders)).Methods("GET")
	router.HandleFunc("/orders", JwtVerify(CreateOrder)).Methods("POST")
	return router
}
