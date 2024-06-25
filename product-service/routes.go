package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", JwtVerify(GetProducts)).Methods("GET")
	router.HandleFunc("/products", JwtVerify(CreateProduct)).Methods("POST")
	router.HandleFunc("/products", JwtVerify(UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products", JwtVerify(DeleteProduct)).Methods("DELETE")
	return router
}
