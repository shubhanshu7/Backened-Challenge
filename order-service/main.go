package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8003", router))
}
