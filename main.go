package main

import (
	"github.com/53jk1/go-base/endpoints"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// main is the entry point of the application
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/validate", endpoints.Validate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}